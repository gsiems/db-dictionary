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
    --border-dark: 2px solid #808080; /* Grey */
    --border-light: 2px solid #C0C0C0; /* Silver */
    --border-width: 2px;
    --line-width: 1px;
    --page-base: #F5F5F5; /* WhiteSmoke */
    --blue-1: #008BB9; /* lighter blue */
    --blue-2: #4A7FA9; /* blue */
    --blue-3: #336791; /* pg blue */
    --blue-4: #483D8B; /* DarkSlateBlue */
    --grey-1: #D3D3D3; /* LightGrey */
    --grey-2: #C0C0C0; /* Silver */
    --grey-3: #DCDCDC; /* Gainsboro */
    --zebra-odd: #E6E6FA; /* Lavender */
    --zebra-even: #DCDCDC; /* Gainsboro */
    --zebra-text: #000000;
    --form-text: #000000;
}
body {
    background-color: var(--page-base);
    font-family: verdana,helvetica,sans-serif;
    margin: 0;
    padding: 0;
}
#topNav {
    overflow: hidden;
    background-color: var(--blue-1);
}
#topNav a {
    background-color: var(--blue-3);
    color: var(--page-base);
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
    background-color: var(--blue-1);
    color: var(--page-base);
}
#topNav a.active {
    background-color: var(--blue-1);
    color: var(--page-base);
}
#pageHead {
    background-color: var(--blue-1);
    border-bottom: var(--border-dark);
    color: var(--page-base);
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
    color: var(--blue-1);
    font-size: 130%;
}
#pageBody hr {
    margin-top: 10px;
    color: var(--blue-1);
}
#pageFoot {
    background-color: var(--blue-1);
    color: var(--page-base);
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
#pageFoot a {
    color: var(--blue-4);
}
#filter-form {
    vertical-align: top;
    /*color: var(--blue-2);*/
    padding-bottom: 1.0ex;
    padding-top: 0.5ex;
}
#filterBy {
    background-color: var(--page-base);
    color: var(--form-text);
    border-bottom: var(--border-light);
    border-right: var(--border-light);
    border-left: var(--border-dark);
    border-top: var(--border-dark);
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
    background-color: var(--grey-3);
    border-bottom: var(--border-dark);
    border-right: var(--border-dark);
    color: var(--blue-3);
    /*cursor: pointer;*/
    padding-left: 4px;
    padding-right: 15px;
}
table.dataTable thead tr .headerSortDown {
    background-color: var(--grey-2);
    background-image: url("../img/desc.gif");
}
table.dataTable thead tr .headerSortUp {
    background-color: var(--grey-2);
    background-image: url("../img/asc.gif");
}
table.dataTable tbody tr:nth-child(odd) {
    background-color: var(--zebra-odd);
    color: var(--zebra-text);
}
table.dataTable tbody tr:nth-child(even) {
    background-color: var(--zebra-even);
    color: var(--zebra-text);
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
    background-color: var(--zebra-odd);
    color: var(--zebra-text);
    border-left: var(--border-light);
    border-top: var(--border-light);
    border-bottom: var(--border-dark);
    border-right: var(--border-dark);
    padding: 5px;
    font-size: 100%;
    font-family: Bitstream Vera Sans Mono, Consolas, Courier New, DejaVu Sans Mono, Liberation Mono, Lucida Console, Monaco, monospace;
}
`
}
