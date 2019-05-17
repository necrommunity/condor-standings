package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"
)

// Teams is a var for each team

func getProp(t models.Result, field string) string {
	r := reflect.ValueOf(t)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

func httpTeamResults(c *gin.Context) {
	// Local variables
	w := c.Writer

	// Used to get JSON results
	url := "https://some.pizza/api/teamresults"
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
			"Total":           0,
		},
		"Frozen Cheese": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
		"Italian Carrots": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
		"Regular Ham": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
		"Stinkin' Rebels": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
	}
	var teamAll = map[string]map[string]int{
		"Dark Cookies": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
		"Frozen Cheese": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
		"Italian Carrots": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
		"Regular Ham": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
		"Stinkin' Rebels": map[string]int{
			"Dark Cookies":    0,
			"Frozen Cheese":   0,
			"Italian Carrots": 0,
			"Regular Ham":     0,
			"Stinkin' Rebels": 0,
			"Total":           0,
		},
	}
	// Parse through JSON results and fill in team map
	wins := 0
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
		wins = team[tWinner][tLoser]
		if tWinner != tLoser {
			team[tWinner]["Total"] = team[tWinner]["Total"] + 1
			team[tWinner][tLoser] = wins + 1
		}
	}
	wins = 0

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
		wins = teamAll[tWinner][tLoser]
		teamAll[tWinner]["Total"] = teamAll[tWinner]["Total"] + 1
		teamAll[tWinner][tLoser] = wins + 1
	}

	headings := []string{"Dark Cookies", "Frozen Cheese", "Italian Carrots", "Regular Ham", "Stinkin' Rebels", "Total"}

	// Send data to template
	data := TemplateData{
		Title:      "Team Results",
		Results:    team,
		ResultsAll: teamAll,
		Headers:    headings,
		TeamList:   teamList,
	}
	httpServeTemplate(w, "teamresults", data)
}
