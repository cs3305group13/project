package gamejoin

import (
	"database/sql"
	"fmt"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/mysql_db/find"
	"github.com/cs3305/group13_2022/project/mysql_db/insert"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
)

// ############# TABLE JOIN STUFF BELOW. ##############

func CheckTableExists( tx *sql.Tx, tablesTableName, tableCode string ) bool {
	tableExists := find.FindRowByValue(tx, tablesTableName, "table_id", tableCode, "table_id")

	if tableExists == "" {
		return false
	} 

	// table exists
	return true
}

func UpdatePlayersSelectedGame(DB *mysql_db.DB, tx *sql.Tx, playersTableName, tableID, username, seatNumber string) (funds string) {
	userFunds := gameinfo.GetPlayersFunds(DB, playersTableName, username)

	if userFunds < 5.0 {
		userFunds = 30.0
	}
	userFundsString := fmt.Sprintf("%.2f", userFunds)  // convert float64 to string

	playerState := "NOT_READY"
	resetPlayerCards := ""
	moneyInPot := "0.0"

	columnNames := "username, funds, table_id, seat_number, player_state, player_cards, money_in_pot, time_since_request"
	values := fmt.Sprintf(`"%s", "%s", "%s", "%s", "%s", "%s", "%s", CURRENT_TIMESTAMP()`, username, userFundsString, tableID, seatNumber, playerState, resetPlayerCards, moneyInPot)
	_, inserted := insert.InsertTableEntry(tx, playersTableName, columnNames, values)
	if ! inserted {
		panic("Player was not inserted into players table.")
	}
	
	return userFundsString
}


// Used for add ing player to players table when signing up
func AddPlayer(tx *sql.Tx, playersTableName, username, funds string) bool {

	columnNames := `username, funds, table_id, seat_number, player_state, player_cards, money_in_pot, time_since_request`

	tableID := "0"
	seatNumber := "0"
	playerState := "LEFT"
	playerCards := ""
	moneyInPot := "0.0"
	values := fmt.Sprintf(`"%s", "%s", "%s" , "%s", "%s", "%s", "%s", CURRENT_TIMESTAMP`, username, funds, tableID, seatNumber, playerState, playerCards, moneyInPot)

	_, inserted := insert.InsertTableEntry(tx, playersTableName, columnNames, values)

	return inserted
}