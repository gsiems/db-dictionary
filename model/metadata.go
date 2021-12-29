package model

import (
	"path"
	"strings"
	"time"

	"github.com/gsiems/db-dictionary/config"
	m "github.com/gsiems/go-db-meta/model"
)

type Database struct {
	File         string
	Version      string
	CharacterSet string
	Name         string
	Alias        string
	Owner        string
	Comment      string
}

type MetaData struct {
	TmspGenerated  string
	DBEngine       string
	CommentsFormat string
	OutputDir      string
	File           string
	Version        string
	CharacterSet   string
	Name           string
	Alias          string
	Owner          string
	Comment        string
	Cfg            config.Config
	//ConnectInfo      *m.ConnectInfo
	Database          Database
	Schemas           []Schema
	Tables            []Table
	Columns           []Column
	Domains           []Domain
	Indexes           []Index
	CheckConstraints  []CheckConstraint
	UniqueConstraints []UniqueConstraint
	ForeignKeys       []ForeignKey
	PrimaryKeys       []PrimaryKey
	Dependencies      []Dependency
	Dependents        []Dependency
	UserTypes         []UserType
}

func Init(cfg config.Config) *MetaData {

	var md MetaData

	md.Cfg = cfg

	t := time.Now()
	md.TmspGenerated = t.Format("2006-01-02 15:04:05")

	if cfg.DbName != "" {
		md.Alias = cfg.DbName
	} else {
		tn := path.Base(cfg.File)
		te := path.Ext(tn)
		if te != "" {
			tn = strings.TrimSuffix(tn, te)
		}
		md.Alias = tn
	}

	switch cfg.CommentsFormat {
	case "markdown":
		md.CommentsFormat = cfg.CommentsFormat
	default:
		md.CommentsFormat = "none"
	}

	if cfg.OutputDir != "" {
		md.OutputDir = cfg.OutputDir
	} else {
		md.OutputDir = "."
	}

	return &md
}

//~ func (md *MetaData) LoadConnectInfo(x *m.ConnectInfo) {
//~ md.connectInfo = x
//~ }

func (md *MetaData) LoadCatalog(x *m.Catalog) {
	md.File = x.CatalogName.String
	md.Version = x.DBMSVersion.String
	md.CharacterSet = x.DefaultCharacterSetName.String
	md.Name = x.CatalogName.String
	md.Owner = x.CatalogOwner.String
	md.Comment = x.Comment.String
}

/*
   <h2>Tables that display potential oddness</h2>

       <th>Table</th>
       <th>No PK</th>
       <th>No indices</th>
       <th>Duplicate indices</th>
       <th>Only one column</th>
       <th>No data</th>
       <th>No relationships</th>
       <th>Denormalized?</th>

   <h2>Columns that display potential oddness</h2>

       <th>Table</th>
       <th>Column</th>
       <th>Nullable and part of a unique constraint</th>
       <th>Nullable and part of a unique index</th>
       <th>Nullable with a default value</th>
       <th>Defaults to NULL or 'NULL'</th>


*/
