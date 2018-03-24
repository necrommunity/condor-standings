package models

import (
	"database/sql"
)

// Users struct captures all users found
type Users struct{}

// UserAccounts gets each row for all profiles
type UserAccounts struct {
	Username sql.NullString
}

// GetUsers gets player data to populate the player's profile page
func (*Users) GetUsers() ([]UserAccounts, int, error) {
	var rows *sql.Rows
	if v, err := db.Query(`
		SELECT
			u.discord_name,
		FROM
			users u
		ORDER BY
			u.discord_name ASC

	`); err == sql.ErrNoRows {
		return nil, 0, nil
	} else {
		rows = v
	}
	defer rows.Close()

	// Iterate over the user profile results
	userAccounts := make([]UserAccounts, 0)
	for rows.Next() {
		var row UserAccounts
		if err := rows.Scan(
			&row.Username,
		); err != nil {
			return nil, 0, err
		}

		userAccounts = append(userAccounts, row)
	}

	// Find total amount of users
	var allUsers int
	if err := db.QueryRow(`
		SELECT count(id)
		FROM users
	`).Scan(&allUsers); err != nil {
		return nil, 0, err
	}

	return userAccounts, allUsers, nil
}
