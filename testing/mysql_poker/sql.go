package mysql_poker

import (

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils"
)

// Deletes all entries in players table and re-inserts
// players necessary for running tests
func RefreshPlayerTable( DB *mysql_db.DB ) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := `
	
	DELETE FROM cs2208_jr30.dummy_players;

    INSERT INTO cs2208_jr30.dummy_players 
	VALUES 
		("derek", 30.0, 1, 1, "PLAYING", "2h3h", 0.0, CURRENT_TIMESTAMP()),
		("jason", 30.0, 1, 2, "PLAYING", "4h5h", 0.0, CURRENT_TIMESTAMP()),
		("john", 30.0, 1, 3, "PLAYING", "6h7h", 0.0, CURRENT_TIMESTAMP()),
		("barry", 30.0, 1, 4, "PLAYING", "8h9h", 0.0, CURRENT_TIMESTAMP()),
		("ahmed", 30.0, 1, 5, "PLAYING", "2d3d", 0.0, CURRENT_TIMESTAMP()),
		("laura", 30.0, 1, 6, "PLAYING", "4d5d", 0.0, CURRENT_TIMESTAMP()),
		("alejandro", 30.0, 1, 7, "PLAYING", "TsJs", 0.0, CURRENT_TIMESTAMP()),
		("dan", 30.0, 1, 8, "PLAYING", "6d7d", 0.0, CURRENT_TIMESTAMP());
		    `

	_, err := db.Exec(query)
	utils.CheckError(err)

}

// Deletes all entries in poker table and re-inserts
// tables necessary for running tests
func RefreshPokerTable( DB *mysql_db.DB ) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := ` DELETE FROM cs2208_jr30.dummy_poker_tables;
	
	           INSERT INTO cs2208_jr30.dummy_poker_tables
			   VALUES (1, "QsKsAs", "john", 1.0, "derek", 1.0);
			 `

	_, err := db.Exec(query)
	utils.CheckError(err)
}

// Deletes all entries in tables table and re-inserts
// tables necessary for running tests
func RefreshTablesTable( DB *mysql_db.DB ) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := ` 
	DELETE FROM cs2208_jr30.dummy_tables;

	INSERT INTO cs2208_jr30.dummy_tables 
	VALUES (1,
	        DATE_SUB(NOW(),
			INTERVAL 48 HOUR),
			"barry",
	        "Ah2h3h4h5h6h7h8h9hThJhQhKhAd2d3d4d5d6d7d8d9dTdJdQdKdAs2s3s4s5s6s7s8s9sTsJsQsKsAc2c3c4c5c6c7c8c9cTcJcQcKc",
			"",
			true);
			 `

	_, err := db.Exec(query)
	utils.CheckError(err)
}