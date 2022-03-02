package gameinteraction

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/utils"
)

func GameInProgress(w http.ResponseWriter, r *http.Request, tx *sql.Tx, tablesTableName string) bool {
	
	return false
}


func TryTakeMoneyFromPlayer(DB *mysql_db.DB, tx *sql.Tx, playersTableName, pokerTablesTableName, tableID, playerName, bid string) (taken bool) {
	playersFunds := gameinfo.GetPlayersFunds(tx, playersTableName, playerName)
	
	playersBid, err := strconv.ParseFloat(bid, 64)
	utils.CheckError(err)

	if playersFunds < playersBid {
		taken = false
		return taken
	}

	highestBidder, highestBid := gameinfo.GetHighestBidder(DB, pokerTablesTableName, tableID)

	if highestBid < playersBid {
		highestBidder = playerName
		highestBid = playersBid
		query := fmt.Sprintf(`UPDATE %s 
	                          SET highest_bidder = "%s",
						          highest_bid = %v
						      WHERE table_id = %s;`, pokerTablesTableName, highestBidder, highestBid, tableID)
		
        res, err := tx.Exec(query)
		utils.CheckError(err)

		rowsAffected := utils.GetNumberOfRowsAffected(res)
		if rowsAffected != 1 {
			panic("exactly one row should have been affected")
		}
	}

	query := fmt.Sprintf(`UPDATE %s
	                      SET funds = funds - %v
						  WHERE username = "%s";`, playersTableName, playersBid, playerName)
	res, err := tx.Exec(query)
	utils.CheckError(err)

	rowsAffected := utils.GetNumberOfRowsAffected(res)
	if rowsAffected != 1 {
		panic("exactly one row should have been affected")
	}

	taken = true
	return taken
}

func PlayersTurn(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName string) bool {

	return false
}


func PlayerFolded(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName string) {

}


func PlayerRaised(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName, playersTableName, raiseAmount string) {
	
}