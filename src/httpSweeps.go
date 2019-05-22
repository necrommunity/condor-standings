package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	// "strings"
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"
)

//  --------------------------------
//  The landing page for the website
//  --------------------------------

// func httpSweeps(c *gin.Context) {
func httpSweeps(c *gin.Context) {
// 	// Local variables
	w := c.Writer

	url := "https://some.pizza/api/sweeps"
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
	var SweepsResults []models.Sweep
	jsonErr := json.Unmarshal(body, &SweepsResults)
	if jsonErr != nil {
		log.Error(jsonErr)
	}
	
	var AutoGenned []models.Sweep
	var Challenges []models.Sweep
	totalSweeps := len(SweepsResults)

	for _, sweep := range SweepsResults {
		// temp, err := fmt.Printf("%d", sweep.AutoGenned.String)
		// log.Info("Auto: ", sweep.AutoGenned, sweep.MatchID, err)
		for i := range sweep.AutoGenned {
			if int(sweep.AutoGenned[i]) == int(1) {
				sweep.AutoGen = true
				AutoGenned = append(AutoGenned, sweep)
			} else {
				sweep.AutoGen = false
				Challenges = append(Challenges, sweep)
			}
		}
	}

	data := TemplateData{
		Title: 			 "Season 8 Sweeps",
		AutoGens:        AutoGenned,
		Challenges:      Challenges, 
		TotalSweeps:     totalSweeps,
	}
	httpServeTemplate(w, "sweep", data)
}
