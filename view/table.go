package view

import (
	"html/template"
	"os"
	"sort"

	m "github.com/gsiems/db-dictionary/model"
)

type tableView struct {
	Title            string
	PathPrefix       string
	DBMSVersion      string
	DBName           string
	DBComment        string
	HasDBComment     bool
	SchemaName       string
	SchemaComment    string
	HasSchemaComment bool
	TableName        string
	TableComment     string
	HasTableComment  bool
	TmspGenerated    string
	Query string
	HasQuery bool
	Columns          []m.Column
}

func SortColumns(columns []m.Column) {
	sort.Slice(columns, func(i, j int) bool {
		return columns[i].OrdinalPosition < columns[j].OrdinalPosition
	})
}

func RenderTables(d *m.Dictionary, s *[]m.Schema, t *[]m.Table) (err error) {

	var navbar string
	if d.HasForeignServers {
		navbar = `    <div id="NavBar">
      <ul id="navlist">
        <li><a href="{{.PathPrefix}}foreign_servers.html">Foreign servers</a></li>
        <li><a href="{{.PathPrefix}}index.html">Schemas</a></li>
        <li><a href="../tables.html">Tables</a></li>
      </ul>
    </div>`
	} else {
		navbar = `    <div id="NavBar">
      <ul id="navlist">
        <li><a href="{{.PathPrefix}}index.html">Schemas</a></li>
        <li><a href="../tables.html">Tables</a></li>
      </ul>
    </div>`
	}

	const body = `  <body>
    <div id="PageHead"><h1>{{.Title}}</h1>
      <table>
        <tr><th>Generated:</th><td>{{.TmspGenerated}}</td></tr>
        <tr><th>Schema:</th><td>{{.SchemaName}}</td></tr>{{if .HasSchemaComment}}
        <tr><th>Schema Comment:</th><td><div class="comments">{{.SchemaComment}}</div></td></tr>{{end}}
        <tr><th>Table:</th><td>{{.TableName}}</td></tr>{{if .HasTableComment}}
        <tr><th>Table Comment:</th><td><div class="comments">{{.TableComment}}</div></td></tr>{{end}}
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
            <td class="TC1"><div class="comments">{{.Comment}}</div></td>
          </tr>{{end}}
        <tbody>
      </table>

      {{if .HasQuery}}<h2>Query</h2>
      <pre>
{{ .Query }}
      </pre>{{end}}
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
				SchemaName:    vs.Name,
				SchemaComment: vs.Comment,
				TableName:     vt.Name,
				TableComment:  vt.Comment,
				Query: vt.ViewDefinition,
				//BinFile: d.BinFile,
				Columns: vt.Columns,
			}

			if context.DBComment != "" {
				context.HasDBComment = true
			}
			if context.SchemaComment != "" {
				context.HasSchemaComment = true
			}
			if context.TableComment != "" {
				context.HasTableComment = true
			}
			if context.Query != "" {
				context.HasQuery = true
			}

			SortColumns(context.Columns)

			templates, err := template.New("doc").Parse(head + navbar + body + foot)
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
