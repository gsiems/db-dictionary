package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	m "github.com/gsiems/db-dictionary/model"
)

// OpenDB opens a database connection and returns a DB reference.
func OpenDB(c *m.ConnectInfo) (*m.DB, error) {

	dsn := fmt.Sprintf("user=%s dbname=%s host=%s port=%s", c.Username, c.DbName, c.Host, c.Port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &m.DB{db}, db.Ping()
}
