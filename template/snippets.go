package template

import (
	"fmt"
	"path"
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
)

func pageHeader(i int, md *m.MetaData) string {

	b := ""

  // resolve the css files
	var css []string
	if md.Cfg.CSSFiles != "" {
		x := strings.Split(md.Cfg.CSSFiles, ",")
		for _, v := range x {
			css = append(css, path.Base(v))
		}
	}
  // if no custom css specified then use the default
	if len(css) == 0 {
		css = append(css, "blues.css")
	}

	switch i {
	case 1, 2:
		ri := ""
		si := ""

		switch i {
		case 1:
			ri = "../"
			si = ""
		case 2:
			ri = "../../"
			si = "../"
		}

		dom := ""
		if len(md.Domains) > 0 {
			dom = `
      <a href="` + si + `domains.html">Domains</a>`
		}

		//class="active"

		var ci []string
		for _, v := range css {
			ci = append(ci, `    <link rel="stylesheet" href="`+ri+`css/`+v+`" type="text/css">`)
		}

		b = fmt.Sprintf(`<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
`+strings.Join(ci, "\n")+`
    <script type="text/javascript" src="`+ri+`js/filter.js"></script>
  </head>
  <body>
    <div id="topNav">
      <a href="`+ri+`index.html">Schemas</a>
      <a href="`+si+`columns.html">Columns</a>
      <a href="`+si+`constraints.html">Constraints</a>%s
      <a href="`+si+`tables.html">Tables</a>
      <a href="`+si+`odd-things.html">Odd things</a>
    </div>`, dom)

	default:

		var ci []string
		for _, v := range css {
			ci = append(ci, `    <link rel="stylesheet" href="css/`+v+`" type="text/css">`)
		}

		b = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
` + strings.Join(ci, "\n") + `
    <script type="text/javascript" src="js/filter.js"></script>
  </head>
  <body>
<!--    <div id="topNav">
      <a class="active" href="index.html">Schemas</a>
    </div> -->`
	}

	return b
}

func sectionHeader(s string) string {
	return fmt.Sprintf(`
      <h2>%s</h2>`, s)
}

func pageFooter() string {
	return `
    </div>
    <div id="pageFoot">Generated by <a href="https://github.com/gsiems/db-dictionary">db-dictionary</a></div>
    <br />
  </body>
</html>
`
}

func reportHead(s, t, rc bool) string {

	var schemaTxt string
	var tableTxt string
	var rowCount string

	if s {
		schemaTxt = `
        <div class="headingItem">Schema:</div><div class="headingItem">{{.SchemaName}}</div>{{if .SchemaComment}}
        <div class="headingItem"></div><div class="headingItem">{{.SchemaComment|safeHTML}}</div>{{end}}`
	}
	if t {
		tableTxt = `
        <div class="headingItem">Table:</div><div class="headingItem">{{.TableName}}</div>{{if .TableComment}}
        <div class="headingItem"></div><div class="headingItem">{{.TableComment|safeHTML}}</div>{{end}}`

	}
	if rc {
		rowCount = `
        <div class="headingItem">Row Count:</div><div class="headingItem">{{.RowCount}}</div>`
	}

	return `
    <div id="pageHead"><h1>{{.Title}}</h1>
      <div class="headingContainer">
        <div class="headingItem">Generated:</div><div class="headingItem">{{.TmspGenerated}}</div>{{if .DBMSVersion}}
        <div class="headingItem">Database Version:</div><div class="headingItem">{{.DBMSVersion}}</div>{{end}}
        <div class="headingItem">Database:</div><div class="headingItem">{{.DBName}}</div>{{if .DBComment}}
        <div class="headingItem"></div><div class="headingItem">{{.DBComment|safeHTML}}</div>{{end}}` + schemaTxt + tableTxt + rowCount + `
        <div class="headingItem">Filter:</div><div class="headingItem"><form id="filter-form" onsubmit="return false;"><input name="filter" id="filterBy" value="" maxlength="32" size="32" type="text" onkeyup="filterTables()"></form></div>
      </div>
    </div>`
}

func tpltSchemas() string {
	return reportHead(false, false, false) + `
    <div id="pageBody">
      <br/>
      <table width="100.0%" id="dataTable-schema" class="dataTable">
        <thead>
        <tr>
          <th>Schema</th>
          <th>Owner</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Schemas}}
          <tr>
            <td><a href="{{.Name}}/tables.html">{{.Name}}</a></td>
            <td>{{.Owner}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
}

