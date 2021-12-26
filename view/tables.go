package view

import (
	"html/template"
	"os"
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
)

type tablesView struct {
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	Tables        []m.Table
}

func sortTables(x []m.Table) {
	sort.Slice(x, func(i, j int) bool {
		switch strings.Compare(x[i].SchemaName, x[j].SchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].Name < x[j].Name
	})
}

func makeTableList(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		context := tablesView{
			Title:         "Tables for " + md.Alias + "." + vs.Name,
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		var pageParts []string
		pageParts = append(pageParts, pageHeader(1))

		context.Tables = md.FindTables(vs.Name)
		if len(context.Tables) > 0 {
			pageParts = append(pageParts, tpltSchemaTables())
			sortTables(context.Tables)
		} else {
			pageParts = append(pageParts, "      <p><b>No tables extracted for this schema.</b></p>")
		}

		pageParts = append(pageParts, pageFooter())

		templates, err := template.New("doc").Parse(strings.Join(pageParts, ""))
		if err != nil {
			return err
		}

		dirName := md.OutputDir + "/" + vs.Name
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
