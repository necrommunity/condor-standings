package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"
)

//  --------------------------------
//  The landing page for the website
//  --------------------------------

func httpHome(c *gin.Context) {
	// Local variables
	w := c.Writer

	url := "https://wow.freepizza.how/api/event"
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
	var EventTables []models.ReturnedTable
	jsonErr := json.Unmarshal(body, &EventTables)
	if jsonErr != nil {
		log.Error(jsonErr)
	}

	data := TemplateData{
		Title:       "Home",
		FoundTables: EventTables,
	}
	httpServeTemplate(w, "home", data)
}
