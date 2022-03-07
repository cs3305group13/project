package change

import (
	"database/sql"
	"fmt"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils"
)

// Function changes rows column values using the values in the `operation` parameter which is passed to a `SET` MySQL query.
//
// compareColumnName - A column which uniquely identifies a singular row entry.
// valueToBeCompared - A value which uniquely identifies the row entry.
// operation - An operation which specifies the column names and column values which should be inserted into the row entry.
//
// Example operation input:
//
//		operation := fmt.Sprintf(`funds = "%s",
//								  table_id = "%s",
//		                          seat_number = "%s",
//		                          player_cards = "%s"`, 30.0, 7385492, 2, resetPlayerCards)
//
//
func ChangeValueInRows(tx *sql.Tx, tableName, compareColumnName, valueToBeCompared, operation string) (numOfRowsAffected int64, changed bool) {

	var query_1 = fmt.Sprintf(`UPDATE %s
	                           SET %s
				               WHERE %s = "%s";`, tableName, operation, compareColumnName, valueToBeCompared)
	
	res, err := tx.Exec(query_1)
	
	utils.CheckError(err)

	numOfRowsAffected = mysql_db.GetNumberOfRowsAffected(res)

	if numOfRowsAffected == 0 {
		return numOfRowsAffected, false
	}

	return numOfRowsAffected, true
}

// Function attempts to reset an entry which is older than the specified hours since now, if not found returns 0 otherwise returns rowID
// 
// The table this function runs on must have a unique primary key.
func CreateRowFromRowOlderThan(tx *sql.Tx, tableName, primaryKey, timeColumn, hours, setOperation string) (rowID int64, found bool) {

	query := fmt.Sprintf(`SELECT %s
	                      FROM %s
						  WHERE  %s <= DATE_SUB(NOW(), INTERVAL %v HOUR)
						  ORDER BY %s LIMIT 1;`, primaryKey, tableName, timeColumn, hours, timeColumn)
	
	err := tx.QueryRow(query).Scan(&rowID)
	if err == sql.ErrNoRows {

	    return 0, false
	} else {
	    utils.CheckError(err)
	}

	var execQuery = fmt.Sprintf(`UPDATE %s
                                 SET %s = CURRENT_TIMESTAMP(), 
		                             %s
	                             WHERE %s <= DATE_SUB(NOW(), INTERVAL %v HOUR) 
	                             ORDER BY %s LIMIT 1;`, tableName, timeColumn, setOperation, timeColumn, hours, timeColumn)


	result, err := tx.Exec(execQuery)
	utils.CheckError(err)
	
	numberOfRowsAffected := mysql_db.GetNumberOfRowsAffected(result)
	if numberOfRowsAffected != 1 {
		panic("One row should have been affected here.")
	}
	utils.CheckError(err)

	return rowID, true
}