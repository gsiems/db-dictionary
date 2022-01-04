package view

import (
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
	t "github.com/gsiems/db-dictionary-core/template"
)

// tableView contains the data used for generating a table information page
type tableView struct {
	Title             string
	DBMSVersion       string
	DBName            string
	DBComment         string
	SchemaName        string
	SchemaComment     string
	TableName         string
	TableComment      string
	TableType         string
	RowCount          int64
	TmspGenerated     string
	Query             string
	Columns           []m.Column
	Indexes           []m.Index
	PrimaryKeys       []m.PrimaryKey
	ParentKeys        []m.ForeignKey
	ChildKeys         []m.ForeignKey
	CheckConstraints  []m.CheckConstraint
	UniqueConstraints []m.UniqueConstraint
	Dependencies      []m.Dependency
	Dependents        []m.Dependency
}

// makeTablePages marshals the data needed for, and then creates, the table (or
// view, or materialized view) information pages for schemas
func makeTablePages(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		oddThings := initOddThings(md, vs)

		for _, vt := range md.FindTables(vs.Name) {

			context := tableView{
				Title:         "Table - " + md.Alias + "." + vs.Name + "." + vt.Name,
				TmspGenerated: md.TmspGenerated,
				DBName:        md.Name,
				DBComment:     md.Comment,
				SchemaName:    vs.Name,
				SchemaComment: vs.Comment,
				TableName:     vt.Name,
				TableType:     vt.TableType,
				RowCount:      vt.RowCount,
				Query:         vt.Query,
				TableComment:  vt.Comment,
			}

			var tmplt t.T
			tmplt.AddPageHeader(2, md)
			tmplt.AddTableHead(context.TableType)

			// Columns
			tmplt.AddSectionHeader("Columns")
			tmplt.AddTableColumns(context.TableType)
			context.Columns = md.FindColumns(vs.Name, vt.Name)
			if len(context.Columns) > 1 {
				sortColumns(context.Columns)
			}

			// Constraints
			switch strings.ToUpper(context.TableType) {
			case "TABLE", "BASE TABLE":
				context.PrimaryKeys = md.FindPrimaryKeys(vs.Name, vt.Name)
				context.CheckConstraints = md.FindCheckConstraints(vs.Name, vt.Name)
				context.UniqueConstraints = md.FindUniqueConstraints(vs.Name, vt.Name)

				if len(context.PrimaryKeys) > 0 || len(context.CheckConstraints) > 0 || len(context.UniqueConstraints) > 0 {
					tmplt.AddSectionHeader("Constraints")
					tmplt.AddSnippet("TableConstraintsHeader")

					// primary key
					if len(context.PrimaryKeys) > 0 {
						tmplt.AddSnippet("TablePrimaryKey")
					}

					// check constraints
					if len(context.CheckConstraints) > 0 {
						tmplt.AddSnippet("TableCheckConstraints")
						sortCheckConstraints(context.CheckConstraints)
					}

					// unique constraints
					if len(context.UniqueConstraints) > 0 {
						tmplt.AddSnippet("TableUniqueConstraints")
						sortUniqueConstraints(context.UniqueConstraints)
					}

					tmplt.AddSnippet("TableConstraintsFooter")
				}
			}

			// Indices
			switch strings.ToUpper(context.TableType) {
			case "TABLE", "BASE TABLE", "MATERIALIZED VIEW":
				context.Indexes = md.FindIndexes(vs.Name, vt.Name)
				if len(context.Indexes) > 0 {
					tmplt.AddSectionHeader("Indices")
					tmplt.AddSnippet("TableIndexes")
					sortIndexes(context.Indexes)
				}
			}

			// Foreign Keys
			switch strings.ToUpper(context.TableType) {
			case "TABLE", "BASE TABLE":
				context.ParentKeys = md.FindParentKeys(vs.Name, vt.Name)
				context.ChildKeys = md.FindChildKeys(vs.Name, vt.Name)

				if len(context.ParentKeys) > 0 || len(context.ChildKeys) > 0 {
					tmplt.AddSectionHeader("Foreign Keys")
					if len(context.ParentKeys) > 0 {
						tmplt.AddSnippet("TableParentKeys")
						sortForeignKeys(context.ParentKeys)
					}
					if len(context.ChildKeys) > 0 {
						tmplt.AddSnippet("TableChildKeys")
						sortForeignKeys(context.ParentKeys)
					}
				}
			}

			// Dependencies
			context.Dependencies = md.FindDependencies(vs.Name, vt.Name)
			context.Dependents = md.FindDependents(vs.Name, vt.Name)

			if len(context.Dependencies) > 0 || len(context.Dependents) > 0 {
				tmplt.AddSectionHeader("Dependencies")
			}

			if len(context.Dependencies) > 0 {
				tmplt.AddSnippet("TableDependencies")
				sortDependencies(context.Dependencies)
			}

			if len(context.Dependents) > 0 {
				tmplt.AddSnippet("TableDependents")
				sortDependencies(context.Dependents)
			}

			//switch strings.ToUpper(context.TableType) {
			//case "TABLE":
			// Foreign Data Wrapper
			// tpltTableFDW
			//}

			// Query
			if len(context.Query) > 0 {
				tmplt.AddSectionHeader("Query")
				tmplt.AddSnippet("TableQuery")
			}

			tmplt.AddPageFooter()

			dirName := md.OutputDir + "/" + vs.Name + "/tables/"

			err = tmplt.RenderPage(dirName, vt.Name, context)
			if err != nil {
				return err
			}
			oddThings.checkOddThings(&context)
		}
		err := oddThings.makeOddnessPage()
		if err != nil {
			return err
		}
	}

	return err
}
