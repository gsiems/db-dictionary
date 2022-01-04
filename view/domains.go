package view

import (
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
	t "github.com/gsiems/db-dictionary-core/template"
)

// domainsView contains the data used for generating the schema domains page
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

// sortDomains sets the default sort order for a list of domains
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

// makeDomainsList marshals the data needed for, and then creates, a schema domains page
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
			return nil
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
