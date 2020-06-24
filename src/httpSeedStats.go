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

func httpSeedStats(c *gin.Context) {
	// Local variables
	w := c.Writer


	url := "https://some.pizza/api/seedstats"
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

	var SeedStats []models.SeedStats
	jsonErr := json.Unmarshal(body, &SeedStats)
	if jsonErr != nil {
		log.Error(jsonErr)
	}
	var totalSeeds int
	for i, times := range SeedStats {
		SeedStats[i].DisplayAvgTime = formatTime(times.AvgTime)
		SeedStats[i].DisplayMinTime = formatTime(times.MinTime)
		SeedStats[i].DisplayMaxTime = formatTime(times.MaxTime)
		// log.Info(SeedStats[i])
		totalSeeds += times.NumOfSeeds
	}

	data := TemplateData{
		Title: 			 "Condor X Seed Stats",
		SeedStatData:	 SeedStats,
		TotalSeeds:		 totalSeeds,
	}

	httpServeTemplate(w, "seeds", data)
	
}
