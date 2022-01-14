package view

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
)

// makeCSS generates the needed CSS file(s)
func makeCSS(md *m.MetaData) (err error) {

	dirName := md.OutputDir + "/css"
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0744)
		if err != nil {
			return err
		}
	}

	if md.Cfg.CSSFiles != "" {
		err = copyCSSFiles(dirName, md.Cfg.CSSFiles)
	} else {
		err = writeDefaultCSS(dirName)
	}

	return err
}

func copyCSSFiles(dirName, files string) (err error) {

	f := strings.Split(files, ",")
	for _, source := range f {

		input, err := ioutil.ReadFile(source)
		if err != nil {
			return err
		}

		target := dirName + "/" + path.Base(source)
		err = ioutil.WriteFile(target, input, 0644)
		if err != nil {
			return err
		}
	}

	return err
}

func writeDefaultCSS(dirName string) (err error) {

	outfile, err := os.Create(dirName + "/blues.css")
	if err != nil {
		return err
	}
	defer outfile.Close()

	_, err = outfile.WriteString(defaultCSS())
	return err

}

// defaultCSS is the CSS for the default css file. Currently embedded for the purpose
// of creating a single-file deployment with minimal fuss
func defaultCSS() string {
	return `
:root {
    /* Define the color pallet */
    --grey-1: hsl(0, 0%, 95%);
    --grey-2: hsl(0, 0%, 85%);
    --grey-3: hsl(0, 0%, 70%);
    --grey-4: hsl(0, 0%, 50%);
    --blue-1: hsl(240, 67%, 94%); /* Lavender */
    --blue-2: hsl(207, 48%, 58%);
    --blue-3: hsl(207, 39%, 48%);
    --blue-4: hsl(207, 48%, 38%); /* pg blue */
    --blue-5: hsl(207, 39%, 18%);
    /* */
    --page-base: var(--grey-1);
    --odd-zebra: var(--blue-1);
    --even-zebra: var(--grey-2);
    --dark-border: 2px solid var(--grey-4);
    --light-border: 2px solid var(--grey-2);
    --border-width: 2px;
    --line-width: 1px;
    --light-text: var(--grey-1);
    --medium-text: var(--blue-3);
    --dark-text: var(--blue-5);
}
body {
    background-color: var(--page-base);
    font-family: verdana,helvetica,sans-serif;
    margin: 0;
    padding: 0;
}
#topNav {
    overflow: hidden;
    background-color: var(--blue-2);
}
#topNav a {
    background-color: var(--blue-4);
    color: var(--light-text);
    border-left: var(--light-border);
    border-top: var(--light-border);
    border-bottom: var(--dark-border);
    border-right: var(--dark-border);
    float: left;
    padding: 2px 20px;
    text-align: center;
    text-decoration: none;
}
#topNav a:hover {
    background-color: var(--blue-2);
    color: var(--light-text);
}
#topNav a.active {
    background-color: var(--blue-2);
    color: var(--light-text);
}
#pageHead {
    background-color: var(--blue-2);
    border-bottom: var(--dark-border);
    color: var(--light-text);
    margin-bottom: 0px;
    padding-bottom: 0px;
    padding-left: 10px;
    padding-right: 10px;
    padding-top: 5px;
}
#pageHead h1 {
    border-bottom-style: solid;
    border-bottom-width: var(--line-width);
    font-size: 150%;
    font-weight: bold;
    margin-top: 10px;
    padding-top: 0.5ex;
    padding-bottom: 0.5ex;
}
#pageHead .headingContainer {
  display: grid;
  grid-template-columns: auto auto;
  justify-content: start;
}
#pageHead .headingItem {
    padding-right: 5px;
    text-align: left;
    vertical-align: top;
}
#pageBody {
    font-size: 90%;
    margin-bottom: 10px;
    margin-left: 10px;
    margin-right: 10px;
}
#pageBody h2 {
    color: var(--medium-text);
    font-size: 130%;
}
#pageBody h3 {
    color: var(--medium-text);
    font-style: italic;
}
#pageBody hr {
    margin-top: 10px;
    color: var(--blue-2);
}
#pageFoot {
    background-color: var(--blue-2);
    color: var(--light-text);
    border-left: var(--light-border);
    border-top: var(--light-border);
    border-bottom: var(--dark-border);
    border-right: var(--dark-border);
    font-size: 90%;
    font-weight: bold;
    margin-top: 10px;
    margin-left: 10px;
    margin-right: 10px;
    padding: 2px 7px 2px 2px;
}
#pageFoot a {
    color: var(--dark-text);
}
#filter-form {
    vertical-align: top;
    padding-bottom: 1.0ex;
    padding-top: 0.5ex;
}
#filterBy {
    background-color: var(--page-base);
    color: var(--dark-text);
    border-bottom: var(--light-border);
    border-right: var(--light-border);
    border-left: var(--dark-border);
    border-top: var(--dark-border);
}
table.dataTable {
    border-bottom: var(--dark-border);
    border-right: var(--dark-border);
    border-left: var(--light-border);
    border-top: var(--light-border);
    text-align: left;
    vertical-align: top;
    width: 100%;
}
table.dataTable thead tr th {
    background-color: var(--grey-2);
    border-bottom: var(--dark-border);
    border-right: var(--dark-border);
    color: var(--blue-4);
    /*cursor: pointer;*/
    padding-left: 4px;
    padding-right: 15px;
}
table.dataTable thead tr .headerSortUp {
    background-color: var(--grey-3);
}
table.dataTable thead tr .headerSortUp::after {
    content: "\25B2";
}
table.dataTable thead tr .headerSortDown {
    background-color: var(--grey-3);
}
table.dataTable thead tr .headerSortDown::after {
    content: "\25BC";
}

table.dataTable tbody tr:nth-child(odd) {
    background-color: var(--odd-zebra);
    color: var(--dark-text);
}
table.dataTable tbody tr:nth-child(even) {
    background-color: var(--even-zebra);
    color: var(--dark-text);
}
table.dataTable tbody td {
    padding-left: 4px;
    padding-right: 4px;
    vertical-align: top;
}
table.dataTable a {
    color: var(--blue-4);
}
table.dataTable tbody td.tcnw {
    white-space: nowrap;
}
table.dataTable tbody td.tcc {
    text-align: center;
}
table.dataTable tbody td.tcn {
    /* Numeric table cells. */
    text-align: right;
    white-space: nowrap;
}
pre {
    background-color: var(--odd-zebra);
    color: var(--dark-text);
    border-left: var(--light-border);
    border-top: var(--light-border);
    border-bottom: var(--dark-border);
    border-right: var(--dark-border);
    padding: 5px;
    font-size: 100%;
    font-family: Bitstream Vera Sans Mono, Consolas, Courier New, DejaVu Sans Mono, Liberation Mono, Lucida Console, Monaco, monospace;
}
`
}
