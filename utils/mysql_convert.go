package utils

import (
	"database/sql"
	"strconv"
)

func GetNumberOfRowsAffected(result sql.Result) int64 {

	rowsAffected, err := result.RowsAffected()

	CheckError(err)

	return rowsAffected
}

func StringNumberOfRowsAffected(result sql.Result) string {
	string := strconv.FormatInt(GetNumberOfRowsAffected(result), 10)

	return string
}

func GetLastInsertedID(result sql.Result) int64 {
	lastInsertedID, err := result.LastInsertId()

	CheckError(err)

	return lastInsertedID
}

func StringLastInsertedID(result sql.Result) string {
	string := strconv.FormatInt(GetLastInsertedID(result), 10)

	return string
}
