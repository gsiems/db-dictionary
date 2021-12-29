package view

import (
	"strings"

	m "github.com/gsiems/db-dictionary/model"
	t "github.com/gsiems/db-dictionary/template"
)

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

/*
   <h2>Tables that display potential oddness</h2>

       <th>Table</th>
       <th>No PK</th>
       <th>No indices</th>
       <th>Duplicate indices</th>
       <th>Only one column</th>
       <th>No data</th>
       <th>No relationships</th>
       <th>Denormalized?</th>

   <h2>Columns that display potential oddness</h2>

       <th>Table</th>
       <th>Column</th>
       <th>Nullable and part of a unique constraint</th>
       <th>Nullable and part of a unique index</th>
       <th>Nullable with a default value</th>
       <th>Defaults to NULL or 'NULL'</th>


*/

func makeTablePages(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {
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
			sortColumns(context.Columns)

			// Constraints
			switch strings.ToUpper(context.TableType) {
			case "TABLE":
				tmplt.AddSectionHeader("Constraints")
				tmplt.AddSnippet("TableConstraintsHeader")

				// primary key
				context.PrimaryKeys = md.FindPrimaryKeys(vs.Name, vt.Name)
				if len(context.PrimaryKeys) > 0 {
					tmplt.AddSnippet("TablePrimaryKey")
				}

				// check constraints
				context.CheckConstraints = md.FindCheckConstraints(vs.Name, vt.Name)
				if len(context.CheckConstraints) > 0 {
					tmplt.AddSnippet("TableCheckConstraints")
					sortCheckConstraints(context.CheckConstraints)
				}

				// unique constraints
				context.UniqueConstraints = md.FindUniqueConstraints(vs.Name, vt.Name)
				if len(context.UniqueConstraints) > 0 {
					tmplt.AddSnippet("TableUniqueConstraints")
					sortUniqueConstraints(context.UniqueConstraints)
				}

				tmplt.AddSnippet("TableConstraintsFooter")
			}

			// Indices
			switch strings.ToUpper(context.TableType) {
			case "TABLE", "MATERIALIZED VIEW":
				tmplt.AddSectionHeader("Indices")
				context.Indexes = md.FindIndexes(vs.Name, vt.Name)
				if len(context.Indexes) > 0 {
					tmplt.AddSnippet("TableIndexes")
					sortIndexes(context.Indexes)
				}
			}

			// Foreign Keys
			switch strings.ToUpper(context.TableType) {
			case "TABLE":
				tmplt.AddSectionHeader("Foreign Keys")
				context.ParentKeys = md.FindParentKeys(vs.Name, vt.Name)
				context.ChildKeys = md.FindChildKeys(vs.Name, vt.Name)

				if len(context.ParentKeys) > 0 || len(context.ChildKeys) > 0 {
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
		}
	}

	return err
}
