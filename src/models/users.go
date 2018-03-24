package models

import (
	"database/sql"
)

// Users struct captures all users found
type Users struct{}

// UserAccount gets each row for all profiles
type UserAccount struct {
	DiscordID       sql.NullInt64
	DiscordUsername sql.NullString
}

// FoundUserNDWC struct stores the info needed for NDWC events
type FoundUserNDWC struct {
	DiscordID       sql.NullInt64
	DiscordUsername sql.NullInt64
	NDWCGroup       sql.NullString
}

// GetUsers gets player data to populate the player's profile page
func (*Users) GetUsers() ([]UserAccount, int, error) {
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

func (*Tables) GetRacersForEvent(eventName) ([]FoundUserNDWC, Totalerr)
