package models

import (
	"database/sql"
)

// Tables struct captures all users found
type Tables struct{}

// FoundTable struct captures all the event tables
type FoundTable struct {
	TableName  sql.NullString
	PrettyName sql.NullString
}

// GetTables finds all the database names
func (*Tables) GetTables() ([]FoundTable, error) {
	var rows *sql.Rows
	if v, err := db.Query(`
    SELECT
			DISTINCT(ist.table_schema),
			nl.league_name
    FROM
			INFORMATION_SCHEMA.TABLES ist
		LEFT JOIN
			necrobot.leagues_old nl
				ON nl.schema_name = ist.table_schema
    WHERE
      table_schema not in ('information_schema', 'necrobot' )
    ORDER BY
      create_time DESC
  `); err == sql.ErrNoRows {
		return nil, nil
	} else {
		rows = v
	}
	defer rows.Close()

	// Iterate over the user profile results
	tables := make([]FoundTable, 0)
	for rows.Next() {
		var row FoundTable
		if err := rows.Scan(
			&row.TableName,
			&row.PrettyName,
		); err != nil {
			return nil, err
		}

		tables = append(tables, row)
	}
	return tables, nil
}
