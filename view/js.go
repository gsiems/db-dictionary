package view

import (
	"os"

	m "github.com/gsiems/db-dictionary-core/model"
)

// makeJS generates the needed javascript file(s)
func makeJS(md *m.MetaData) (err error) {

	dirName := md.OutputDir + "/js"
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0744)
		if err != nil {
			return err
		}
	}

	outfile, err := os.Create(dirName + "/filter.js")
	if err != nil {
		return err
	}
	defer outfile.Close()

	_, err = outfile.WriteString(tableFilter())
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
  - single column or multi-column sort?
  - if multi then need to store array of sorted columns
    - click plus modifier? (shift, ctrl, ?) to indicate appending column
      - if is already the last column in array then toggle sort for the column
      - if is already in the list, but is not last, then... ignore?
    - click without modifier, reset sorted columns array

*/
