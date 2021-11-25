package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gsiems/db-dictionary/util"
	e "github.com/gsiems/go-db-meta/engine/pg"
	m "github.com/gsiems/go-db-meta/model"
)

var (
	showVersion bool
	version     = "0.1"
	base        string
	dbName      string
	debug       bool
	host        string
	port        string
	quiet       bool
	schemas     string
	userName    string
	xclude      string
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `usage: pg_dictionary [flags]

Database connection flags

  -db      The database to connect to.

  -host    The hostname that the database is on. Defaults to localhost.

  -port    The port that the database is listening on. Defaults to 5432.

  -user    The username to connect as. Defaults to the OS user.


Extract database/schema(s) DDL flags

  -b       The base directory to write the generated results to.
           Overrides the BASE_DIR environment variable. Defaults to the
           current directory.

  -s       The comma separated list of schemas to extract.

  -x       The comma separated list of schemas to exclude.
           Ignored if the -s flag is supplied.

Other flags

  -debug

  -q       Quiet mode. Do not print any error messages.

  -version Display the version information

`)
	}
	flag.BoolVar(&debug, "debug", false, "")
	flag.BoolVar(&quiet, "q", false, "")
	flag.BoolVar(&showVersion, "version", false, "")
	flag.StringVar(&dbName, "db", "", "")
	flag.StringVar(&host, "host", "", "")
	flag.StringVar(&port, "port", "", "")
	flag.StringVar(&userName, "user", "", "")
	flag.StringVar(&base, "b", "", "")
	flag.StringVar(&schemas, "s", "", "")
	flag.StringVar(&xclude, "x", "", "")

	flag.Parse()

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var c m.ConnectInfo
	c.Username = userName
	c.Host = host
	c.Port = port
	c.DbName = dbName
	//c.Debug = debug

	db, err := e.OpenDB(&c)
	util.FailOnErr(quiet, err)
	defer func() {
		if cerr := db.CloseDB(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	catalog, err := e.CurrentCatalog(db)
	util.FailOnErr(quiet, err)
	catalogName := catalog.CatalogName.String
	catalogOwner := catalog.CatalogOwner.String
	catalogComment := catalog.Comment.String

	fmt.Printf("Catalog Name: %q\n", catalogName)
	fmt.Printf("Catalog Owner: %q\n", catalogOwner)
	fmt.Printf("Catalog Comment: %q\n", catalogComment)


	schemata, err := e.Schemata(db, schemas, xclude)
	util.FailOnErr(quiet, err)

	for _, schema := range schemata {
		fmt.Printf("    Schema Name: %q\n", schema.SchemaName.String)
		fmt.Printf("    Schema Owner: %q\n", schema.SchemaOwner.String)
		fmt.Printf("    Schema Comment: %q\n", schema.Comment.String)
	}


	tables, err := e.Tables(db, "")
	util.FailOnErr(quiet, err)
	for _, table := range tables {
		fmt.Printf("        Table Schema: %q\n", table.TableSchema.String)
		fmt.Printf("        Table Name: %q\n", table.TableName.String)
		fmt.Printf("        Table Owner: %q\n", table.TableOwner.String)
		fmt.Printf("        Table Type: %q\n", table.TableType.String)
		fmt.Printf("        Table Comment: %q\n", table.Comment.String)
	}

	columns, err := e.Columns(db, "", "")
	util.FailOnErr(quiet, err)
	for _, column := range columns {
		fmt.Printf("        Table Schema: %q\n", column.TableSchema.String)
		fmt.Printf("        Table Name: %q\n", column.TableName.String)
		fmt.Printf("        Column Name: %q\n", column.ColumnName.String)
		fmt.Printf("        Data Type: %q\n", column.DataType.String)
		fmt.Printf("        Column Comment: %q\n", column.Comment.String)
	}



}

/*

func (db *DB) CurrentCatalog(q string) (Catalog, error) {

type Catalog struct {
	CatalogName             sql.NullString `json:"catalogName"`
	CatalogOwner            sql.NullString `json:"catalogOwner"`
	DefaultCharacterSetName sql.NullString `json:"defaultCharacterSetName"`
	DBMSVersion             sql.NullString `json:"dbmsVersion"`
	Comment                 sql.NullString `json:"comment"`
}

func (db *DB) Schemata(q, nclude, xclude string) ([]Schema, error) {

type Schema struct {
	CatalogName                sql.NullString `json:"catalogName"`
	SchemaName                 sql.NullString `json:"schemaName"`
	SchemaOwner                sql.NullString `json:"schemaOwner"`
	DefaultCharacterSetCatalog sql.NullString `json:"defaultCharacterSetCatalog"`
	DefaultCharacterSetSchema  sql.NullString `json:"defaultCharacterSetSchema"`
	DefaultCharacterSetName    sql.NullString `json:"defaultCharacterSetName"`
	Comment                    sql.NullString `json:"comment"`
}

func Domains(db *m.DB, schema string) ([]m.Domain, error) {

type Domain struct {
	DomainCatalog sql.NullString `json:"domainCatalog"`
	DomainSchema  sql.NullString `json:"domainSchema"`
	DomainName    sql.NullString `json:"domainName"`
	DomainOwner   sql.NullString `json:"domainOwner"`
	DataType      sql.NullString `json:"dataType"`
	DomainDefault sql.NullString `json:"domainDefault"`
	Comment       sql.NullString `json:"comment"`
}

func (db *DB) Tables(q, tableSchema string) ([]Table, error) {

type Table struct {
	TableCatalog   sql.NullString `json:"tableCatalog"`
	TableSchema    sql.NullString `json:"tableSchema"`
	TableName      sql.NullString `json:"tableName"`
	TableOwner     sql.NullString `json:"tableOwner"`
	TableType      sql.NullString `json:"tableType"`
	RowCount       sql.NullInt64  `json:"rowCount"`
	Comment        sql.NullString `json:"comment"`
	ViewDefinition sql.NullString `json:"viewDefinition"`
}

func Columns(db *m.DB, tableSchema, tableName string) ([]m.Column, error) {

type Column struct {
	TableCatalog    sql.NullString `json:"tableCatalog"`
	TableSchema     sql.NullString `json:"tableSchema"`
	TableName       sql.NullString `json:"tableName"`
	ColumnName      sql.NullString `json:"columnName"`
	OrdinalPosition sql.NullInt32  `json:"ordinalPosition"`
	ColumnDefault   sql.NullString `json:"columnDefault"`
	IsNullable      sql.NullString `json:"isNullable"`
	DataType        sql.NullString `json:"dataType"`
	DomainCatalog   sql.NullString `json:"domainCatalog"`
	DomainSchema    sql.NullString `json:"domainSchema"`
	DomainName      sql.NullString `json:"domainName"`
	Comment         sql.NullString `json:"comment"`
}

func (db *DB) PrimaryKeys(q, tableSchema, tableName string) ([]PrimaryKey, error) {

type PrimaryKey struct {
	TableCatalog      sql.NullString `json:"tableCatalog"`
	TableSchema       sql.NullString `json:"tableSchema"`
	TableName         sql.NullString `json:"tableName"`
	ConstraintName    sql.NullString `json:"constraintName"`
	ConstraintColumns sql.NullString `json:"constraintColumns"`
	ConstraintStatus  sql.NullString `json:"constraintStatus"`
	Comment           sql.NullString `json:"comment"`
}

func (db *DB) ReferentialConstraints(q, tableSchema, tableName string) ([]ReferentialConstraint, error) {

type ReferentialConstraint struct {
	TableCatalog      sql.NullString `json:"tableCatalog"`
	TableSchema       sql.NullString `json:"tableSchema"`
	TableName         sql.NullString `json:"tableName"`
	TableColumns      sql.NullString `json:"tableColumns"`
	ConstraintName    sql.NullString `json:"constraintName"`
	RefTableCatalog   sql.NullString `json:"refTableCatalog"`
	RefTableSchema    sql.NullString `json:"refTableSchema"`
	RefTableName      sql.NullString `json:"refTableName"`
	RefTableColumns   sql.NullString `json:"refTableColumns"`
	RefConstraintName sql.NullString `json:"refConstraintName"`
	MatchOption       sql.NullString `json:"matchOption"`
	UpdateRule        sql.NullString `json:"updateRule"`
	DeleteRule        sql.NullString `json:"deleteRule"`
	IsEnforced        sql.NullString `json:"isEnforced"`
	//is_deferrable
	//initially_deferred
	Comment sql.NullString `json:"comment"`
}

func CheckConstraints(db *m.DB, tableSchema, tableName string) ([]m.CheckConstraint, error) {

type CheckConstraint struct {
	TableCatalog   sql.NullString `json:"tableCatalog"`
	TableSchema    sql.NullString `json:"tableSchema"`
	TableName      sql.NullString `json:"tableName"`
	ConstraintName sql.NullString `json:"constraintName"`
	CheckClause    sql.NullString `json:"checkClause"`
	Status         sql.NullString `json:"status"`
	Comment        sql.NullString `json:"comment"`
}

func Indexes(db *m.DB, tableSchema, tableName string) ([]m.Index, error) {

type Index struct {
	IndexCatalog sql.NullString `json:"indexCatalog"`
	IndexSchema  sql.NullString `json:"indexSchema"`
	IndexName    sql.NullString `json:"indexName"`
	IndexType    sql.NullString `json:"indexType"`
	IndexColumns sql.NullString `json:"indexColumns"`
	TableCatalog sql.NullString `json:"tableCatalog"`
	TableSchema  sql.NullString `json:"tableSchema"`
	TableName    sql.NullString `json:"tableName"`
	IsUnique     sql.NullString `json:"isUnique"`
	Comment      sql.NullString `json:"comment"`
}



func (db *DB) Types(q, tableSchema string) ([]Type, error) {

type Type struct {
	TypeCatalog sql.NullString `json:"typeCatalog"`
	TypeSchema  sql.NullString `json:"typeSchema"`
	TypeName    sql.NullString `json:"typeName"`
	TypeOwner   sql.NullString `json:"typeOwner"`
	//DataType    sql.NullString `json:"dataType"`
	Comment sql.NullString `json:"comment"`
}


type Sequence struct {
	SequenceCatalog sql.NullString `json:"sequenceCatalog"`
	SequenceSchema  sql.NullString `json:"sequenceSchema"`
	SequenceName    sql.NullString `json:"sequenceName"`
	SequenceOwner   sql.NullString `json:"sequenceOwner"`
	DataType        sql.NullString `json:"dataType"`
	MinimumValue    sql.NullString `json:"minimumValue"`
	MaximumValue    sql.NullString `json:"maximumValue"`
	Increment       sql.NullString `json:"increment"`
	CycleOption     sql.NullString `json:"cycleOption"`
	StartValue      sql.NullString `json:"startValue"`
	Comment         sql.NullString `json:"comment"`
}

// Sequences returns a slice of Sequences for the (schema) parameter
func (db *DB) Sequences(q, sequenceSchema string) ([]Sequence, error) {



*/
