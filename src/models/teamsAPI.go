package models

import (
	"database/sql"
)

// TeamAPI struct captures all info found
type TeamAPI struct{}

// Result gets each row for all profiles
type Result struct {
	Racer1 string `json:"racer1"`
	Team1  string `json:"team1"`
	Racer2 string `json:"racer2"`
	Team2  string `json:"team2"`
	Winner int    `json:"winner"`
}

// GetResultsAllTeams gets all needed information about a specific event
func (*TeamAPI) GetResultsAllTeams() ([]Result, error) {
	var rows *sql.Rows
	var results []Result
	if v, err := db.Query(`
		SELECT
		    u1.twitch_name AS racer1,
		    s7t1.team AS team1,
		    u2.twitch_name AS racer2,
		    s7t2.team AS team2,
		    mr.winner
		FROM
		    season_7.matches m
		        LEFT JOIN
		    necrobot.users u1 ON u1.user_id = m.racer_1_id
		        LEFT JOIN
		    necrobot.users u2 ON u2.user_id = m.racer_2_id
		        LEFT JOIN
		    season_7.season_7_teams s7t1 ON s7t1.user_id = m.racer_1_id
		        LEFT JOIN
		    season_7.season_7_teams s7t2 ON s7t2.user_id = m.racer_2_id
		        LEFT JOIN
		    season_7.match_races mr ON mr.match_id = m.match_id
		        LEFT JOIN
		    season_7.race_runs rr ON rr.race_id = mr.race_id
		WHERE
		    rr.level = - 2
		    
		GROUP BY mr.race_id
		`); err == sql.ErrNoRows {
		return results, nil
	} else if err != nil {
		return results, err
	} else {
		rows = v
	}
	defer rows.Close()

	results = make([]Result, 0)
	for rows.Next() {
		var result Result
		if err := rows.Scan(
			&result.Racer1,
			&result.Team1,
			&result.Racer2,
			&result.Team2,
			&result.Winner,
		); err != nil {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}

// GetResults gets all needed information about a specific event
func (*TeamAPI) GetResults() ([]Result, error) {
	var rows *sql.Rows
	var results []Result
	if v, err := db.Query(`
		SELECT
		    u1.twitch_name AS racer1,
		    s7t1.team AS team1,
		    u2.twitch_name AS racer2,
		    s7t2.team AS team2,
		    mr.winner
		FROM
		    season_7.matches m
		        LEFT JOIN
		    necrobot.users u1 ON u1.user_id = m.racer_1_id
		        LEFT JOIN
		    necrobot.users u2 ON u2.user_id = m.racer_2_id
		        LEFT JOIN
		    season_7.season_7_teams s7t1 ON s7t1.user_id = m.racer_1_id
		        LEFT JOIN
		    season_7.season_7_teams s7t2 ON s7t2.user_id = m.racer_2_id
		        LEFT JOIN
		    season_7.match_races mr ON mr.match_id = m.match_id
		        LEFT JOIN
		    season_7.race_runs rr ON rr.race_id = mr.race_id
		WHERE
		    rr.level = - 2
		    
		GROUP BY mr.race_id
		`); err == sql.ErrNoRows {
		return results, nil
	} else if err != nil {
		return results, err
	} else {
		rows = v
	}
	defer rows.Close()

	results = make([]Result, 0)
	for rows.Next() {
		var result Result
		if err := rows.Scan(
			&result.Racer1,
			&result.Team1,
			&result.Racer2,
			&result.Team2,
			&result.Winner,
		); err != nil {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}
