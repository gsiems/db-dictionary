package model

import (
	"database/sql"
)

// Database struct for the basic database information
type Database struct {
	DatabaseName  sql.NullString `db:"database_name"  json:"databaseName"`
	Owner         sql.NullString `db:"owner"          json:"owner"`
	Encoding      sql.NullString `db:"encoding"       json:"encoding"`
	VersionNumber sql.NullString `db:"version_number" json:"versionNumber"`
	Comments      sql.NullString `db:"comments"       json:"comments"`
}

func (db *DB) DatabaseInfo(q string) (d Database, err error) {

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&d.DatabaseName,
			&d.Owner,
			&d.Encoding,
			&d.VersionNumber,
			&d.Comments,
		)
	}

	return
}
