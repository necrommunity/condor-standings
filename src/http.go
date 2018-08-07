package main

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"
)

const (
	sessionName = "condor.sid"
)

var (
	sessionStore sessions.CookieStore
	myHTTPClient = &http.Client{ // We don't want to use the default http.Client structure because it has no default timeout set
		Timeout: 10 * time.Second,
	}
)

// ----------------
// Data structures
// ----------------

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}

type Event struct {
	EventName    string `json:"eventName"`
	Participants []models.Participant
}

type TemplateData struct {
	Title        string
	UsersTotal   int
	UserAccounts []models.UserAccount
	FoundTables  []models.FoundTable
}

//  -----------------------
//  Initialization function
//  -----------------------

func httpInit() {
	// Create a new Gin HTTP router
	gin.SetMode(gin.ReleaseMode) // Comment this out to debug HTTP stuff
	httpRouter := gin.Default()

	// Read some HTTP server configuration values from environment variables
	// (they were loaded from the .env file in main.go)
	sessionSecret := os.Getenv("SESSION_SECRET")
	if len(sessionSecret) == 0 {
		log.Info("The \"SESSION_SECRET\" environment variable is blank; aborting HTTP initalization.")
		return
	}

	domain := os.Getenv("DOMAIN")
	if len(domain) == 0 {
		log.Info("The \"DOMAIN\" environment variable is blank; aborting HTTP initalization.")
		return
	}
	tlsCertFile := os.Getenv("TLS_CERT_FILE")
	tlsKeyFile := os.Getenv("TLS_KEY_FILE")
	useTLS := true
	if len(tlsCertFile) == 0 || len(tlsKeyFile) == 0 {
		useTLS = false
	}

	// Create a session store
	sessionStore = sessions.NewCookieStore([]byte(sessionSecret))
	options := sessions.Options{
		Path:   "/",
		Domain: domain,
		MaxAge: 5, // 5 seconds
		// After getting a cookie via "/login", the client will immediately
		// establish a WebSocket connection via "/ws", so the cookie only needs
		// to exist for that time frame
		Secure: true,
		// Only send the cookie over HTTPS:
		// https://www.owasp.org/index.php/Testing_for_cookies_attributes_(OTG-SESS-002)
		HttpOnly: true,
		// Mitigate XSS attacks:
		// https://www.owasp.org/index.php/HttpOnly
	}
	if !useTLS {
		options.Secure = false
	}
	sessionStore.Options(options)
	httpRouter.Use(sessions.Sessions(sessionName, sessionStore))

	// Use the Tollbooth Gin middleware for rate limiting
	limiter := tollbooth.NewLimiter(1, nil) // Limit each user to 1 request per second
	// When a user requests "/", they will also request the CSS and images;
	// this middleware is smart enough to know that it is considered part of the first request
	// However, it is still not possible to spam download CSS or image files
	limiterMiddleware := tollbooth_gin.LimitHandler(limiter)
	httpRouter.Use(limiterMiddleware)

	// Use a custom middleware for Google Analytics tracking
	// GATrackingID = os.Getenv("GA_TRACKING_ID")
	// if len(GATrackingID) != 0 {
	// 	httpRouter.Use(httpMwGoogleAnalytics)
	// }

	// Path handlers (for the website)
	httpRouter.GET("/", httpHome)
	httpRouter.GET("/api", httpAPI)                   // Handles static API
	httpRouter.GET("/api/event", httpEventDocAPI)     // Handles specific event calls
	httpRouter.GET("/api/event/:event", httpEventAPI) // Handles specific event calls
	httpRouter.GET("/api/teamresults", httpTeamAPI)
	// Static handlers (for the website)
	httpRouter.Static("/public", "../public")

	// Figure out the port that we are using for the HTTP server
	var port int
	if useTLS {
		// We want all HTTP requests to be redirected to HTTPS
		// (but make an exception for Let's Encrypt)
		// The Gin router is using the default serve mux, so we need to create a
		// new fresh one for the HTTP handler
		HTTPServeMux := http.NewServeMux()
		HTTPServeMux.Handle("/.well-known/acme-challenge/", http.FileServer(http.FileSystem(http.Dir("../letsencrypt"))))
		HTTPServeMux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
		}))

		// ListenAndServe is blocking, so start listening on a new goroutine
		go func() {
			http.ListenAndServe(":80", HTTPServeMux) // Nothing before the colon implies 0.0.0.0
			log.Fatal("http.ListenAndServe ended for port 80.", nil)
		}()

		// 443 is the default port for HTTPS
		port = 443
	} else {
		// 80 is the defeault port for HTTP
		port = 80
	}

	// Start listening and serving requests (which is blocking)
	log.Info("Listening on port " + strconv.Itoa(port) + ".")
	if useTLS {
		if err := http.ListenAndServeTLS(
			":"+strconv.Itoa(port), // Nothing before the colon implies 0.0.0.0
			tlsCertFile,
			tlsKeyFile,
			httpRouter,
		); err != nil {
			log.Fatal("http.ListenAndServeTLS failed:", err)
		}
		log.Fatal("http.ListenAndServeTLS ended prematurely.", nil)
	} else {
		// Listen and serve (HTTP)
		if err := http.ListenAndServe(
			":"+strconv.Itoa(port), // Nothing before the colon implies 0.0.0.0
			httpRouter,
		); err != nil {
			log.Fatal("http.ListenAndServe failed:", err)
		}
		log.Fatal("http.ListenAndServe ended prematurely.", nil)
	}
}

//
//	HTTP miscellaneous subroutines
//

func httpServeTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	lp := path.Join("views", "layout.tmpl")
	fp := path.Join("views", templateName+".tmpl")

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Create the template
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Error("Failed to create the template: " + err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Execute the template and send it to the user
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		if strings.HasSuffix(err.Error(), ": write: broken pipe") ||
			strings.HasSuffix(err.Error(), ": client disconnected") ||
			strings.HasSuffix(err.Error(), ": http2: stream closed") ||
			strings.HasSuffix(err.Error(), ": write: connection timed out") {

			// Broken pipe errors can occur when the user presses the "Stop" button while the template is executing
			// We don't want to reporting these errors to Sentry
			// https://stackoverflow.com/questions/26853200/filter-out-broken-pipe-errors-from-template-execution
			// I don't know exactly what the other errors mean
			log.Info("Ordinary error when executing the template: " + err.Error())
		} else {
			log.Error("Failed to execute the template: " + err.Error())
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
