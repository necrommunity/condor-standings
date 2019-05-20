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

// Why in the world does this take 8 seconds to run D:
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

// Why in the world does this take 8 seconds to run D:
// GetEventInfo gets specific event info
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

/*
// GetEventInfo gets specific event info
func (*EventAPI) GetEventInfo(eventName string) (Event, error) {
	var rows *sql.Rows
	TheEvent := Event{EventName: eventName, Participants: nil}
	if v, err := db.Query(`
		SELECT
		    u.discord_name AS Username,
				u.twitch_name as tUsername,
		    SUM(CASE
		        WHEN rr.rank = 1 THEN 1
		        ELSE 0
		    END) AS Points,
		    COUNT(rr.user_id) AS Played
		FROM
		    ` + eventName + `.entrants e
		        LEFT JOIN
		    necrobot.users u ON u.user_id = e.user_id
		        LEFT JOIN
		    ` + eventName + `.race_runs rr ON rr.user_id = e.user_id
		WHERE
			u.discord_id IS NOT NULL
			AND u.discord_name IS NOT NULL
			AND u.twitch_name IS NOT NULL
		GROUP BY u.discord_name
		ORDER BY Points DESC , Played DESC
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
			&participant.EventPoints,
			&participant.EventPlayed,
		); err != nil {
			return TheEvent, err
		}

		participants = append(participants, participant)
	}

	TheEvent.Participants = participants

	return TheEvent, nil
}

// GetUsers gets player data to populate the player's profile page
/*func (*Users) GetEventInfo(eventName string) ([]UserAccount, int, error) {
	var rows *sql.Rows
	if v, err := db.Query(`
		SELECT
		    u.discord_id,
				u.discord_name,

		FROM
		    users u
		WHERE
		    u.discord_id IS NOT NULL
				AND u.discord_name IS NOT NULL
		ORDER BY u.discord_name ASC
	`); err == sql.ErrNoRows {
		return nil, 0, nil
	} else {
		rows = v
	}
	defer rows.Close()

	// Iterate over the user profile results
	userAccounts := make([]UserAccount, 0)
	for rows.Next() {
		var row UserAccount
		if err := rows.Scan(
			&row.DiscordID,
			&row.DiscordUsername,
		); err != nil {
			return nil, 0, err
		}

		userAccounts = append(userAccounts, row)
	}

	// Find total amount of users
	var allUsers int
	if err := db.QueryRow(`
		SELECT count(discord_id)
		FROM users
		WHERE discord_id IS NOT NULL
	`).Scan(&allUsers); err != nil {
		return nil, 0, err
	}

	return userAccounts, allUsers, nil
}

func (*Tables) GetRacersForEvent(eventName string) ([]FoundUserNDWC, int, error) {
	var rows *sql.Rows
	if v, err := db.Query(`
		SELECT
				u.discord_id,
				u.discord_name
		FROM
				users u
		WHERE
				u.discord_id IS NOT NULL
				AND u.discord_name IS NOT NULL
		ORDER BY u.discord_name ASC
	`); err == sql.ErrNoRows {
		return nil, 0, nil
	} else {
		rows = v
	}
	defer rows.Close()

	// Iterate over the user profile results
	userAccounts := make([]UserAccount, 0)
	for rows.Next() {
		var row UserAccount
		if err := rows.Scan(
			&row.DiscordID,
			&row.DiscordUsername,
		); err != nil {
			return nil, 0, err
		}

		userAccounts = append(userAccounts, row)
	}

}*/
