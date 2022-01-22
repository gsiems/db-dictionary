package template

import (
	"html/template"
	"os"
	"path"
	"regexp"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
)

// T contains the tempate snippets that get concatenated to create the page template
type T struct {
	sectionCount int
	snippets     []string
}

// C is the empty interface used for allowing the different page type
// structures (contexts) to use the same page generation function
type C interface {
}

// AddSnippet add a template snippet to the template source accumulator
func (t *T) AddSnippet(s string) {
	switch s {
	case "Schemas":
		t.snippets = append(t.snippets, tpltSchemas())
	case "SchemaTables":
		t.snippets = append(t.snippets, tpltSchemaTables())
	case "SchemaDomains":
		t.snippets = append(t.snippets, tpltSchemaDomains())
	case "SchemaColumns":
		t.snippets = append(t.snippets, tpltSchemaColumns())
	case "SchemaConstraintsHeader":
		t.snippets = append(t.snippets, tpltSchemaConstraintsHeader())
	case "SchemaCheckConstraints":
		t.snippets = append(t.snippets, tpltSchemaCheckConstraints())
	case "SchemaUniqueConstraints":
		t.snippets = append(t.snippets, tpltSchemaUniqueConstraints())
	case "SchemaFKConstraints":
		t.snippets = append(t.snippets, tpltSchemaFKConstraints())
	case "TableConstraintsHeader":
		t.snippets = append(t.snippets, tpltTableConstraintsHeader())
	case "TableConstraintsFooter":
		t.snippets = append(t.snippets, tpltTableConstraintsFooter())
	case "TableCheckConstraints":
		t.snippets = append(t.snippets, tpltTableCheckConstraints())
	case "TablePrimaryKey":
		t.snippets = append(t.snippets, tpltTablePrimaryKey())
	case "TableUniqueConstraints":
		t.snippets = append(t.snippets, tpltTableUniqueConstraints())
	case "TableIndexes":
		t.snippets = append(t.snippets, tpltTableIndexes())
	case "TableParentKeys":
		t.snippets = append(t.snippets, tpltTableParentKeys())
	case "TableChildKeys":
		t.snippets = append(t.snippets, tpltTableChildKeys())
	case "TableDependencies":
		t.snippets = append(t.snippets, tpltTableDependencies())
	case "TableDependents":
		t.snippets = append(t.snippets, tpltTableDependents())
	case "TableFDW":
		t.snippets = append(t.snippets, tpltTableFDW())
	case "TableQuery":
		t.snippets = append(t.snippets, tpltTableQuery())
	case "OddHeader":
		t.snippets = append(t.snippets, tpltOddHeader())
	case "OddTables":
		t.snippets = append(t.snippets, tpltOddTables())
	case "OddColumns":
		t.snippets = append(t.snippets, tpltOddColumns())
	default:
		// Assertion: if the string does not match the known snippets then it
		// must be the actual string to append
		t.snippets = append(t.snippets, s)
	}
}

// AddTableHead adds the snippet for a table report header (based on the table type)
func (t *T) AddTableHead(tabType string) {
	t.snippets = append(t.snippets, tpltTableHead(tabType))
}

// AddTableColumns adds the snippet for the columns list for a table page
// (displayed columns depend on table type: table, view, etc.)
func (t *T) AddTableColumns(tabType string) {
	t.snippets = append(t.snippets, tpltTableColumns(tabType))
}

// AddPageHeader adds the initial snippet for the page to create (HTML head
// plus navigation menu bar)
func (t *T) AddPageHeader(i int, md *m.MetaData) {
	t.AddSnippet(pageHeader(i, md))
}

// AddPageFooter adds the final snippet for generating the end of the page to create
func (t *T) AddPageFooter(i int, md *m.MetaData) {
	t.AddSnippet(pageFooter(i, md))
}

// AddSectionHeader adds a section header for pages with multiple sections
func (t *T) AddSectionHeader(s string) {
	if t.sectionCount > 0 {
		t.AddSnippet("<hr/>")
	}
	t.sectionCount++
	t.AddSnippet(sectionHeader(s))
}

// RenderPage takes the supplied view context and renders/writes the html file
func (t *T) RenderPage(dirName, fileName string, context C, minify bool) error {

	ft := strings.Join(t.snippets, "")

	if minify {
		/*
		   Simple minify
		   running `du -s` on output
		   2192: pre-minification
		   1932: remove leading spaces -- 11.86% reduction
		   1920: remove line breaks -- ~ 12.41% reduction
		*/

		re := regexp.MustCompile(`\n *`)
		ft = re.ReplaceAllString(ft, "\n")
		re2 := regexp.MustCompile(`>\n<`)
		ft = re2.ReplaceAllString(ft, "><")
		re3 := regexp.MustCompile(`\}\n<`)
		ft = re3.ReplaceAllString(ft, "}<")
	}

	// parse the template
	templates, err := template.New("doc").Funcs(template.FuncMap{
		"safeHTML": func(u string) template.HTML { return template.HTML(u) },
		"checkMark": func(u string) template.HTML {
			switch strings.ToUpper(u) {
			case "X", "YES", "Y":
				return "✓"
			}
			return ""
		},
	}).Parse(ft)
	if err != nil {
		return err
	}

	// ensure that the file directory exists
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0744)
		if err != nil {
			return err
		}
	}

	// create the file
	outfile, err := os.Create(path.Join(dirName, fileName+".html"))
	if err != nil {
		return err
	}
	defer outfile.Close()

	// render and write the file
	err = templates.Lookup("doc").Execute(outfile, context)
	if err != nil {
		return err
	}

	return err
}
