package view

import (
	"html/template"
	"os"
	"sort"

	m "github.com/gsiems/db-dictionary/model"
)

type tableView struct {
	Title         string
	PathPrefix    string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TableName     string
	TableComment  string
	TableType     string
	TmspGenerated string
	Query         string
	Columns       []m.Column
}

func SortColumns(columns []m.Column) {
	sort.Slice(columns, func(i, j int) bool {
		return columns[i].OrdinalPosition < columns[j].OrdinalPosition
	})
}

func RenderTables(d *m.Dictionary, s *[]m.Schema, t *[]m.Table) (err error) {

	const body1 = `    <div id="NavBar">
      <ul id="navlist">
        <li><a href="{{.PathPrefix}}index.html">Schemas</a></li>
        <li><a href="../tables.html">Tables</a></li>
        <li><a href="../columns.html">Columns</a></li>
      </ul>
    </div>
    <div id="PageHead"><h1>{{.Title}}</h1>
      <table>
        <tr><th>Generated:</th><td>{{.TmspGenerated}}</td><td></td></tr>
        <tr><th>Database:</th><td>{{.DBName}}</td><td class="TCcomment">{{.DBComment}}</td></tr>
        <tr><th>Schema:</th><td>{{.SchemaName}}</td><td class="TCcomment">{{.SchemaComment}}</td></tr>
        <tr><th>Table:</th><td>{{.TableName}}</td><td class="TCcomment">{{.TableComment}}</td></tr>
        <tr><th>Table Type:</th><td>{{.TableType}}</td><td></td></tr>
      </table>
    </div>
    <div id="PageBody">
      <h2>Columns</h2>
      <table width="100.0%" id="tablesorter-data" class="tablesorter">
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
            <td class="TC1">{{.Name}}</td>
            <td class="TC1">{{.OrdinalPosition}}</td>
            <td class="TC1">{{.DataType}}</td>
            <td class="TC1">{{.IsNullable}}</td>
            <td class="TC1">{{.Default}}</td>
            <td class="TCcomment">{{.Comment}}</td>
          </tr>{{end}}
        <tbody>
      </table>
      `

	head := header()
	foot := footer()

	for _, vt := range *t {
		for _, vs := range *s {
			if vt.SchemaName != vs.Name {
				continue
			}
			context := tableView{
				Title:         d.DBName + "." + vs.Name + "." + vt.Name,
				PathPrefix:    "../../",
				TmspGenerated: d.TmspGenerated,
				DBName:        d.DBName,
				DBComment:     d.DBComment,
				SchemaName:    vs.Name,
				SchemaComment: vs.Comment,
				TableName:     vt.Name,
				TableComment:  vt.Comment,
				TableType:     vt.TableType,
				Query:         vt.ViewDefinition,
				Columns:       vt.Columns,
			}

			if d.DBAlias != "" {
				context.Title = d.DBAlias + "." + vs.Name + "." + vt.Name
			}

			SortColumns(context.Columns)

			var query string
			if len(context.Query) > 0 {
				query = `
      <h2>Query</h2>
      <pre>
{{.Query}}
      </pre>`
			}

			templates, err := template.New("doc").Parse(head + body1 + query + foot)
			if err != nil {
				return err
			}

			dirName := d.OutputDir + "/" + vs.Name + "/tables/"
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
