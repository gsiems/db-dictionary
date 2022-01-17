package view

import (
	m "github.com/gsiems/db-dictionary-core/model"
)

// makeJS generates the needed javascript file(s)
func makeJS(md *m.MetaData) (err error) {

	dirName := md.OutputDir + "/js"
	err = ensurePath(dirName)
	if err != nil {
		return err
	}

	err = writeDefaultJS(dirName)
	if err != nil {
		return err
	}

	err = copyFileList(dirName, md.Cfg.JSFiles)

	return err
}

func writeDefaultJS(dirName string) (err error) {

	err = writeFile(dirName+"/filter.js", tableFilter())

	return err
}

// tableFilter returns the javascript used for providing the data filtering
// functionality on the various data dictionary pages
func tableFilter() string {
	return `
function filterTables() {
  var input, filter, table, tbody, tr, td, i, j, k, txtValue, newDisplay;
  input = document.getElementById("filterBy");
  filter = input.value.toUpperCase();

  table = document.getElementsByClassName("dataTable");
  for (i = 0; i < table.length; i++) {

    tbody = table[i].getElementsByTagName("tbody");
    if (tbody) {

      tr = tbody[0].getElementsByTagName("tr");
      for (j = 0; j < tr.length; j++) {

        td = tr[j].getElementsByTagName("td");
        if (td) {
          newDisplay = "none"
          for (k = 0; k < td.length; k++) {
            if (td[k]) {
              txtValue = td[k].textContent || td[k].innerText;
              if (txtValue.toUpperCase().indexOf(filter) > -1) {
                newDisplay = "";
              }
            }
          }
          if (tr[j].style.display != newDisplay) {
            tr[j].style.display = newDisplay;
          }
        }
      }
    }
  }
}
`

}

////////////////////////////////////////////////////////////////////////////////

/*
Sorting data:

One possible? approach:

  - a page may have multiple tables that can be sorted
  - click on cell header to toggle sort
  - sort should be aware of the datatype presented in a table column (text, numeric)
    - something like: td[i].getAttribute("class") to determine datatype for sorting
      (assert that number, date, etc. columns have class name that indicates the datatype)
      - means needing different <th><td> CSS for different datatypes
  - set class of table header cell to indicate sort order {asc, desc, none}
  - for header cell, if clicked:
    - was shift-clicked?
      - yes: is the column in the list of already sorted columns?
        - yes: is the column the last column in the list?
          - yes: toggle the sort of the column {asc, desc, none}
          - no: ignore the click
        - no: append the column to the list
      - no: (just clicked)
        - reset the columns list and append the clicked column
    - read whatever attributes to determine sort order
    - sort columns

*/
