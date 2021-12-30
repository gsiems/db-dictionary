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

type oddTable struct {
	TableName        string
	NoPK             string //
	NoIndices        string //
	DuplicateIndices string //
	OneColumn        string //
	NoData           string //
	NoRelationships  string //
	Denormalized     string
}

type oddColumn struct {
	TableName       string
	ColumnName      string
	NullUnique      string
	NullWithDefault string
	NullAsDefault   string
}

type oddness struct {
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	OddTables     []oddTable
	OddColumns    []oddColumn
}

func makeTablePages(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		//otm := make(map[string]oddTable)
		//oc := make(map[string]oddColumn)

		o := oddness{
			Title:         "Odd things - " + md.Alias + "." + vs.Name,
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

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

			var otm []string

			if vt.RowCount == 0 {
				otm = append(otm, "NoData")
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
			} else {
				otm = append(otm, "OneColumn")
			}

			// Constraints
			switch strings.ToUpper(context.TableType) {
			case "TABLE":
				tmplt.AddSectionHeader("Constraints")
				tmplt.AddSnippet("TableConstraintsHeader")

				// primary key
				context.PrimaryKeys = md.FindPrimaryKeys(vs.Name, vt.Name)
				if len(context.PrimaryKeys) > 0 {
					tmplt.AddSnippet("TablePrimaryKey")
				} else {
					otm = append(otm, "NoPK")
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

					dupChk := make(map[string]int)
					for _, idx := range context.Indexes {
						dupChk[idx.IndexColumns]++
					}
					for _, kount := range dupChk {
						if kount > 1 {
							otm = append(otm, "DuplicateIndices")
						}
					}
				} else {
					otm = append(otm, "NoIndices")
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
				} else {
					otm = append(otm, "NoRelationships")
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

			switch strings.ToUpper(context.TableType) {
			case "TABLE":

				if len(otm) > 0 {
					ot := oddTable{
						TableName: vt.Name,
					}

					for _, v := range otm {
						switch v {
						case "NoPK":
							ot.NoPK = "X"
						case "NoIndices":
							ot.NoIndices = "X"
						case "DuplicateIndices":
							ot.DuplicateIndices = "X"
						case "OneColumn":
							ot.OneColumn = "X"
						case "NoData":
							ot.NoData = "X"
						case "NoRelationships":
							ot.NoRelationships = "X"
						case "Denormalized":
							ot.Denormalized = "X"
						}
					}
					o.OddTables = append(o.OddTables, ot)
				}
			}
		}

		// Create odd things page
		var tmplt t.T
		tmplt.AddPageHeader(1, md)
		tmplt.AddSnippet("OddHeader")
		tmplt.AddSectionHeader("Tables that display potential oddness")
		if len(o.OddTables) > 0 {
			tmplt.AddSnippet("OddTables")
		} else {
			tmplt.AddSnippet("      <p><b>No table oddities were extracted for this schema.</b></p>")
		}

		tmplt.AddSectionHeader("Columns that display potential oddness")
		if len(o.OddColumns) > 0 {
			tmplt.AddSnippet("OddColumns")
		} else {
			tmplt.AddSnippet("      <p><b>No column oddities were extracted for this schema.</b></p>")
		}

		tmplt.AddPageFooter()

		dirName := md.OutputDir + "/" + vs.Name
		err = tmplt.RenderPage(dirName, "odd-things", o)
		if err != nil {
			return err
		}
	}

	return err
}
