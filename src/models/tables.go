package models

import (
	"database/sql"
)

// Tables struct captures all users found
type Tables struct{}

// FoundTable struct captures all the event tables
type FoundTable struct {
	TableName sql.NullString
}

// GetTables finds all the database names
func (*Tables) GetTables() ([]FoundTable, error) {
	var rows *sql.Rows
	if v, err := db.Query(`
    SELECT
      DISTINCT(table_schema)
    FROM
      INFORMATION_SCHEMA.TABLES
    WHERE
      table_schema not in ('information_schema', 'necrobot' )
    ORDER BY
      table_schema ASC
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
		); err != nil {
			return nil, err
		}

		tables = append(tables, row)
	}
	return tables, nil
}
