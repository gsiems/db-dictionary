package view

import (
	"os"

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

	outfile, err := os.Create(dirName + "/main.css")
	if err != nil {
		return err
	}
	defer outfile.Close()

	_, err = outfile.WriteString(mainCSS())
	return err
}

// mainCSS is the CSS for the main.css file. Currently embedded for the purpose
// of creating a single-file deployment with minimal fuss
func mainCSS() string {
	return `
:root {
    /* Define the color pallet */
    --border-dark: 2px solid #666666;
    --border-light: 2px solid #cccccc;
    --border-width: 2px;
    --line-width: 1px;
    --page-background: #f5f5f5;
    --primary-dark: #336791;
    --primary-med-dark: #4a7Fa9;
    --primary-medium: #008bb9;
    --zebra-one: #e6E6fa;
    --zebra-text: #030303;
    --zebra-two: #dadada;
}
body {
    background-color: var(--page-background);
    font-family: verdana,helvetica,sans-serif;
    margin: 0;
    padding: 0;
}

#topNav {
    overflow: hidden;
    background-color: var(--primary-medium);
}
#topNav a {
    background-color: var(--primary-dark);
    color: var(--page-background);
    border-left: var(--border-light);
    border-top: var(--border-light);
    border-bottom: var(--border-dark);
    border-right: var(--border-dark);
    float: left;
    padding: 2px 20px;
    text-align: center;
    text-decoration: none;
}
#topNav a:hover {
    background-color: var(--primary-med-dark);
    color: var(--zebra-one);
}
#topNav a.active {
    background-color: var(--primary-med-dark);
    color: var(--zebra-one);
}

#pageHead {
    background-color: var(--primary-medium);
    border-bottom: var(--border-dark);
    color: var(--page-background);
    margin-bottom: 10px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
    padding-top: 5px;
}
#pageHead h1 {
    border-bottom-style: solid;
    border-bottom-width: var(--line-width);
    color: var(--page-background);
    font-size: 150%;
    font-weight: bold;
    margin-top: 10px;
    padding-top: 0.5ex;
    padding-bottom: 0.5ex;
}
#pageHead table th {
    color: var(--page-background);
    text-align: left;
    vertical-align: top;
    white-space: nowrap;
}
#pageHead table tr {
    color: var(--page-background);
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
    color: var(--primary-dark);
    font-size: 130%;
}
#pageBody hr {
    margin-top: 15px;
    color: var(--primary-medium);
}

#pageFoot {
    background-color: var(--primary-medium);
    color: var(--page-background);

    border-left: var(--border-light);
    border-top: var(--border-light);
    border-bottom: var(--border-dark);
    border-right: var(--border-dark);

    font-size: 90%;
    font-weight: bold;
    margin-top: 10px;
    margin-left: 10px;
    margin-right: 10px;

    padding: 2px 7px 2px 2px;
}
#filter-form {
    vertical-align: top;
    color: var(--page-background);
    padding-bottom: 1.0ex;
    padding-top: 0.5ex;
}

table.dataTable {
    border-bottom: var(--border-dark);
    border-right: var(--border-dark);
    border-left: var(--border-light);
    border-top: var(--border-light);
    text-align: left;
    vertical-align: top;
    width: 100%;
}
table.dataTable thead tr th {
    background-color: #ddd;
    border-bottom: var(--border-dark);
    border-right: var(--border-dark);
    color: var(--primary-dark);
    /*cursor: pointer;*/
    padding-left: 4px;
    padding-right: 15px;
}
table.dataTable thead tr .headerSortDown {
    background-color: #bfbfbf;
    background-image: url("../img/desc.gif");
}
table.dataTable thead tr .headerSortUp {
    background-color: #bfbfbf;
    background-image: url("../img/asc.gif");
}
table.dataTable tbody tr:nth-child(odd) {
    background-color: var(--zebra-one);
    color: var(--zebra-text);
}
table.dataTable tbody tr:nth-child(even) {
    background-color: var(--zebra-two);
    color: var(--zebra-text);
}
table.dataTable tbody tr td {
    padding-left: 4px;
    padding-right: 4px;
    vertical-align: top;
}
table.dataTable tbody td.tcnw {
    white-space: nowrap;
}
table.dataTable tbody td.tcc {
    text-align: center;
}
table.dataTable tbody td.tcn {
    /* Numeric table cells. */
    padding-right: 8px;
    text-align: right;
    white-space: nowrap;
}
pre {
    background-color: var(--zebra-one);
    color: var(--zebra-text);
    border-left: var(--border-light);
    border-top: var(--border-light);
    border-bottom: var(--border-dark);
    border-right: var(--border-dark);
    padding: 5px;
    font-family: Bitstream Vera Sans Mono, Consolas, Courier New, DejaVu Sans Mono, Liberation Mono, Lucida Console, Monaco, monospace;
}
`
}
