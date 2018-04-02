// @APIVersion 1.0.0
// @Title My Cool API
// @Description My API usually works as expected. But sometimes its not true
// @BasePath http://wow.freepizza.how/
// @Contact sillypairs@gmail.com
// @TermsOfServiceUrl http://google.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause

package main // In Go, executable commands must always use package main

import (
	"github.com/joho/godotenv"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"
	"os"
	"path"
)

var (
	projectPath = path.Join(os.Getenv("GOPATH"), "src", "github.com", "sillypears", "condor-standings")
	db          *models.Models
)

func main() {
	// Initalize the Logger
	log.Init()
	// Welcome message
	log.Info("+-------------------------------+")
	log.Info("|        Starting server        |")
	log.Info("+-------------------------------+")

	// Load the ".env" file which contains environment variables with secret values
	if err := godotenv.Load(path.Join(projectPath, ".env")); err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	// Initialize the database model
	if v, err := models.Init(); err != nil {
		log.Fatal("Failed to open the database:", err)
	} else {
		db = v
	}
	defer db.Close()

	// Start the http server
	httpInit()
}
