// @SubApi API [/api]
package main

import (
	"encoding/json"
	"flag"
	"strings"
	"fmt"
	// "math"

	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"

	"regexp"
)

var (
	staticContent = flag.String("staticPath", "../public/swagger-ui", "Path to folder with Swagger UI")
	apiurl        = flag.String("api", "http://127.0.0.1", "The base path URI of the API service")
)

// @Title API
// @Description Lists all APIs
// @Accept plain
// @Produce json
// @Success 200 {object} json
// @Failure 404 {object} APIError "No Events Found"
// @Router /api [get]

func httpAPI(c *gin.Context) {
	// Local variables

	w := c.Writer

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(json.RawMessage(apiDescriptionsJson["api"]))
	
	if err != nil {
		log.Error("Couldn't generate JSON")
		return
	}

	w.Write(jsonData)

}

// @Title Events
// @Description Lists all events found by name
// @Accept plain
// @Produce json
// @Success 200 {object} models.ReturnedTable
// @Failure 404 {object} APIError "No Events Found"
// @Router /api/event [get]
func httpEventDocAPI(c *gin.Context) {
	// Local variables
	w := c.Writer

	returnedTables := make([]models.ReturnedTable, 0)

	foundTables, err := db.Tables.GetTables()
	if err != nil {
		log.Error("Could not get tables: ", err)
	}

	for _, tname := range foundTables {
		var rTable models.ReturnedTable
		rTable.EventName = tname.TableName.String
		if tname.PrettyName.Valid {
			rTable.PrettyName = tname.PrettyName.String
		} else {
			rTable.PrettyName = tname.TableName.String
		}
		returnedTables = append(returnedTables, rTable)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.MarshalIndent(returnedTables, "", "\t")
	if err != nil {
		log.Error("Couldn't generate JSON")
		//w.Write([]byte("Please search for a user"))
		return
	}
	w.Write(jsonData)

}

// @Title Event Listing
// @Description Lists everything found for the event
// @Accept plain
// @Produce json
// @Param event	path	string	true	"Event Name"
// @Success 200 {object} models.Event
// @Failure 404 {object} APIError "Event not found"
// @Router /api/event/{event} [get]
// httpEventAPI gets listings for the events
func httpEventAPI(c *gin.Context) {
	// Local variables

	w := c.Writer
	var foundEvent models.Event
	// Set the header
	w.Header().Set("Content-Type", "application/json")

	// Parse the event name from the URL
	event := c.Params.ByName("event")
	if event == "" {
		w.Write([]byte("{\"Error\": \"No event found\"}"))
		return
	}
	groups, err := regexp.MatchString("ndwc*", event)
	if err != nil {
		log.Error("Something funky  happened with regex")
	}
	if groups == true {
		foundEvent, _ = db.EventAPI.GetEventInfoGroups(event)
	} else {
		foundEvent, _ = db.EventAPI.GetEventInfo(event)
	}
	if err != nil {
		log.Error("Couldn't get event info: ", err)
	}
	season, err := regexp.MatchString("season_*", event)
	if err != nil {
		log.Error("Regex is bad")
	}

	editedEvent := models.Event{
		EventName: foundEvent.EventName,
	}
	var parts []models.Participant

	if season == true {
		for _, participant := range foundEvent.Participants {
			part := models.Participant{
				DiscordUsername: participant.DiscordUsername,
				TwitchUsername:  participant.TwitchUsername,
				EventWins:       participant.EventWins,
				EventLosses:     participant.EventLosses,
				EventPoints:     participant.EventPoints,
				EventPlayed:     participant.EventPlayed,
			}
			if event == "season_7" {
				part.TierName = s7tiers[strings.ToLower(participant.TwitchUsername)]
			}
			parts = append(parts, part)
		}
		editedEvent.Participants = parts
		foundEvent = editedEvent
	}

	jsonData, err := json.MarshalIndent(foundEvent, "", "\t")
	if err != nil {
		log.Error("Couldn't generate JSON")
		//w.Write([]byte("Please search for a user"))
		return
	}

	w.Write(jsonData)
}

// @Title Team Results Listing
// @Description Lists everything found for the season 7 teams
// @Accept plain
// @Produce json
// @Success 200 {object} models.Result
// @Failure 404 {object} APIError "Nothing found"
// @Router /api/teamresults [get]
// httpTeamAPI gets the team results and returns json
func httpTeamAPI(c *gin.Context) {
	w := c.Writer

	w.Header().Set("Content-Type", "application/json")

	results, err := db.TeamAPI.GetResults()
	if err != nil {
		log.Error("Couldn't get team info, ", err)
	}

	jsonData, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		log.Error("Couldn't parse results: ", err)
		return
	}
	w.Write(jsonData)
}

