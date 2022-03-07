package mysql_db

import (
	"database/sql"
	"strconv"

	"github.com/cs3305/group13_2022/project/utils"
)

func GetNumberOfRowsAffected( result sql.Result ) int64 {

	rowsAffected, err := result.RowsAffected()
	
	utils.CheckError(err)

	return rowsAffected
}

func StringNumberOfRowsAffected( result sql.Result ) string {
	string := strconv.FormatInt( GetNumberOfRowsAffected(result), 10 )

	return string 
}

func GetLastInsertedID( result sql.Result ) int64 {
	lastInsertedID, err := result.LastInsertId()

	utils.CheckError(err)

	return lastInsertedID
}

func StringLastInsertedID( result sql.Result ) string {
	string := strconv.FormatInt( GetLastInsertedID(result), 10 )

	return string 
}