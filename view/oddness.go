package view

import (
	"path"
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
	t "github.com/gsiems/db-dictionary/template"
)

// oddTable contains the data used for generating the table portion of the schema odd-things page
type oddTable struct {
	TableName        string
	NoPK             string
	NoIndices        string
	DuplicateIndices string
	OneColumn        string
	NoData           string
	NoRelationships  string
	Denormalized     string
}

// oddColumn contains the data used for generating the column portion of the schema odd-things page
type oddColumn struct {
	TableName       string
	ColumnName      string
	NullUnique      string
	NullWithDefault string //
	NullAsDefault   string //
}

// oddnessView contains the data used for generating the schema odd-things page
type oddnessView struct {
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	OutputDir     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	OddTables     []oddTable
	OddColumns    []oddColumn
}

// oddness accumulates odd things and the template needed for generating the schema odd-things page
type oddness struct {
	minify  bool
	context oddnessView
	oT      t.T
}

// sortOddTables sets the default sort order for a list of odd table things
func sortOddTables(x []oddTable) {
	sort.Slice(x, func(i, j int) bool {
		return x[i].TableName < x[j].TableName
	})
}

// sortOddColumns sets the default sort order for a list of odd columns things
func sortOddColumns(x []oddColumn) {
	sort.Slice(x, func(i, j int) bool {

		switch strings.Compare(x[i].TableName, x[j].TableName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].ColumnName < x[j].ColumnName
	})
}

// initOddThings initializes the oddness accumulator for a schema
func initOddThings(md *m.MetaData, vs m.Schema) *oddness {

	var o oddness

	o.minify = md.Cfg.Minify

	o.context = oddnessView{
		Title:         "Odd things - " + md.Alias + "." + vs.Name,
		TmspGenerated: md.TmspGenerated,
		DBName:        md.Name,
		DBComment:     md.Comment,
		SchemaName:    vs.Name,
		SchemaComment: vs.Comment,
		OutputDir:     md.OutputDir,
	}

	o.oT.AddPageHeader(1, md)
	o.oT.AddSnippet("OddHeader")

	return &o
}

// checkOddThings scans the data for a table view and identifies odd characteristics
func (o *oddness) checkOddThings(tView *tableView) {

	//
	switch strings.ToUpper(tView.TableType) {
	case "TABLE", "BASE TABLE":
		//
	default:
		return
	}

	// potential odd table things
	var otm []string

	if tView.RowCount == 0 {
		otm = append(otm, "NoData")
	}

	if len(tView.PrimaryKeys) == 0 {
		otm = append(otm, "NoPK")
	}

	if len(tView.Columns) == 1 {
		otm = append(otm, "OneColumn")
	}

	if len(tView.Indexes) > 0 {

		dupChk := make(map[string]int)
		for _, idx := range tView.Indexes {
			dupChk[idx.IndexColumns]++
		}
		for _, kount := range dupChk {
			if kount > 1 {
				otm = append(otm, "DuplicateIndices")
			}
		}
	} else {
		otm = append(otm, "NoIndices")
	}
	if len(tView.ParentKeys) == 0 && len(tView.ChildKeys) == 0 {
		otm = append(otm, "NoRelationships")
	}

	// potential odd column things
	for _, vc := range tView.Columns {
		var ocm []string

		// Nullable yet has a default value
		if vc.IsNullable == "Y" && vc.Default != "" {
			ocm = append(ocm, "NullWithDefault")
		}

		// Has null or the string literal 'null' as the default
		switch {
		case strings.ToLower(vc.Default) == "null":
			ocm = append(ocm, "NullAsDefault")
		case strings.HasPrefix(strings.ToLower(vc.Default), "'null'"):
			ocm = append(ocm, "NullAsDefault")
		}

		// NullUnique Is nullable and part of unique constraint/index
		//if vc.IsNullable == "Y" {
		// TODO
		//}

		// Maybe is denormalized?
		lastByByte := vc.Name[len(vc.Name)-1:]
		switch lastByByte {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			otm = append(otm, "Denormalized")
		}

		//
		if len(ocm) > 0 {
			ts := oddColumn{
				TableName:  vc.TableName,
				ColumnName: vc.Name,
			}

			for _, v := range ocm {
				switch v {
				case "NullUnique":
					ts.NullUnique = "X"
				case "NullWithDefault":
					ts.NullWithDefault = "X"
				case "NullAsDefault":
					ts.NullAsDefault = "X"
				}
			}
			o.context.OddColumns = append(o.context.OddColumns, ts)
		}
	}

	//
	if len(otm) > 0 {
		ts := oddTable{
			TableName: tView.TableName,
		}

		for _, v := range otm {
			switch v {
			case "NoPK":
				ts.NoPK = "X"
			case "NoIndices":
				ts.NoIndices = "X"
			case "DuplicateIndices":
				ts.DuplicateIndices = "X"
			case "OneColumn":
				ts.OneColumn = "X"
			case "NoData":
				ts.NoData = "X"
			case "NoRelationships":
				ts.NoRelationships = "X"
			case "Denormalized":
				ts.Denormalized = "X"
			}
		}
		o.context.OddTables = append(o.context.OddTables, ts)
	}

}

// makeOddnessPage creates a schema odd things page using the accumulated oddness data
func (o *oddness) makeOddnessPage(md *m.MetaData) (err error) {

	o.oT.AddSectionHeader("Tables that display potential oddness")
	if len(o.context.OddTables) > 0 {
		sortOddTables(o.context.OddTables)
		o.oT.AddSnippet("OddTables")
	} else {
		o.oT.AddSnippet("      <p><b>No table oddities were extracted for this schema.</b></p>")
	}

	o.oT.AddSectionHeader("Columns that display potential oddness")
	if len(o.context.OddColumns) > 0 {
		sortOddColumns(o.context.OddColumns)
		o.oT.AddSnippet("OddColumns")
	} else {
		o.oT.AddSnippet("      <p><b>No column oddities were extracted for this schema.</b></p>")
	}

	o.oT.AddPageFooter(1, md)

	dirName := path.Join(o.context.OutputDir, o.context.SchemaName)
	err = o.oT.RenderPage(dirName, "odd-things", o.context, o.minify)
	if err != nil {
		return err
	}
	return err
}
