package view

import (
	"html/template"
	"os"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
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

			var pageParts []string

			pageParts = append(pageParts, pageHeader(2))
			pageParts = append(pageParts, tpltTableHead(context.TableType))

			// Columns
			pageParts = append(pageParts, sectionHeader("Columns"))
			pageParts = append(pageParts, tpltTableColumns(context.TableType))
			context.Columns = md.FindColumns(vs.Name, vt.Name)
			sortColumns(context.Columns)

			// Constraints
			switch strings.ToUpper(context.TableType) {
			case "TABLE":
				pageParts = append(pageParts, sectionHeader("Constraints"))
				pageParts = append(pageParts, tpltTableConstraintsHeader())

				// primary key
				context.PrimaryKeys = md.FindPrimaryKeys(vs.Name, vt.Name)
				if len(context.PrimaryKeys) > 0 {
					pageParts = append(pageParts, tpltTablePrimaryKey())
				}

				// check constraints
				context.CheckConstraints = md.FindCheckConstraints(vs.Name, vt.Name)
				if len(context.CheckConstraints) > 0 {
					pageParts = append(pageParts, tpltTableCheckConstraints())
					sortCheckConstraints(context.CheckConstraints)
				}

				// unique constraints
				context.UniqueConstraints = md.FindUniqueConstraints(vs.Name, vt.Name)
				if len(context.UniqueConstraints) > 0 {
					pageParts = append(pageParts, tpltTableUniqueConstraints())
					sortUniqueConstraints(context.UniqueConstraints)
				}

				pageParts = append(pageParts, tpltTableConstraintsFooter())
			}

			// Indices
			switch strings.ToUpper(context.TableType) {
			case "TABLE", "MATERIALIZED VIEW":
				pageParts = append(pageParts, sectionHeader("Indices"))
				context.Indexes = md.FindIndexes(vs.Name, vt.Name)
				if len(context.Indexes) > 0 {
					pageParts = append(pageParts, tpltTableIndexes())
					sortIndexes(context.Indexes)
				}
			}

			// Foreign Keys
			switch strings.ToUpper(context.TableType) {
			case "TABLE":
				pageParts = append(pageParts, sectionHeader("Foreign Keys"))
				context.ParentKeys = md.FindParentKeys(vs.Name, vt.Name)
				context.ChildKeys = md.FindChildKeys(vs.Name, vt.Name)

				if len(context.ParentKeys) > 0 || len(context.ChildKeys) > 0 {
					if len(context.ParentKeys) > 0 {
						pageParts = append(pageParts, tpltTableParentKeys())
						sortForeignKeys(context.ParentKeys)
					}
					if len(context.ChildKeys) > 0 {
						pageParts = append(pageParts, tpltTableChildKeys())
						sortForeignKeys(context.ParentKeys)
					}
				}
			}

			// Dependencies
			pageParts = append(pageParts, sectionHeader("Dependencies"))

			context.Dependencies = md.FindDependencies(vs.Name, vt.Name)
			if len(context.Dependencies) > 0 {
				pageParts = append(pageParts, tpltTableDependencies())
				sortDependencies(context.Dependencies)
			}

			context.Dependents = md.FindDependents(vs.Name, vt.Name)
			if len(context.Dependents) > 0 {
				pageParts = append(pageParts, tpltTableDependents())
				sortDependencies(context.Dependents)
			}

			//switch strings.ToUpper(context.TableType) {
			//case "TABLE":
			// Foreign Data Wrapper
			// tpltTableFDW
			//}

			// Query
			if len(context.Query) > 0 {
				pageParts = append(pageParts, sectionHeader("Query"))
				pageParts = append(pageParts, tpltTableQuery())
			}

			pageParts = append(pageParts, pageFooter())

			templates, err := template.New("doc").Funcs(template.FuncMap{
				"safeHTML": func(u string) template.HTML { return template.HTML(u) },
			}).Parse(strings.Join(pageParts, ""))
			if err != nil {
				return err
			}

			dirName := md.OutputDir + "/" + vs.Name + "/tables/"
			_, err = os.Stat(dirName)
			if os.IsNotExist(err) {
				err = os.Mkdir(dirName, 0744)
				if err != nil {
					return err
				}
			}

			outfile, err := os.Create(dirName + "/" + vt.Name + ".html")
			if err != nil {
				return err
			}
			defer outfile.Close()

			err = templates.Lookup("doc").Execute(outfile, context)
			if err != nil {
				return err
			}
		}
	}

	return err
}
