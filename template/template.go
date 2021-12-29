package template

import (
	"html/template"
	"os"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
)

type T struct {
	pageParts []string
}

type C interface {
}

func (t *T) AddSnippet(s string) {
	switch s {
	case "Schemas":
		t.pageParts = append(t.pageParts, tpltSchemas())
	case "SchemaTables":
		t.pageParts = append(t.pageParts, tpltSchemaTables())
	case "SchemaDomains":
		t.pageParts = append(t.pageParts, tpltSchemaDomains())
	case "SchemaColumns":
		t.pageParts = append(t.pageParts, tpltSchemaColumns())
	case "SchemaConstraintsHeader":
		t.pageParts = append(t.pageParts, tpltSchemaConstraintsHeader())
	case "SchemaCheckConstraints":
		t.pageParts = append(t.pageParts, tpltSchemaCheckConstraints())
	case "SchemaUniqueConstraints":
		t.pageParts = append(t.pageParts, tpltSchemaUniqueConstraints())
	case "SchemaFKConstraints":
		t.pageParts = append(t.pageParts, tpltSchemaFKConstraints())
	case "TableConstraintsHeader":
		t.pageParts = append(t.pageParts, tpltTableConstraintsHeader())
	case "TableConstraintsFooter":
		t.pageParts = append(t.pageParts, tpltTableConstraintsFooter())
	case "TableCheckConstraints":
		t.pageParts = append(t.pageParts, tpltTableCheckConstraints())
	case "TablePrimaryKey":
		t.pageParts = append(t.pageParts, tpltTablePrimaryKey())
	case "TableUniqueConstraints":
		t.pageParts = append(t.pageParts, tpltTableUniqueConstraints())
	case "TableIndexes":
		t.pageParts = append(t.pageParts, tpltTableIndexes())
	case "TableParentKeys":
		t.pageParts = append(t.pageParts, tpltTableParentKeys())
	case "TableChildKeys":
		t.pageParts = append(t.pageParts, tpltTableChildKeys())
	case "TableDependencies":
		t.pageParts = append(t.pageParts, tpltTableDependencies())
	case "TableDependents":
		t.pageParts = append(t.pageParts, tpltTableDependents())
	case "TableFDW":
		t.pageParts = append(t.pageParts, tpltTableFDW())
	case "TableQuery":
		t.pageParts = append(t.pageParts, tpltTableQuery())
	default:
		// Assertion: if the string does not match then it must be the actual string to append
		t.pageParts = append(t.pageParts, s)
	}
}

func (t *T) AddTableHead(tabType string) {
	t.pageParts = append(t.pageParts, tpltTableHead(tabType))
}

func (t *T) AddTableColumns(tabType string) {
	t.pageParts = append(t.pageParts, tpltTableColumns(tabType))
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
	}).Parse(strings.Join(t.pageParts, ""))
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
