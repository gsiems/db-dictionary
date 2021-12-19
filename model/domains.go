package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type Domain struct {
	DBName     string
	SchemaName string
	Name       string
	Owner      string
	DataType   string
	Default    string
	Comment    string
}

func (md *MetaData) LoadDomains(x *[]m.Domain) {
	for _, v := range *x {
		domain := Domain{
			DBName:     v.DomainCatalog.String,
			SchemaName: md.chkSchemaName(v.DomainSchema.String),
			Name:       v.DomainName.String,
			Owner:      v.DomainOwner.String,
			DataType:   v.DataType.String,
			Default:    v.DomainDefault.String,
			Comment:    v.Comment.String,
		}
		md.Domains = append(md.Domains, domain)
	}
	fmt.Printf("%d domains loaded\n", len(md.Domains))
}

func (md *MetaData) FindDomains(schemaName string) (d []Domain) {

	for _, v := range md.Domains {
		if schemaName == "" || schemaName == v.SchemaName {
			d = append(d, v)
		}
	}

	return d
}
