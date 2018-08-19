// @SubApi API [/api]
package main

import (
	"encoding/json"
	"flag"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
	"github.com/sillypears/condor-standings/src/models"

	"regexp"
)

var (
	staticContent = flag.String("staticPath", "../public/swagger-ui", "Path to folder with Swagger UI")
	apiurl        = flag.String("api", "http://127.0.0.1", "The base path URI of the API service")
)

// ReturnedTables creates a struct for json output
type ReturnedTables struct {
	EventNames []string `json:"eventNames"`
}

func httpAPI(c *gin.Context) {
	// Local variables

	w := c.Writer

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(json.RawMessage(apiDescriptionsJson))
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
// @Success 200 {object} ReturnedTables
// @Failure 404 {object} APIError "No Events Found"
// @Router /api/event [get]
func httpEventDocAPI(c *gin.Context) {
	// Local variables
	w := c.Writer

	returnedTables := make([]string, 0)

	foundTables, err := db.Tables.GetTables()
	if err != nil {
		log.Error("Could not get tables: ", err)
	}

	for _, tname := range foundTables {
		returnedTables = append(returnedTables, tname.TableName.String)
	}

	var TablesReturned ReturnedTables
	TablesReturned.EventNames = returnedTables

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.MarshalIndent(TablesReturned, "", "\t")
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
// @Success 200 {object} Event
// @Failure 404 {object} APIError "Event not found"
// @Router /api/event/{event} [get]
// httpEventAPI gets listings for the events
func httpEventAPI(c *gin.Context) {
	// Local variables

	// <rant>
	// This is fucking terrible and you should never actually do this but circumstances
	// have lead to things being forced like this and I can't do much about it
	// so here we are. Live and learn children.
	// </rant>
	tiers := map[string]string{
		strings.ToLower("abu__yazan"):              "Crystal",
		strings.ToLower("incnone"):                 "Crystal",
		strings.ToLower("jackofgames"):             "Crystal",
		strings.ToLower("mayaundefined"):           "Crystal",
		strings.ToLower("mudjoe2"):                 "Crystal",
		strings.ToLower("pancelor"):                "Crystal",
		strings.ToLower("ptrevordactyl"):           "Crystal",
		strings.ToLower("RoyalGoof"):               "Crystal",
		strings.ToLower("spootybiscuit"):           "Crystal",
		strings.ToLower("Squega"):                  "Crystal",
		strings.ToLower("Staekk"):                  "Crystal",
		strings.ToLower("thedarkfreaack"):          "Crystal",
		strings.ToLower("Tufwfo"):                  "Crystal",
		strings.ToLower("ARTQ"):                    "Obsidian",
		strings.ToLower("biggiemac42"):             "Obsidian",
		strings.ToLower("cyber_1"):                 "Obsidian",
		strings.ToLower("kingtorture"):             "Obsidian",
		strings.ToLower("mantasmbl"):               "Obsidian",
		strings.ToLower("moyuma"):                  "Obsidian",
		strings.ToLower("Paratroopa1"):             "Obsidian",
		strings.ToLower("ratata_ratata"):           "Obsidian",
		strings.ToLower("Ratracing"):               "Obsidian",
		strings.ToLower("raviolinguini"):           "Obsidian",
		strings.ToLower("reijigazpacho"):           "Obsidian",
		strings.ToLower("Revalize"):                "Obsidian",
		strings.ToLower("Siveure"):                 "Obsidian",
		strings.ToLower("supervillain_joe"):        "Obsidian",
		strings.ToLower("tetel__"):                 "Obsidian",
		strings.ToLower("yuka34"):                  "Obsidian",
		strings.ToLower("alex42918"):               "Titanium",
		strings.ToLower("Ancalagor"):               "Titanium",
		strings.ToLower("bastet222"):               "Titanium",
		strings.ToLower("call_me_kaye"):            "Titanium",
		strings.ToLower("chef_mayhem"):             "Titanium",
		strings.ToLower("definitely_not_him"):      "Titanium",
		strings.ToLower("duneaught"):               "Titanium",
		strings.ToLower("firebrde"):                "Titanium",
		strings.ToLower("flygluffet"):              "Titanium",
		strings.ToLower("Lucoa"):                   "Titanium",
		strings.ToLower("Saakas0206"):              "Titanium",
		strings.ToLower("SailorMint"):              "Titanium",
		strings.ToLower("seanpwolf"):               "Titanium",
		strings.ToLower("Uniowen"):                 "Titanium",
		strings.ToLower("WuffWuff"):                "Titanium",
		strings.ToLower("yamiramiz"):               "Titanium",
		strings.ToLower("arborelia"):               "Blood",
		strings.ToLower("brumekuroi"):              "Blood",
		strings.ToLower("cohomerlogist"):           "Blood",
		strings.ToLower("dsmidna"):                 "Blood",
		strings.ToLower("ekimekim"):                "Blood",
		strings.ToLower("EpicSuccess"):             "Blood",
		strings.ToLower("fiverfiverone"):           "Blood",
		strings.ToLower("Jamblar"):                 "Blood",
		strings.ToLower("justsparkyyes"):           "Blood",
		strings.ToLower("Kailaria"):                "Blood",
		strings.ToLower("Kova46"):                  "Blood",
		strings.ToLower("Kyakh"):                   "Blood",
		strings.ToLower("MegaMissingn0"):           "Blood",
		strings.ToLower("Minhs2"):                  "Blood",
		strings.ToLower("professionaltwitchtroll"): "Blood",
		strings.ToLower("rachelsatx"):              "Blood",
		strings.ToLower("Rotomington"):             "Blood",
		strings.ToLower("Slimo"):                   "Blood",
		strings.ToLower("wow_tomato"):              "Blood",
		strings.ToLower("Zin"):                     "Blood",
	}
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
				EventPoints:     participant.EventPoints,
				EventPlayed:     participant.EventPlayed,
				TierName:        tiers[strings.ToLower(participant.TwitchUsername)],
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
