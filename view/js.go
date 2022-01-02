package view

import (
	"os"

	m "github.com/gsiems/db-dictionary/model"
)

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
