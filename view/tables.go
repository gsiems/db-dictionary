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
	TmspGenerated    string
	Tables           []m.Table
}

func SortTables(tables []m.Table) {
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].Name < tables[j].Name
	})
}

func RenderTableList(d *m.Dictionary, s *[]m.Schema, t *[]m.Table) (err error) {

	var navbar string
	if d.HasForeignServers {
		navbar = `    <div id="NavBar">
      <ul id="navlist">
        <li><a href="{{.PathPrefix}}foreign_servers.html">Foreign servers</a></li>
        <li><a href="{{.PathPrefix}}index.html">Schemas</a></li>
        <li><a href="tables.html">Tables</a></li>
      </ul>
    </div>`
	} else {
		navbar = `    <div id="NavBar">
      <ul id="navlist">
        <li><a href="{{.PathPrefix}}index.html">Schemas</a></li>
        <li><a href="tables.html">Tables</a></li>
      </ul>
    </div>`
	}

	const body = `  <body>
    <div id="PageHead"><h1>{{.Title}}</h1>
      <table>
        <tr><th>Generated:</th><td>{{.TmspGenerated}}</td></tr>
        <tr><th>Schema:</th><td>{{.SchemaName}}</td></tr>{{if .HasSchemaComment}}
        <tr><th>Schema Comment:</th><td><div class="comments">{{.SchemaComment}}</div></td></tr>{{end}}
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
            <td class="TC1"><a href="{{.Name}}/tables.html">{{.Name}}</a></td>
            <td class="TC1">{{.Owner}}</td>
            <td class="TC1">{{.TableType}}</td>
            <td class="TC1">{{.RowCount}}</td>
            <td class="TC1"><div class="comments">{{.Comment}}</div></td>
          </tr>{{end}}
        <tbody>
      </table>`

	head := header()
	foot := footer()

	for _, vs := range *s {

		context := tableView{
			Title:         "Tables for " + d.DBName + "." + vs.Name,
			PathPrefix:    "../",
			TmspGenerated: d.TmspGenerated,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
			//BinFile: d.BinFile,
			//Tables: *s,
		}

		if d.DBComment != "" {
			context.HasDBComment = true
		}
		if vs.Comment != "" {
			context.HasSchemaComment = true
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

		templates.Lookup("doc").Execute(outfile, context)

	}

	return err
}
