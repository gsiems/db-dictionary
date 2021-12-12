package view

import (
	"html/template"
	"os"
	"sort"

	m "github.com/gsiems/db-dictionary/model"
)

type tablesView struct {
	Title         string
	PathPrefix    string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	Tables        []m.Table
}

func SortTables(tables []m.Table) {
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].Name < tables[j].Name
	})
}

func RenderTableList(d *m.Dictionary, s *[]m.Schema, t *[]m.Table) (err error) {

	var navbar string
	if d.HasForeignServers {
		navbar = `  <body>
    <div id="NavBar">
      <ul id="navlist">
        <li><a href="{{.PathPrefix}}foreign_servers.html">Foreign servers</a></li>
        <li><a href="{{.PathPrefix}}index.html">Schemas</a></li>
        <li><a href="tables.html">Tables</a></li>
        <li><a href="columns.html">Columns</a></li>
      </ul>
    </div>`
	} else {
		navbar = `  <body>
    <div id="NavBar">
      <ul id="navlist">
        <li><a href="{{.PathPrefix}}index.html">Schemas</a></li>
        <li><a href="tables.html">Tables</a></li>
        <li><a href="columns.html">Columns</a></li>
      </ul>
    </div>`
	}

	const body = `    <div id="PageHead"><h1>{{.Title}}</h1>
      <table>
        <tr><th>Generated:</th><td>{{.TmspGenerated}}</td><td></td></tr>
        <tr><th>Database:</th><td>{{.DBName}}</td><td class="TCcomment">{{.DBComment}}</td></tr>
        <tr><th>Schema:</th><td>{{.SchemaName}}</td><td class="TCcomment">{{.SchemaComment}}</td></tr>
      </table>
    </div>
    <div id="PageBody">
      <table width="100.0%" id="tablesorter-data" class="tablesorter">
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
            <td class="TC1"><a href="tables/{{.Name}}.html">{{.Name}}</a></td>
            <td class="TC1">{{.Owner}}</td>
            <td class="TC1">{{.TableType}}</td>
            <td class="TC1">{{.RowCount}}</td>
            <td class="TCcomment">{{.Comment}}</td>
          </tr>{{end}}
        <tbody>
      </table>`

	head := header()
	foot := footer()

	for _, vs := range *s {

		context := tablesView{
			Title:         "Tables for " + d.DBName + "." + vs.Name,
			PathPrefix:    "../",
			TmspGenerated: d.TmspGenerated,
			DBName:        d.DBName,
			DBComment:     d.DBComment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		for _, vt := range *t {
			if vt.SchemaName == vs.Name {
				context.Tables = append(context.Tables, vt)
			}
		}
		SortTables(context.Tables)

		templates, err := template.New("doc").Parse(head + navbar + body + foot)
		if err != nil {
			return err
		}

		dirName := d.OutputDir + "/" + vs.Name
		_, err = os.Stat(dirName)
		if os.IsNotExist(err) {
			err = os.Mkdir(dirName, 0744)
			if err != nil {
				return err
			}
		}

		outfile, err := os.Create(dirName + "/tables.html")
		if err != nil {
			return err
		}
		defer outfile.Close()

		err = templates.Lookup("doc").Execute(outfile, context)
		if err != nil {
			return err
		}
	}

	return err
}
