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
		team[tWinner]["Total"] = team[tWinner]["Total"] + 1
		team[tWinner][tLoser] = wins + 1
	}
	teamList := map[string][]string{
		"Dark Cookies": []string{
			"abu__yazan",
			"biggiemac42",
			"chef_mayhem",
			"Flygluffet",
			"incnone",
			"Jamblur",
			"JustSparkyYes",
			"Kova46",
			"Lucoa",
			"mantasMBL",
			"Megamissingn0",
			"Spootybiscuit",
			"WuffWuff",
		},
		"Frozen Cheese": []string{
			"arborelia",
			"bastet222",
			"dsmidna",
			"duneaught",
			"firebrde",
			"Kailaria",
			"moyuma", 
			"Slimo",
			"Squega",
			"Staekk",
			"wow_tomato",
			"yamiramiz",
		},
		"Italian Carrots": []string{
			"ARTQ",
			"brumekuroi",
			"Cyber_1",
			"Minhs2",
			"mpr",
			"pancelor",
			"professionaltwitchtroll",
			"ratata_ratata",
			"Rachelsatx",
			"Revalize",
			"Siveure",
			"thedarkfreaaack",
			"yuka34",
		},
		"Regular Hams": []string{
			"alex42918",
			"Ancalagor",
			"Call_Me_Kaye",
			"definitely_not_HIM",
			"EpicSuccess",
			"JackOfGames",
			"mudjoe2",
			"Paratroopa1",
			"ptrevortactyl",
			"Rotomington",
			"saakas0206",
			"seanpwolf",
			"uniowen",
		},
		"Stinkin' Rebels": []string{
			"ekimekim",
			"fiverfiverone",
			"KingTorture",
			"mayaundefined",
			"Ratracing",
			"reijigazpacho",
			"RoyalGoof",
			"sailormint",
			"supervillain_joe",
			"tetel",
			"Tufwfo",
		},
	}
	headings := []string{"Dark Cookies", "Frozen Cheese", "Italian Carrots", "Regular Ham", "Stinkin' Rebels", "Total"}

	// Send data to template
	data := TemplateData{
		Title:    "Team Results",
		Results:  team,
		Headers:  headings,
		TeamList: teamList,
	}
	httpServeTemplate(w, "teamresults", data)
}