func tpltSchemaTables() string {
	return reportHead(true, false, false) + `
    <div id="pageBody">
      <br/>
      <table width="100.0%" id="dataTable-tab" class="dataTable">
        <thead>
        <tr>
          <th>Table</th>
          <th>Owner</th>
          <th>Type</th>
          <th>Rows</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Tables}}
          <tr>
            <td><a href="tables/{{.Name}}.html">{{.Name}}</a></td>
            <td>{{.Owner}}</td>
            <td>{{.TableType}}</td>
            <td class="tcn">{{.RowCount}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
}

func tpltSchemaDomains() string {
	return reportHead(true, false, false) + `
    <div id="pageBody">
      <br/>
      <table width="100.0%" id="dataTable-dom" class="dataTable">
        <thead>
        <tr>
          <th>Name</th>
          <th>Data Type</th>
          <th>Default</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Domains}}
          <tr>
            <td>{{.Name}}</td>
            <td>{{.DataType}}</td>
            <td>{{.Default}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
}

func tpltSchemaColumns() string {
	return reportHead(true, false, false) + `
    <div id="pageBody">
      <br/>
      <table width="100.0%" id="dataTable-col" class="dataTable">
        <thead>
        <tr>
          <th>Table</th>
          <th>Column</th>
          <th>Position</th>
          <th>Data Type</th>
          <th>Nulls</th>
          <th>Default</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Columns}}
          <tr>
            <td><a href="tables/{{.TableName}}.html">{{.TableName}}</a></td>
            <td>{{.Name}}</td>
            <td class="tcn">{{.OrdinalPosition}}</td>
            <td>{{.DataType}}</td>
            <td class="tcc">{{.IsNullable|checkMark}}</td>
            <td>{{.Default}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
}

func tpltSchemaConstraintsHeader() string {
	return reportHead(true, false, false) + `
    <div id="pageBody">`
}

func tpltSchemaCheckConstraints() string {
	return `
      <table width="100.0%" id="dataTable-chk" class="dataTable">
        <thead>
        <tr>
          <th>Table</th>
          <th>Constraint</th>
          <th>Search Condition</th>
          <th>Status</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .CheckConstraints}}
          <tr>
            <td><a href="tables/{{.TableName}}.html">{{.TableName}}</a></td>
            <td>{{.Name}}</td>
            <td>{{.CheckClause}}</td>
            <td>{{.Status}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
}

func tpltSchemaUniqueConstraints() string {
	return `
      <table width="100.0%" id="dataTable-uniq" class="dataTable">
        <thead>
        <tr>
          <th>Table</th>
          <th>Constraint</th>
          <th>Columns</th>
          <th>Status</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .UniqueConstraints}}
          <tr>
            <td><a href="tables/{{.TableName}}.html">{{.TableName}}</a></td>
            <td>{{.Name}}</td>
            <td>{{.Columns}}</td>
            <td>{{.Status}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
}

func tpltSchemaFKConstraints() string {
	return `
      <table width="100.0%" id="dataTable-fk" class="dataTable">
        <thead>
        <tr>
          <th>Table</th>
          <th>Constraint</th>
          <th>Columns</th>
          <th>Is Indexed</th>
          <th>Referenced Table</th>
          <th>Referenced Columns</th>
          <th>On Update</th>
          <th>On Delete</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .ParentKeys}}
          <tr>
            <td><a href="tables/{{.TableName}}.html">{{.TableName}}</a></td>
            <td>{{.Name}}</td>
            <td>{{.TableColumns}}</td>
            <td class="tcc">{{.IsIndexed|checkMark}}</td>
            <td>{{.RefSchemaName}}.<a href="../{{.RefSchemaName}}/tables/{{.RefTableName}}.html">{{.RefTableName}}</a>
            <td>{{.RefTableColumns}}</td>
            <td>{{.UpdateRule}}</td>
            <td>{{.DeleteRule}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
}

func tpltTableHead(tabType string) string {

	switch tabType {
	case "TABLE", "MATERIALIZED VIEW":
		return reportHead(true, true, true) + `
    <div id="pageBody">`
	}

	return reportHead(true, true, false) + `
    <div id="pageBody">`

}

