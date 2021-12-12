package view

import (
	"html/template"
	"os"
	"sort"

	m "github.com/gsiems/db-dictionary/model"
)

type schemasView struct {
	Title         string
	PathPrefix    string
	TmspGenerated string
	DBMSVersion   string
	DBName        string
	DBComment     string
	Schemas       []m.Schema
}

func makeCSS(d *m.Dictionary) (err error) {

	dirName := d.OutputDir + "/css"
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0744)
		if err != nil {
			return err
		}
	}

	outfile, err := os.Create(dirName + "/main.css")
	if err != nil {
		return err
	}
	defer outfile.Close()

	_, err = outfile.WriteString(css())
	return err
}

func SortSchemas(schemas []m.Schema) {
	sort.Slice(schemas, func(i, j int) bool {
		return schemas[i].Name < schemas[j].Name
	})
}

func RenderSchemaList(d *m.Dictionary, s *[]m.Schema) (err error) {

	if d.OutputDir != "." {
		dirName := d.OutputDir
		_, err = os.Stat(dirName)
		if os.IsNotExist(err) {
			err = os.Mkdir(dirName, 0744)
			if err != nil {
				return err
			}
		}

	}

	err = makeCSS(d)
	if err != nil {
		return err
	}

	var navbar string
	if d.HasForeignServers {
		navbar = `    <div id="NavBar">
	      <ul id="navlist">
	        <li><a href="foreign_servers.html">Foreign servers</a></li>
	        <li><a href="index.html">Schemas</a></li>
	      </ul>
	    </div>`
	} else {
		navbar = `    <div id="NavBar">
	      <ul id="navlist">
	        <li><a href="index.html">Schemas</a></li>
	      </ul>
	    </div>`
	}

	const body = `  <body>
    <div id="PageHead"><h1>{{.Title}}</h1>
      <table width="100.0%">
        <tr><th>Generated:</th><td>{{.TmspGenerated}}</td><td></td></tr>
        <tr><th>Database Version:</th><td colspan="2">{{.DBMSVersion}}</td></tr>
        <tr><th>Database:</th><td>{{.DBName}}</td><td class="TCcomment">{{.DBComment}}</td></tr>
      </table>
    </div>
    <div id="PageBody">
      <table width="100.0%" id="tablesorter-data" class="tablesorter">
        <thead>
        <tr>
          <th>Schema</th>
          <th>Owner</th>
          <th>Comment</th>
        </tr>
        </thead>
        <tbody>{{range .Schemas}}
          <tr>
            <td class="TC1"><a href="{{.Name}}/tables.html">{{.Name}}</a></td>
            <td class="TC1">{{.Owner}}</td>
            <td class="TCcomment">{{.Comment}}</td>
          </tr>{{end}}
        <tbody>
      </table>`

	head := header()
	foot := footer()

	context := schemasView{
		Title:         "Schemas for " + d.DBName,
		TmspGenerated: d.TmspGenerated,
		DBMSVersion:   d.DBMSVersion,
		DBName:        d.DBName,
		DBComment:     d.DBComment,
		Schemas:       *s,
	}

	SortSchemas(context.Schemas)

	templates, err := template.New("doc").Parse(head + navbar + body + foot)
	if err != nil {
		return err
	}

	//
	outfile, err := os.Create(d.OutputDir + "/index.html")
	if err != nil {
		return err
	}
	defer outfile.Close()

	err = templates.Lookup("doc").Execute(outfile, context)
	if err != nil {
		return err
	}

	return err
}
