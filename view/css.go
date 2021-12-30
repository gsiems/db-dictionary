package view

import (
	"os"

	m "github.com/gsiems/db-dictionary/model"
)

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

	_, err = outfile.WriteString(css())
	return err
}

func css() string {
	return `
body {
    background-color: white;
    font-family: verdana,helvetica,sans-serif;
    margin: 0;
    padding: 0;
}
h1 {
    border-bottom-style: solid;
    border-bottom-width: 2px;
    color: #eee;
    font-size: 150%;
    font-weight: bold;
    margin-top: 0;
    padding-bottom: 0.5ex;
}
h2 {
    border-bottom-style: solid;
    border-bottom-width: 2px;
    color: #009ace;
    font-size: 140%;
    padding-bottom: 0.5ex;
}
#NavBar ul {
    background-color: #003366;
    color: white;
    float: left;
    margin-left: 0;
    padding-left: 0;
    width: 100%;
}
#NavBar ul li {
    display: inline;
}
#NavBar ul li a {
    background-color: #003366;
    border-right: 1px solid #fff;
    color: white;
    float: left;
    padding: 0.2em 1em;
    text-decoration: none;
}
#NavBar ul li a:hover {
    background-color: #336699;
    color: #fff;
}
#PageHead {
    background-color: #009ace;
    border-bottom: 2px solid #999;
    border-color: #999;
    color: #eee;
    margin-bottom: 10px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
    padding-top: 5px;
}
#PageHead table th {
    color: #eee;
    text-align: left;
    vertical-align: top;
    white-space: nowrap;
}
#PageHead table tr {
    color: #eee;
    text-align: left;
    vertical-align: top;
}
#PageBody {
    font-size: 90%;
    margin-bottom: 10px;
    margin-left: 10px;
    margin-right: 10px;
}
#PageFoot {
    background-color: #009ace;
    color: #eee;
    border-bottom: 2px solid #999;
    border-top: 2px solid #999;
    font-size: 90%;
    font-weight: bold;
    margin-top: 10px;
    padding: 2px 7px 2px 2px;
    text-align: right;
}
#filter-form {
    color: #eee;
    padding-bottom: 1.0ex;
    padding-top: 0.5ex;
}
/* TC1: Standard table text cells */
.TC1 {
    padding-left: 4px;
    padding-right: 4px;
    vertical-align: top;
}
/* TC2: Non-wrapping table text cells */
.TC2 {
    padding-left: 4px;
    padding-right: 4px;
    vertical-align: top;
    white-space: nowrap;
}
/* TCc: Centered table text cells */
.TCc {
    padding-left: 4px;
    padding-right: 4px;
    text-align: center;
    vertical-align: top;
}
/* TCn: Numeric table cells. Note that this is probably backwards for RTL languages. */
.TCn {
    padding-left: 4px;
    padding-right: 8px;
    text-align: right;
    vertical-align: top;
    white-space: nowrap;
}
.TCcomment {
    padding-left: 8px;
    padding-right: 8px;
    text-align: left;
    vertical-align: top;
}
table.tablesorter {
    border-bottom: 1px solid #777;
    border-left: 1px solid #777;
    border-right: 1px solid #777;
    text-align: left;
    vertical-align: top;
    width: 100%;
}
table.tablesorter thead tr th {
    background-color: #ddd;
    background-position: right center;
    background-repeat: no-repeat;
    border-bottom: 2px solid #777;
    border-color: #777;
    color: #009ace;
    cursor: pointer;
    padding-left: 4px;
    padding-right: 15px;
}
table.tablesorter thead tr .headerSortDown {
    background-color: #bfbfbf;
    background-image: url("../img/desc.gif");
}
table.tablesorter thead tr .headerSortUp {
    background-color: #bfbfbf;
    background-image: url("../img/asc.gif");
}
table.tablesorter tr:nth-child(odd) {
    background-color: #eee;
    color: #333;
}
table.tablesorter tr:nth-child(even) {
    background-color: #ddd;
    color: #333;
}
pre {
    background-color: #eee;
    color: #333;
    padding: 5px;
    font-family: Bitstream Vera Sans Mono, Consolas, Courier New, DejaVu Sans Mono, Liberation Mono, Lucida Console, Monaco, monospace;
}
`
}
