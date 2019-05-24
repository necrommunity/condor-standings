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


func httpUser(c *gin.Context) {
	w := c.Writer
	
	data := TemplateData{
		Title: 			 "Season 8 Users",
		
	}
	httpServeTemplate(w, "user", data)

}



func httpUserInfo(c *gin.Context) {
	w := c.Writer

	userName := c.Params.ByName("user")
	
	if userName == "" {
		w.Write([]byte("{\"Error\": \"No username found\"}"))
		return
	}

	url := fmt.Sprintf("https://some.pizza/api/user/%s",  userName)
	log.Info(url)
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
	var UserMatches []models.Match
	jsonErr := json.Unmarshal(body, &UserMatches)
	if jsonErr != nil {
		log.Error(jsonErr)
	}

	data := TemplateData{
		Title: 			 userName + " Match Info",
		UserMatchInfo:   UserMatches,
	}

	httpServeTemplate(w, "users", data)
}
