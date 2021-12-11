package model

import (
	"time"

	"github.com/gsiems/db-dictionary/config"
	m "github.com/gsiems/go-db-meta/model"
)

type Dictionary struct {
	//BinFile                 string
	HasForeignServers bool
	TmspGenerated     string
	OutputDir         string
	DBMSVersion       string
	CharacterSet      string
	DBName            string
	DBOwner           string
	DBComment         string
	CommentsFormat    string
}

func DBDictionary(dbe string, cfg config.Config, s m.Catalog) (r Dictionary, err error) {

	r = Dictionary{
		DBMSVersion:  s.DBMSVersion.String,
		CharacterSet: s.DefaultCharacterSetName.String,
		DBName:       s.CatalogName.String,
		DBOwner:      s.CatalogOwner.String,
		DBComment:    s.Comment.String,
	}

	r.HasForeignServers = dbe == "pg"

	if cfg.OutputDir != "" {
		r.OutputDir = cfg.OutputDir
	} else {
		r.OutputDir = "."
	}

	switch cfg.CommentsFormat {
	case "markdown":
		r.CommentsFormat = cfg.CommentsFormat
	default:
		r.CommentsFormat = "none"
	}

	// TODO: IncludeJS flag

	t := time.Now()
	r.TmspGenerated = t.Format("2006-01-02 15:04:05")

	return r, err
}
