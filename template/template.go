package template

import (
	"html/template"
	"os"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
)

type T struct {
	snippets []string
}

type C interface {
}

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
		// Assertion: if the string does not match then it must be the actual string to append
		t.snippets = append(t.snippets, s)
	}
}

func (t *T) AddTableHead(tabType string) {
	t.snippets = append(t.snippets, tpltTableHead(tabType))
}

func (t *T) AddTableColumns(tabType string) {
	t.snippets = append(t.snippets, tpltTableColumns(tabType))
}

func (t *T) AddPageHeader(i int, md *m.MetaData) {
	t.AddSnippet(pageHeader(i, md))
}

func (t *T) AddPageFooter() {
	t.AddSnippet(pageFooter())
}

func (t *T) AddSectionHeader(s string) {
	t.AddSnippet(sectionHeader(s))
}

func (t *T) RenderPage(dirName, fileName string, context C) error {

	templates, err := template.New("doc").Funcs(template.FuncMap{
		"safeHTML": func(u string) template.HTML { return template.HTML(u) },
	}).Parse(strings.Join(t.snippets, ""))
	if err != nil {
		return err
	}

	//dirName := md.OutputDir + "/" + vs.Name + "/tables/"
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0744)
		if err != nil {
			return err
		}
	}

	//outfile, err := os.Create(dirName + "/" + vt.Name + ".html")
	outfile, err := os.Create(dirName + "/" + fileName + ".html")
	if err != nil {
		return err
	}
	defer outfile.Close()

	err = templates.Lookup("doc").Execute(outfile, context)
	if err != nil {
		return err
	}

	return err
}
