package model

import (
	"time"

	"github.com/gsiems/db-dictionary/config"
	m "github.com/gsiems/go-db-meta/model"
)

type Dictionary struct {
	DBFile            string
	HasForeignServers bool
	TmspGenerated     string
	OutputDir         string
	DBMSVersion       string
	CharacterSet      string
	DBName            string
	DBAlias           string
	DBOwner           string
	DBComment         string
	CommentsFormat    string
}

func (md *MetaData) DBDictionary(dbe string, cfg config.Config, s m.Catalog) (r Dictionary, err error) {

	r = Dictionary{
		DBMSVersion:  s.DBMSVersion.String,
		CharacterSet: s.DefaultCharacterSetName.String,
		DBFile:       s.CatalogName.String,
		DBName:       s.CatalogName.String,
		DBOwner:      s.CatalogOwner.String,
		DBComment:    md.renderComment(s.Comment.String),
	}

	switch dbe {
	case "pg":
		r.HasForeignServers = true
	case "sqlite":
		r.DBAlias = cfg.DbName
	}

	switch cfg.CommentsFormat {
	case "markdown":
		r.CommentsFormat = cfg.CommentsFormat
	default:
		r.CommentsFormat = "none"
	}

	if cfg.OutputDir != "" {
		r.OutputDir = cfg.OutputDir
	} else {
		r.OutputDir = "."
	}

	// TODO: IncludeJS flag

	t := time.Now()
	r.TmspGenerated = t.Format("2006-01-02 15:04:05")

	return r, err
}