func tpltTableColumns(tabType string) string {
	switch tabType {
	case "VIEW", "MATERIALIZED VIEW":
		return `
      <table width="100.0%" id="dataTable-col" class="dataTable">
        <thead>
        <tr>
          <th>Column</th>
          <th>Position</th>
          <th>Data Type</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Columns}}
          <tr>
            <td>{{.Name}}</td>
            <td class="tcn">{{.OrdinalPosition}}</td>
            <td>{{.DataType}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
	default:
		return `
      <table width="100.0%" id="dataTable-col" class="dataTable">
        <thead>
        <tr>
          <th>Column</th>
          <th>Ordinal Position</th>
          <th>Data Type</th>
          <th>Nulls</th>
          <th>Default</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Columns}}
          <tr>
            <td>{{.Name}}</td>
            <td class="tcn">{{.OrdinalPosition}}</td>
            <td>{{.DataType}}</td>
            <td class="tcc">{{.IsNullable|checkMark}}</td>
            <td>{{.Default}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      <br />`
	}
}

func tpltTableConstraintsHeader() string {
	return `
      <table width="100.0%" id="dataTable-cons" class="dataTable">
        <thead>
        <tr>
          <th>Name</th>
          <th>Type</th>
          <th>Columns</th>
          <th>Search Condition</th>
          <th>Status</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>`
}

func tpltTableConstraintsFooter() string {
	return `
        </tbody>
      </table>
      <br />`
}

func tpltTableCheckConstraints() string {
	return `{{range .CheckConstraints}}
          <tr>
            <td class="tcnw">{{.Name}}</td>
            <td class="tcnw">Check</td>
            <td class="tcnw"></td>
            <td class="tcnw">{{.CheckClause}}</td>
            <td class="tcnw">{{.Status}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}`
}

func tpltTablePrimaryKey() string {
	return `{{range .PrimaryKeys}}
          <tr>
            <td class="tcnw">{{.Name}}</td>
            <td class="tcnw">Primary Key</td>
            <td class="tcnw">{{.Columns}}</td>
            <td class="tcnw"></td>
            <td class="tcnw"></td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}`
}

func tpltTableUniqueConstraints() string {
	return `{{range .UniqueConstraints}}
          <tr>
            <td class="tcnw">{{.Name}}</td>
            <td class="tcnw">Unique</td>
            <td class="tcnw">{{.Columns}}</td>
            <td class="tcnw"></td>
            <td class="tcnw"></td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}`
}

func tpltTableIndexes() string {
	return `
      <table width="100.0%" id="dataTable-idx" class="dataTable">
        <thead>
        <tr>
          <th>Name</th>
          <th>Columns</th>
          <th>Is Unique</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Indexes}}
          <tr>
            <td class="tcnw">{{.Name}}</td>
            <td class="tcnw">{{.IndexColumns}}</td>
            <td class="tcc">{{.IsUnique|checkMark}}</td>
            <td>{{.Comment|safeHTML}}</td>
          </tr>{{end}}
        </tbody>
      </table>
      <br />`
}

func tpltTableParentKeys() string {
	return `
      <p><b>Parents (references)</b></p>
      <table width="100.0%" id="dataTable-parent" class="dataTable">
        <thead>
        <tr>
          <th>Name</th>
          <th>Columns</th>
          <th>Is Indexed</th>
          <th>Referenced Table</th>
          <th>Referenced Columns</th>
          <th>On Update</th>
          <th>On Delete</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .ParentKeys}}
        <tr>
          <td class="tcnw">{{.Name}}</td>
          <td>{{.TableColumns}}</td>
          <td class="tcc">{{.IsIndexed|checkMark}}</td>
          <td class="tcnw">{{.RefSchemaName}}.<a href="../../{{.RefSchemaName}}/tables/{{.RefTableName}}.html">{{.RefTableName}}</a></td>
          <td>{{.RefTableColumns}}</td>
          <td>{{.UpdateRule}}</td>
          <td>{{.DeleteRule}}</td>
          <td>{{.Comment|safeHTML}}</td>
        </tr>{{end}}
        </tbody>
      </table>
      <br />`
}

func tpltTableChildKeys() string {
	return `
      <p><b>Children (referenced by)</b></p>
      <table width="100.0%" id="dataTable-child" class="dataTable">
        <thead>
        <tr>
          <th>Name</th>
          <th>Columns</th>
          <th>Referencing Table</th>
          <th>Referencing Columns</th>
          <th>Is Indexed</th>
          <th>On Update</th>
          <th>On Delete</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .ChildKeys}}
        <tr>
          <td class="tcnw">{{.RefConstraintName}}</td>
          <td>{{.RefTableColumns}}</td>
          <td class="tcnw">{{.SchemaName}}.<a href="../../{{.SchemaName}}/tables/{{.TableName}}.html">{{.TableName}}</a></td>
          <td>{{.TableColumns}}</td>
          <td class="tcc">{{.IsIndexed|checkMark}}</td>
          <td>{{.UpdateRule}}</td>
          <td>{{.DeleteRule}}</td>
          <td>{{.Comment|safeHTML}}</td>
        </tr>{{end}}
        </tbody>
      </table>
      <br />`

}

//      <h2>Dependencies</h2>
func tpltTableDependencies() string {
	return `
      <p><b>Parents (this depends on)</b></p>
      <table width="100.0%" id="dataTable-tdo" class="dataTable">
        <thead>
        <tr>
          <th>Object Schema</th>
          <th>Object Name</th>
          <th>Object Type</th>
        </tr>
        </thead>
        <tbody>{{range .Dependencies}}
        <tr>
          <td class="tcnw">{{.DepObjectSchema}}</td>
          <td class="tcnw">{{ if or ( or (eq .DepObjectType "TABLE" ) (eq .DepObjectType "VIEW")) ( or (eq .DepObjectType "MATERIALIZED VIEW") (eq .DepObjectType "FOREIGN TABLE")) }}<a href="../../{{.DepObjectSchema}}/tables/{{.DepObjectName}}.html">{{.DepObjectName}}</a>{{else}}{{.DepObjectName}}{{end}}</td>
          <td class="tcnw">{{.DepObjectType}}</td>
        </tr>{{end}}
        </tbody>
      </table>
      <br />`

}

func tpltTableDependents() string {
	return `
      <p><b>Children (depends on this)</b></p>
      <table width="100.0%" id="dataTable-dot" class="dataTable">
        <thead>
        <tr>
          <th>Object Schema</th>
          <th>Object Name</th>
          <th>Object Type</th>
        </tr>
        </thead>
        <tbody>{{range .Dependents}}
        <tr>
          <td class="tcnw">{{.ObjectSchema}}</td>
          <td class="tcnw">{{ if or ( or (eq .ObjectType "TABLE" ) (eq .ObjectType "VIEW")) ( or (eq .ObjectType "MATERIALIZED VIEW") (eq .ObjectType "FOREIGN TABLE")) }}<a href="../../{{.ObjectSchema}}/tables/{{.ObjectName}}.html">{{.ObjectName}}</a>{{else}}{{.ObjectName}}{{end}}</td>
          <td class="tcnw">{{.ObjectType}}</td>
        </tr>{{end}}
        </tbody>
      </table>
      <br />`
}

func tpltTableFDW() string {
	return `
      <h2>Foreign Data Wrapper</h2>
      <table width="100.0%" id="dataTable-fdw" class="dataTable">
        <thead>
        <tr>
          <th>Wrapper Name</th>
          <th>Server Name</th>
          <th>Wrapper Options</th>
          <th>Comments</th>
        </tr>
        </thead>
        <tbody>{{range .ForeignWrappers}}
        <tr>
          <td class="tcnw">{{.WrapperName}}</td>
          <td>{{.ServerName}}</td>
          <td>{{.WrapperOptions}}</td>
          <td>{{.Comment|safeHTML}}</td>
        </tr>{{end}}
        </tbody>
      </table>
      <br />`
}

func tpltTableQuery() string {
	return `
      <pre>
{{.Query}}
      </pre>`
}

func tpltOddHeader() string {
	return reportHead(true, false, false) + `
    <div id="pageBody">`
}

func tpltOddTables() string {
	return `
      <table width="100.0%" id="dataTable-tab" class="dataTable">
        <thead>
        <tr>
          <th>Table</th>
          <th>No PK</th>
          <th>No indices</th>
          <th>Duplicate indices</th>
          <th>Only one column</th>
          <th>No data</th>
          <th>No relationships</th>
          <th>Denormalized?</th>
        </tr>
        </thead>
        <tbody>{{range .OddTables}}
        <tr>
          <td><a href="tables/{{.TableName}}.html">{{.TableName}}</a></td>
          <td class="tcc">{{.NoPK|checkMark}}</td>
          <td class="tcc">{{.NoIndices|checkMark}}</td>
          <td class="tcc">{{.DuplicateIndices|checkMark}}</td>
          <td class="tcc">{{.OneColumn|checkMark}}</td>
          <td class="tcc">{{.NoData|checkMark}}</td>
          <td class="tcc">{{.NoRelationships|checkMark}}</td>
          <td class="tcc">{{.Denormalized|checkMark}}</td>
        </tr>{{end}}
        </tbody>
      </table>
      <br />`
}

func tpltOddColumns() string {
	return `
      <table width="100.0%" id="dataTable-col" class="dataTable">
        <thead>
        <tr>
          <th>Table</th>
          <th>Column</th>
          <th>Nullable and part of a unique index or constraint</th>
          <th>Nullable with a default value</th>
          <th>Defaults to NULL or 'NULL'</th>
        </tr>
        </thead>
        <tbody>{{range .OddColumns}}
        <tr>
          <td><a href="tables/{{.TableName}}.html">{{.TableName}}</a></td>
          <td>{{.ColumnName}}</td>
          <td class="tcc">{{.NullUnique|checkMark}}</td>
          <td class="tcc">{{.NullWithDefault|checkMark}}</td>
          <td class="tcc">{{.NullAsDefault|checkMark}}</td>
        </tr>{{end}}
        </tbody>
      </table>
      <br />`
}
