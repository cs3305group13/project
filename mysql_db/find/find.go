package find

import (
	"database/sql"
	"fmt"
)

// Function returns value from table row that matches the required value under the specified column.
//
// All string parameters are case sensitive.
func FindRowByValue(tx *sql.Tx, tableName, compareColumnName, valueToBeCompared, returnColumnName string) string {

	var query = fmt.Sprintf(string(`SELECT %s 
	                                FROM %s 
				                    WHERE %s = ?;`), returnColumnName, tableName, compareColumnName)

	var fieldFound string
	err := tx.QueryRow(query, valueToBeCompared).Scan(&fieldFound)

	if err != nil {
		if err.Error() == "ErrNoRows" {
			fieldFound = ""
			return fieldFound
		}
	}

    return fieldFound
}