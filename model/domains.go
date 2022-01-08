package model

import (
	"log"

	m "github.com/gsiems/go-db-meta/model"
)

// Domain contains the metadata for a user defined domain
type Domain struct {
	DBName     string
	SchemaName string
	Name       string
	Owner      string
	DataType   string
	Default    string
	Comment    string
}

// LoadDomains loads the user defined domain information from go-db-meta
// into the dictionary metadata structure
func (md *MetaData) LoadDomains(x *[]m.Domain) {
	for _, v := range *x {
		domain := Domain{
			DBName:     v.DomainCatalog.String,
			SchemaName: md.chkSchemaName(v.DomainSchema.String),
			Name:       v.DomainName.String,
			Owner:      v.DomainOwner.String,
			DataType:   v.DataType.String,
			Default:    v.DomainDefault.String,
			Comment:    md.renderComment(v.Comment.String),
		}
		md.Domains = append(md.Domains, domain)
	}
	if !md.Cfg.Quiet {
		log.Printf("%d domains loaded\n", len(md.Domains))
	}
}

// FindDomains returns the user-defined domain metadata that matches the
// specified schema. If no schema is specified then all user-defined domains
// are returned.
func (md *MetaData) FindDomains(schemaName string) (d []Domain) {

	for _, v := range md.Domains {
		if schemaName == "" || schemaName == v.SchemaName {
			d = append(d, v)
		}
	}

	return d
}
