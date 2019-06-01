package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"
)

//  --------------------------------
//  The landing page for the website
//  --------------------------------

func httpS(c *gin.Context) {
	// Local variables
	w := c.Writer

	sWins := 0
	nonSWins := 0
	var sParticipants []models.Participant 
	var nonSParticipants []models.Participant

	url := "https://some.pizza/api/s"
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
	var ReturnedEvent models.Event
	jsonErr := json.Unmarshal(body, &ReturnedEvent)
	if jsonErr != nil {
		log.Error(jsonErr)
	}
	
	// Find wins based on name
	for _, parts := range ReturnedEvent.Participants {
		if strings.HasPrefix(strings.ToLower(parts.TwitchUsername), "s") {
			sWins += parts.EventWins
			sParticipants = append(sParticipants, parts)
		} else {
			nonSWins += parts.EventWins
			nonSParticipants = append(nonSParticipants, parts)
		}
	}
	totalWins := sWins + nonSWins
	sWinsPerc := fmt.Sprintf("%.2f", (float64(sWins) / float64(totalWins) * 100))
	nonSWinsPerc := fmt.Sprintf("%.2f", (float64(nonSWins) / float64(totalWins) * 100))
	log.Info(totalWins, sWinsPerc)
	data := TemplateData{
		Title: 			 "Season 8 ESSSSSSS",
		SEvent: 		 ReturnedEvent,
		AllSWins: 		 sWins,
		AllSWinsPerc:	 sWinsPerc,
		AllNonSWins: 	 nonSWins,
		AllNonSWinsPerc: nonSWinsPerc,
		AllSParts:		 sParticipants,
		AllNonSParts:	 nonSParticipants,
	}
	httpServeTemplate(w, "s", data)
}
