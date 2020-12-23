package pg

import (
	m "github.com/gsiems/db-dictionary/model"
)

func DatabaseInfo(db *m.DB) (m.Database, error) {

	q := `
SELECT d.datname AS database_name,
        pg_catalog.pg_get_userbyid ( d.datdba ) AS owner,
        pg_catalog.pg_encoding_to_char ( d.encoding ) AS encoding,
        pg_catalog.version () AS version_number,
        pg_catalog.shobj_description ( d.oid, 'pg_database' ) AS comments
    FROM pg_catalog.pg_database d
    WHERE d.datname = pg_catalog.current_database ()
`

	return db.DatabaseInfo(q)
}
