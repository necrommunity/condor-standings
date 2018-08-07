package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// Teams is a var for each team
type Teams struct {
	DarkCookies    int `json:"Dark Cookies"`
	FrozenCheese   int `json:"Frozen Cheese"`
	ItalianCarrots int `json:"Italian Carrots"`
	RegularHam     int `json:"Regular Ham"`
	StinkinRebels  int `json:"Stinkin' Rebels"`
}

func getProp(t models.Result, field string) string {
	r := reflect.ValueOf(t)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

func httpTeamResults(c *gin.Context) {
	// Local variables
	w := c.Writer

	// Used to get JSON results
	url := "https://wow.freepizza.how/api/teamresults"
	urlClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error(err)
	}
	req.Header.Set("User-Agent", "it's me!")
	res, getErr := urlClient.Do(req)
	if getErr != nil {
		log.Error(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Error(readErr)
	}

	// Build the results struct
	var Results []models.Result
	jsonErr := json.Unmarshal(body, &Results)
	if jsonErr != nil {
		log.Error(jsonErr)
	}

	// Build the team map with init values
	var team = map[string]map[string]int{
		"Dark Cookies": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
		},
		"Frozen Cheese": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
		},
		"Italian Carrots": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
		},
		"Regular Ham": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
		},
		"Stinkin' Rebels": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
		},
	}

	// Parse through JSON results and fill in team map
	for i := range Results {
		wi := Results[i].Winner
		lo := 2
		if wi == 2 {
			lo = 1
		}
		tWin := "Team" + strconv.Itoa(wi)
		tLos := "Team" + strconv.Itoa(lo)
		tWinner := getProp(Results[i], tWin)
		tLoser := getProp(Results[i], tLos)
		wins := team[tWinner][tLoser]
		team[tWinner][tLoser] = wins + 1
	}

	// Send data to template
	data := TemplateData{
		Title: "Team Results",
		Teams: team,
	}
	httpServeTemplate(w, "teamresults", data)
}