// @Title S Results Listing
// @Description Lists all racers from season 8 specifically
// @Accept plain
// @Produce json
// @Param event	path	string	true	"Event Name"
// @Success 200 {object} models.Event
// @Failure 404 {object} APIError "Nothing found"
// @Router /api/s [get]
// httpSAPI gets the s8 results for manipulating
func httpSAPI(c *gin.Context) {
	// Local variables

	w := c.Writer
	var foundEvent models.Event
	// Set the header
	w.Header().Set("Content-Type", "application/json")

	// Strictly for season 8
	event := "season_8"

	foundEvent, err := db.EventAPI.GetEventInfo(event)

	if err != nil {
		log.Error("Couldn't get event info: ", err)
	}

	season, err := regexp.MatchString("season_*", event)
	if err != nil {
		log.Error("Regex is bad")
	}

	editedEvent := models.Event{
		EventName: foundEvent.EventName,
	}
	var parts []models.Participant

	if season == true {
		for _, participant := range foundEvent.Participants {
			part := models.Participant{
				DiscordUsername: participant.DiscordUsername,
				TwitchUsername:  participant.TwitchUsername,
				EventWins:       participant.EventWins,
				EventLosses:     participant.EventLosses,
				EventPoints:     participant.EventPoints,
				EventPlayed:     participant.EventPlayed,
			}

			parts = append(parts, part)
		}
		editedEvent.Participants = parts
		foundEvent = editedEvent
	}

	jsonData, err := json.MarshalIndent(foundEvent, "", "\t")
	if err != nil {
		log.Error("Couldn't generate JSON")
		//w.Write([]byte("Please search for a user"))
		return
	}

	w.Write(jsonData)
}

// @Title S8 Sweeps Results 
// @Description Lists all sweeps on autogenned stuff from season 8 specifically
// @Accept plain
// @Produce json
// @Param event	path	string	true	"Event Name"
// @Success 200 {object} models.Event
// @Failure 404 {object} APIError "Nothing found"
// @Router /api/s [get]
// httpSAPI gets the s8 results for manipulating
func httpSweepsAPI(c *gin.Context) {
	// Local variables

	w := c.Writer

	// Set the header
	w.Header().Set("Content-Type", "application/json")

	sweepsResults, err := db.EventAPI.GetSweepsInfo()

	if err != nil {
		log.Error("Couldn't get event info: ", err)
	}


	jsonData, err := json.MarshalIndent(sweepsResults, "", "\t")
	if err != nil {
		log.Error("Couldn't generate JSON")
		//w.Write([]byte("Please search for a user"))
		return
	}

	w.Write(jsonData)
}

func httpUsersAPI(c *gin.Context) {
	// Local variables
	// w := c.Writer

	// userName := c.Params.ByName("user")

	// if userName == "" {
	// 	w.Write([]byte("{\"Error\": \"No username found\"}"))
	// 	return
	// }
	// foundUser, err := db.EventAPI.GetUserRaces(userName)
	// if err != nil {
	// 	log.Error("Couldn't get user, ", userName, " info: ", err)
	// }

	
	// jsonData, err := json.MarshalIndent(foundUser, "", "\t")
	// if err != nil {
	// 	log.Error("Couldn't generate JSON")
	// 	return
	// }

	// w.Write([]byte('HI'))
}

func httpUserAPI(c *gin.Context) {
	// Local variables
	w := c.Writer

	userName := c.Params.ByName("user")

	if userName == "" {
		w.Write([]byte("{\"Error\": \"No username found\"}"))
		return
	}
	FoundUser, err := db.EventAPI.GetUserRaces(userName)
	if err != nil {
		log.Error("Couldn't get user, ", userName, " info: ", err)
	}

	for i := 0; i < len(FoundUser); i++ {
		w := float64(FoundUser[i].RaceTime)/float64(100)
		ms := int((float64(FoundUser[i].RaceTime/100) - w) * -100)
		s := ((FoundUser[i].RaceTime/100) % 60)
		m := ((FoundUser[i].RaceTime/(100*60)) % 60)
		h := ((FoundUser[i].RaceTime/(100*60*60)) % 24)
		var fTime string
		if h > 0 {
			fTime += fmt.Sprintf("%02d:", h)
		}
		fTime += fmt.Sprintf("%02d:%02d.%02d", m, s, ms)
		FoundUser[i].RaceTimeF = fTime
		for j := range FoundUser[i].AutoGenFlag {
			if int(FoundUser[i].AutoGenFlag[j]) == int(1) {
				FoundUser[i].IsAutoGen = true
			}
		}
	}

	jsonData, err := json.MarshalIndent(FoundUser, "", "\t")
	if err != nil {
		log.Error("Couldn't generate JSON")
		return
	}

	w.Write(jsonData)
}
