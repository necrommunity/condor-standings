package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	// "strings"
	"fmt"

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
	var AutoGennedSweeps []models.Sweep
	var Challenges []models.Sweep
	var ChallengesSweeps []models.Sweep

	
	for _, sweep := range SweepsResults {
		
		for i := range sweep.AutoGenned {
			
			if int(sweep.AutoGenned[i]) == int(1) {
				if sweep.Racer1Wins == 3 || sweep.Racer2Wins == 3 {
					sweep.AutoGen = true
					AutoGennedSweeps = append(AutoGennedSweeps, sweep)
				}
				AutoGenned = append(AutoGenned, sweep)
			} else {
				if sweep.Racer1Wins == 3 || sweep.Racer2Wins == 3 {
					sweep.AutoGen = false
					ChallengesSweeps = append(ChallengesSweeps, sweep)
				}
				Challenges = append(Challenges, sweep)
			}
		}
	}

	
	totalAGSweeps := len(AutoGennedSweeps)
	totalAG := len(AutoGenned)
	totalCSweeps := len(ChallengesSweeps)
	totalC := len(Challenges)
	totalMatches := len(SweepsResults)
	totalSweepMatches := totalCSweeps + totalAGSweeps

	// // Find wins based on name
	// for _, parts := range ReturnedEvent.Participants {
	// 	if strings.HasPrefix(parts.TwitchUsername, "s") {
	// 		sWins += parts.EventWins
	// 		sParticipants = append(sParticipants, parts)
	// 	} else {
	// 		nonSWins += parts.EventWins
	// 		nonSParticipants = append(nonSParticipants, parts)
	// 	}
	// }
	
	agSweepsPerc := fmt.Sprintf("%.2f", (float64(totalAGSweeps) / float64(totalSweepMatches) * 100))
	cSweepsPerc := fmt.Sprintf("%.2f", (float64(totalCSweeps) / float64(totalSweepMatches) * 100))
	agPerc := fmt.Sprintf("%.2f", (float64(totalAG) / float64(totalMatches) * 100))
	cPerc := fmt.Sprintf("%.2f", (float64(totalC) / float64(totalMatches) * 100))

	data := TemplateData{
		Title: 			 "Season 8 Sweeps",
		AutoGens:        AutoGenned,
		AutoGensLen:     len(AutoGenned),
		AutoGensSweepLen: len(AutoGennedSweeps),
		TotalAG:         totalAG,
		Challenges:      Challenges, 
		TotalC:          totalC,
		ChallengesLen:   len(Challenges),
		ChalSweepLen:    len(ChallengesSweeps),
		TotalMatches:    totalMatches,
		TotalSweeps:     totalSweepMatches,
		AGSweepsPerc:    agSweepsPerc,
		CSweepsPerc:     cSweepsPerc,
		CPerc:           cPerc,
		AGPerc:          agPerc,
// 		AllSWinsPerc:	 sWinsPerc,
// 		AllNonSWins: 	 nonSWins,
// 		AllNonSWinsPerc: nonSWinsPerc,
// 		AllSParts:		 sParticipants,
// 		AllNonSParts:	 nonSParticipants,
	}
	httpServeTemplate(w, "sweep", data)
}
