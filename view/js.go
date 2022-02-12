package view

import (
	"path"

	m "github.com/gsiems/db-dictionary/model"
)

// makeJS generates the needed javascript file(s)
func makeJS(md *m.MetaData) (err error) {

	dirName := path.Join(md.OutputDir, "js")
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

// writeDefaultJS writes the default javascript file(s)
func writeDefaultJS(dirName string) (err error) {

	err = writeFile(path.Join(dirName, "filter.js"), tableFilter())
	if err != nil {
		return err
	}

	err = writeFile(path.Join(dirName, "sort.js"), tableSorter())
	return err
}

// tableFilter returns the javascript used for providing the data filtering
// functionality on the various data dictionary pages
func tableFilter() string {
	return `
function filterTables() {

  var input, filter, table, tbody, tr, td, i, j, k, txtValue, newDisplay;

  input = document.getElementById('filterBy');
  filter = input.value.toUpperCase();

  table = document.getElementsByClassName('dataTable');
  if (table) {
    for (i = 0; i < table.length; i++) {

      tbody = table[i].getElementsByTagName('tbody');
      if (tbody) {

        tr = tbody[0].getElementsByTagName('tr');
        for (j = 0; j < tr.length; j++) {

          td = tr[j].getElementsByTagName('td');
          if (td) {
            newDisplay = 'none'
            for (k = 0; k < td.length; k++) {
              if (td[k]) {
                txtValue = td[k].textContent || td[k].innerText;
                if (txtValue.toUpperCase().indexOf(filter) > -1) {
                  newDisplay = '';
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
}
`
}

// tableSorter returns the javascript used for providing the basic table sorting
// functionality on the various data dictionary pages
func tableSorter() string {
	return `
function sortTable(table, e) {

  var th

  if (e.target.tagName != 'TH') {
    return;
  }

  th = e.target;

  if (th.classList.contains('headerSortUp')){
    sortTableDesc(table, th.cellIndex, th.dataset.type);
    setSortMarkers(table, th, 'down')
  } else {
    sortTableAsc(table, th.cellIndex, th.dataset.type);
    setSortMarkers(table, th, 'up')
  }
}

function setSortMarkers(table, sorted, direction) {

  var thead, tr, th

  thead = table.getElementsByTagName('thead');
  if (thead) {

    tr = thead[0].getElementsByTagName('tr');
    if (tr) {

      th = tr[0].getElementsByTagName('th');
      if (th) {

        for (i = 0; i < th.length; i++) {
          if(th[i].classList.contains('headerSortUp'))
            th[i].classList.remove('headerSortUp') ;
          if(th[i].classList.contains('headerSortDown'))
            th[i].classList.remove('headerSortDown') ;
        }
      }
    }
  }

  switch (direction) {
  case 'up':
    sorted.classList.add('headerSortUp');
    break;
  case 'down':
    sorted.classList.add('headerSortDown');
    break;
  }
}

function sortTableDesc(table, idx, type) {

  var tbody, rowAry, compare

  tbody = table.querySelector('tbody');
  if (tbody) {
    rowAry = Array.from(tbody.rows);

    switch (type) {
    case 'number':
      compare = function(rowA, rowB) {
        return rowB.cells[idx].innerHTML - rowA.cells[idx].innerHTML;
      };
      break;
    default:
      compare = function(rowA, rowB) {
        return rowB.cells[idx].innerHTML > rowA.cells[idx].innerHTML ? 1 : -1;
      };
      break;
    }

    rowAry.sort(compare);
    tbody.append(...rowAry);
  }
}

function sortTableAsc(table, idx, type) {

  var tbody, rowAry, compare

  tbody = table.querySelector('tbody');
  if (tbody) {

    rowAry = Array.from(tbody.rows);

    switch (type) {
    case 'number':
      compare = function(rowA, rowB) {
        return rowA.cells[idx].innerHTML - rowB.cells[idx].innerHTML;
      };
      break;
    default:
      compare = function(rowA, rowB) {
        return rowA.cells[idx].innerHTML > rowB.cells[idx].innerHTML ? 1 : -1;
      };
      break;
    }

    rowAry.sort(compare);
    tbody.append(...rowAry);
  }
}
`
}
