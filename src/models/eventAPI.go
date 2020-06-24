package models

import (
	
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

// EventAPI struct captures all users found
type EventAPI struct{}

// Participant gets each row for all profiles
type Participant struct {
	DiscordUsername string `json:"discordUsername"`
	TwitchUsername  string `json:"twitchUsername"`
	RacerID         int    `json:"racerID"`
	EventWins       int    `json:"eventWins"`
	EventLosses     int    `json:"eventLosses"`
	EventPoints     int    `json:"eventPoints"`
	EventPlayed     int    `json:"eventPlayed"`
	GroupName       string `json:"groupName"`
	TierName        string `json:"tierName"`
}

// Event holds the event information
type Event struct {
	EventName    string `json:"eventName"`
	Participants []Participant
}

// Sweep holds the data for Sweeps page
type Sweep struct {
	MatchID	    int    `json:"match_id"`
	Racer1	    string `json:"racer1"`
	Racer2 	    string `json:"racer2"`
	StartTime	time.Time    `json:"startTime"`
	Cawmentator sql.NullString `json:"cawmentator"`
	VODLink	    sql.NullString `json:"vodLink"`
	AutoGenned	[]byte   `json:"autogenned"`
	AutoGen     bool	`json:"autoGen"`
	Racer1Wins	int	   `json:"racer1Wins"`
	Racer2Wins	int	   `json:"racer2Wins"`
}

// Match holds the data for a basic event match
type Match struct {
	RaceID int `json:"raceID"`
	RaceTime int `json:"raceTime"`
	RaceTimeF string `json:"raceTimeF"`
	RaceSeed int `json:"raceSeed"`
	RaceWinner int `json:"raceWinner"`
	Racer1ID int `json:"racer1ID"`
	Racer1Name string `json:"racer1Name"`
	Racer2ID int `json:"racer2ID"`
	Racer2Name string `json:"racer2Name"`
	RaceVod sql.NullString `json:"raceVod"`
	AutoGenFlag []byte `json:"autoGenFlag"`
	IsAutoGen bool `json:"isAutoGen"`
}

// UserNames holds data for both search and display names for racers
type UserNames struct {
	TwitchName string `json:"twitchName"`
	DiscordName string `json:"discordName"`
}

// SeedStats holds data for the seed info page
type SeedStats struct {
	SeedNum int `json:"seedNum"`
	AvgTime int `json:"avgTime"`
	MinTime int `json:"minTime"`
	MaxTime int `json:"maxTime"`
	DisplayAvgTime string `json:"displayAvgTime"`
	DisplayMinTime string `json:"displayMinTime"`
	DisplayMaxTime string `json:"displayMaxTime"`

	NumOfSeeds int `json:"numOfSeeds"`
}

// GetEventInfoGroups gets all needed information about a specific event
func (*EventAPI) GetEventInfoGroups(eventName string) (Event, error) {
	var rows *sql.Rows
	TheEvent := Event{EventName: eventName, Participants: nil}
	if v, err := db.Query(`
		SELECT
			u.discord_name AS Username,
			u.twitch_name as tUsername,
			u.user_id as racerId,
			(SELECT	COUNT(rs.winner_id)
				FROM ` + eventName + `.race_summary rs
				WHERE rs.winner_id = u.user_id
			) as wins,
			(SELECT	COUNT(rs.winner_id)
				FROM ` + eventName + `.race_summary rs
				WHERE rs.loser_id = u.user_id
			) as losses,
			e.group
		FROM
			` + eventName + `.entrants e
				LEFT JOIN
			necrobot.users u ON u.user_id = e.user_id
		WHERE
			u.discord_id IS NOT NULL
			AND u.discord_name IS NOT NULL
			AND u.twitch_name IS NOT NULL
			AND u.user_id IS NOT NULL
		GROUP BY u.user_id
		ORDER BY e.group, wins desc, losses asc, tUsername asc;
		`); err == sql.ErrNoRows {
		return TheEvent, nil
	} else if err != nil {
		return TheEvent, err
	} else {
		rows = v
	}
	defer rows.Close()

	// Find all the users and stick them into the event structure
	participants := make([]Participant, 0)
	for rows.Next() {
		var participant Participant

		if err := rows.Scan(
			&participant.DiscordUsername,
			&participant.TwitchUsername,
			&participant.RacerID,
			&participant.EventWins,
			&participant.EventLosses,
			&participant.GroupName,
		); err != nil {
			return TheEvent, err
		}
		participants = append(participants, participant)
	}

	TheEvent.Participants = participants

	return TheEvent, nil
}

// GetEventInfo gets specific event info
func (*EventAPI) GetEventInfo(eventName string) (Event, error) {
	var rows *sql.Rows
	TheEvent := Event{EventName: eventName, Participants: nil}
	if v, err := db.Query(`
		SELECT
		u.discord_name AS Username,
		u.twitch_name as tUsername,
		u.user_id as racerId,
		(SELECT	COUNT(*)
			FROM ` + eventName + `.race_summary rs
			WHERE rs.winner_id = u.user_id
		) as wins,
		(SELECT	COUNT(*)
			FROM ` + eventName + `.race_summary rs
			WHERE rs.loser_id = u.user_id
		) as losses
		
		FROM
			` + eventName + `.entrants e
				LEFT JOIN
			necrobot.users u ON u.user_id = e.user_id
		WHERE
			u.discord_id IS NOT NULL
			AND u.discord_name IS NOT NULL
			AND u.twitch_name IS NOT NULL
			AND u.user_id IS NOT NULL
		GROUP BY u.user_id
		ORDER BY wins desc, losses asc, tUsername asc;
		`); err == sql.ErrNoRows {
		return TheEvent, nil
	} else if err != nil {
		return TheEvent, err
	} else {
		rows = v
	}
	defer rows.Close()

	// Find all the users and stick them into the event structure
	participants := make([]Participant, 0)
	for rows.Next() {
		var participant Participant
		if err := rows.Scan(
			&participant.DiscordUsername,
			&participant.TwitchUsername,
			&participant.RacerID,
			&participant.EventWins,
			&participant.EventLosses,
		); err != nil {
			return TheEvent, err
		}
		participants = append(participants, participant)
	}

	TheEvent.Participants = participants

	return TheEvent, nil
}

// GetSweepsInfo gets specific event info
func (*EventAPI) GetSweepsInfo() ([]Sweep, error) {
	var rows *sql.Rows
	var sweepsInfo []Sweep
	if v, err := db.Query(`
		SELECT 
		  match_id,
		  racer_1_name,
		  racer_2_name,
		  scheduled_time,
		  cawmentator_name,
		  vod,
		  autogenned,
		  racer_1_wins,
		  racer_2_wins 
		FROM
		  season_8.match_info 
		WHERE
		  completed = 1
		`); err == sql.ErrNoRows {
		return sweepsInfo, nil
	} else if err != nil {
		return sweepsInfo, err
	} else {
		rows = v
	}
	defer rows.Close()

	// Find all the users and stick them into the event structure
	for rows.Next() {
		var sweeps Sweep
		if err := rows.Scan(
			&sweeps.MatchID,
			&sweeps.Racer1,
			&sweeps.Racer2,
			&sweeps.StartTime,
			&sweeps.Cawmentator,
			&sweeps.VODLink,
			&sweeps.AutoGenned,
			&sweeps.Racer1Wins,
			&sweeps.Racer2Wins,
		); err != nil {
			return sweepsInfo, err
		}
		sweepsInfo = append(sweepsInfo, sweeps)
	}

	return sweepsInfo, nil
}

// GetUsers takes a username and gets the relevant data
func (*EventAPI) GetUsers(eventName string) ([]UserNames, error) {
	var rows *sql.Rows
	var userList []UserNames
	if v, err := db.Query(`
	SELECT 
    	u.twitch_name, u.discord_name
	FROM
    	` + eventName + `.entrants e
    LEFT JOIN
    	necrobot.users u ON e.user_id = u.user_id
	`); err == sql.ErrNoRows {
		return userList, nil
	} else if err != nil {
		return userList, err
	} else {
		rows = v
	}
	defer rows.Close()

	for rows.Next() {
		var user UserNames
		if err := rows.Scan(
		 &user.TwitchName,
		 &user.DiscordName,

		); err != nil {
			return userList, err
		}
		userList = append(userList, user)
	}

	return userList, nil
}

// GetUserRaces gets the races for a given user
func (*EventAPI) GetUserRaces(userName string, eventName string) ([]Match, error) {
	var rows *sql.Rows
	var matchInfo []Match
	if v, err := db.Query(`
		SELECT 
				rr.race_id,
				mr.winner,
				rr.time,
				r.seed,
				m.racer_1_id,
				m.racer_2_id,
				m.vod,
				(SELECT 
								e.twitch_name
						FROM
								necrobot.users e
						WHERE
								e.user_id = m.racer_1_id) AS racer_1_name,
				(SELECT 
								e.twitch_name
						FROM
								necrobot.users e
						WHERE
								e.user_id = m.racer_2_id) AS racer_2_name,
				m.autogenned
		FROM
				`+ eventName +`.match_races mr
						LEFT JOIN
				`+ eventName +`.races r ON r.race_id = mr.race_id
						LEFT JOIN
				`+ eventName +`.matches m ON m.match_id = mr.match_id
						LEFT JOIN
				`+ eventName +`.race_runs rr ON rr.race_id = r.race_id
						LEFT JOIN
				necrobot.users e ON e.user_id = rr.user_id
		WHERE
				e.twitch_name = ?
		GROUP BY rr.race_id
	`, userName); err == sql.ErrNoRows {
		return matchInfo, nil
	} else if err != nil {
		return matchInfo, err
	} else {
		rows = v
	}
	defer rows.Close()

	for rows.Next() {
		var matches Match
		if err := rows.Scan(
		 &matches.RaceID,
		 &matches.RaceWinner,
		 &matches.RaceTime,
		 &matches.RaceSeed,
		 &matches.Racer1ID,
		 &matches.Racer2ID,
		 &matches.RaceVod,
		 &matches.Racer1Name,
		 &matches.Racer2Name,
		 &matches.AutoGenFlag,
		); err != nil {
			return matchInfo, err
		}
		matchInfo = append(matchInfo, matches)
	}

	return matchInfo, nil
}

// GetSeedStats gets seed stats for Condor X
func (*EventAPI) GetSeedStats() ([]SeedStats, error) {
	var rows *sql.Rows
	var stats []SeedStats
	if v, err := db.Query(`
		SELECT
			substr(r.seed, 1,1) as seed_num,
			ROUND(AVG(rr.time)) as seed_avg,
			ROUND(MIN(rr.time)) as seed_min,
			ROUND(MAX(rr.time)) as seed_max,
			COUNT(rr.time) as seed_count
		FROM
			condor_x.races r
		LEFT JOIN
			condor_x.race_runs rr
			ON rr.race_id = r.race_id
		LEFT JOIN
			condor_x.match_races mr
			ON mr.race_id = r.race_id
		LEFT JOIN
			condor_x.matches m
			ON m.match_id = mr.match_id
		LEFT JOIN
			condor_x.leagues l
			ON l.league_tag = m.league_tag
		WHERE
			r.seed <> 0
			AND rr.rank = 1
			AND l.league_name = "cad"
		GROUP BY
			seed_num
		ORDER BY
	seed_num asc, time asc;
		`); err == sql.ErrNoRows {
		return stats, nil
	} else if err != nil {
		return stats, err
	} else {
		rows = v
	}
	defer rows.Close()

	
	for rows.Next() {
		var stat SeedStats
		if err := rows.Scan(
			&stat.SeedNum,
			&stat.AvgTime,
			&stat.MinTime,
			&stat.MaxTime,
			&stat.NumOfSeeds,
		); err != nil {
			return stats, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}