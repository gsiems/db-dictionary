package view

import (
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
	t "github.com/gsiems/db-dictionary/template"
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

		var tmplt t.T
		tmplt.AddPageHeader(1, md)

		context.Domains = md.FindDomains(vs.Name)
		if len(context.Domains) > 0 {
			tmplt.AddSnippet("SchemaDomains")
			sortDomains(context.Domains)
		} else {
			tmplt.AddSnippet("      <p><b>No domains extracted for this schema.</b></p>")
		}

		tmplt.AddPageFooter()

		dirName := md.OutputDir + "/" + vs.Name
		err = tmplt.RenderPage(dirName, "domains", context)
		if err != nil {
			return err
		}
	}

	return err
}