package view

import (
	"html/template"
	"os"

	m "github.com/gsiems/db-dictionary/model"
)

func RenderColumns(d *m.Dictionary, s *[]m.Schema, t *[]m.Table) (err error) {

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
          <th>Column</th>
          <th>Ordinal Position</th>
          <th>Data Type</th>
          <th>Nulls</th>
          <th>Default</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Tables}}{{range .Columns}}
          <tr>
            <td class="TC1"><a href="tables/{{.TableName}}.html">{{.TableName}}</a></td>
            <td class="TC1">{{.Name}}</td>
            <td class="TC1">{{.OrdinalPosition}}</td>
            <td class="TC1">{{.DataType}}</td>
            <td class="TC1">{{.IsNullable}}</td>
            <td class="TC1">{{.Default}}</td>
            <td class="TCcomment">{{.Comment}}</td>
          </tr>{{end}}{{end}}
        <tbody>
      </table>`

	head := header()
	foot := footer()

	for _, vs := range *s {

		context := tablesView{
			Title:         "Columns for " + d.DBName + "." + vs.Name,
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

		outfile, err := os.Create(dirName + "/columns.html")
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
