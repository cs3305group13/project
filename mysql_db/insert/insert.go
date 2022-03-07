package insert

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/cs3305/group13_2022/project/mysql_db"
)

//
// columnNames - list of column names in order of occurrence
// values - list of values in order matching that of columnNames
// update - a list consisting of a mapping between columnName and value
//
// example:
// columnNames = "id, a, b, c"
// values = '"1", "2", "3", "4"'
//
func InsertTableEntry(tx *sql.Tx, tableName, columnNames, values string) (tableID int64, inserted bool) {
	var updateOperation string

	var columnNamesArray = strings.Split(columnNames, ",")
	var valuesArray = strings.Split(values, ",")

	for i:=0; i<len(columnNamesArray) && i<len(valuesArray); i++ {
		updateOperation += fmt.Sprintf(`%s = %s`, columnNamesArray[i], valuesArray[i])
		if i+1 < len(columnNamesArray) && i+1 < len(valuesArray) {
			updateOperation += ", "
		}
	}

	var query = fmt.Sprintf(`INSERT INTO %s (%s)
	                         VALUES (%s)
							 ON DUPLICATE KEY UPDATE %s;`, tableName, columnNames, values, updateOperation)


	res, err := tx.Exec(query)
	
	if err != nil {
		panic(err)
	}
	// numOfRowsAffected := utils.GetNumberOfRowsAffected(res)

	// if numOfRowsAffected != 1 {
	// 	return tableID, false
	// }

	tableID = mysql_db.GetLastInsertedID(res)

	return tableID, true
}