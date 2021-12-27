package view

import (
	"html/template"
	"os"
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
)

type domainsView struct {
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	Domains       []m.Domain
}

func sortDomains(x []m.Domain) {
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

func makeDomainsList(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		context := domainsView{
			Title:         "Domains for " + md.Alias + "." + vs.Name,
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		var pageParts []string
		pageParts = append(pageParts, pageHeader(1, md))

		context.Domains = md.FindDomains(vs.Name)
		if len(context.Domains) > 0 {
			pageParts = append(pageParts, tpltSchemaDomains())
			sortDomains(context.Domains)
		} else {
			pageParts = append(pageParts, "      <p><b>No domains extracted for this schema.</b></p>")
		}

		pageParts = append(pageParts, pageFooter())

		templates, err := template.New("doc").Funcs(template.FuncMap{
			"safeHTML": func(u string) template.HTML { return template.HTML(u) },
		}).Parse(strings.Join(pageParts, ""))
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

		outfile, err := os.Create(dirName + "/domains.html")
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
