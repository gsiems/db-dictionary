package sqlite

import (
	"database/sql"

	m "github.com/gsiems/db-dictionary/model"
)

func DatabaseInfo(db *m.DB) (d m.Database, err error) {

	d.DatabaseName, err = databaseName(db)
	if err != nil {
		return d, err
	}

	d.Encoding, err = databaseEncoding(db)
	if err != nil {
		return d, err
	}

	/*

	   type Database struct {
	   	DatabaseName  sql.NullString `db:"database_name"  json:"databaseName"`
	   	Owner         sql.NullString `db:"owner"          json:"owner"`
	   	Encoding      sql.NullString `db:"encoding"       json:"encoding"`
	   	VersionNumber sql.NullString `db:"version_number" json:"versionNumber"`
	   	Comments      sql.NullString `db:"comments"       json:"comments"`
	   }




	   sub sql_dbms_ver {
	       my $dbh = shift;
	       return $dbh->FETCH('sqlite_version');
	   }

	*/

	return d, nil
}

func databaseName(db *m.DB) (sql.NullString, error) {

	var v sql.NullString
	rows, err := db.Query("select file from pragma_database_list where seq = 0")
	if err != nil {
		return v, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&v)
	}
	return v, err

}

func databaseEncoding(db *m.DB) (sql.NullString, error) {

	var v sql.NullString
	rows, err := db.Query("select encoding from pragma_encoding")
	if err != nil {
		return v, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&v)
	}
	return v, err

}
