package gameobserver

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/utils"
)

// Function monitors if any players go idle. Should be use with goroutine call.
//
// IMPORTANT - dont test
func GameObserver(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {

	numOfPlayers := gameinfo.GetNumberOfPlayersAtTable( DB, playersTableName, tableID )  // check how many players are in game.
	db := mysql_db.EstablishConnection(DB)
	var tx *sql.Tx
	defer tx.Rollback()
	defer db.Close()
	
	for numOfPlayers != 0 {
		tx = mysql_db.NewTransaction(db)
		
		removeIdleUsers(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID )
		
		
		err := tx.Commit()
		utils.CheckError(err)

		numOfPlayers = gameinfo.GetNumberOfPlayersAtTable( DB, playersTableName, tableID )  // refresh numOfPlayers
		
		time.Sleep(time.Second)
		fmt.Println("slept one second")
	}
}


func removeIdleUsers(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {

	fiveSeconds := "5"

	currentPlayerMakingMove, currentPlayerMakingMoveSeatNumber := gameinfo.GetCurrentPlayerMakingMove(DB, tablesTableName, playersTableName, tableID)
	highestBidder, highestBidderSeatNumber, dealer, dealerSeatNumber := gameinfo.GetDealerAndHighestBidder(DB, playersTableName, pokerTablesTableName, tableID)

	if findIdleImportantUser(DB, playersTableName, tableID, fiveSeconds, highestBidder) != "" {
		setOperation := "highest_bidder = "
		gameflow.SetNextAvailablePlayerAfterThisOne(DB, tx, pokerTablesTableName, playersTableName, tableID, highestBidder, highestBidderSeatNumber, setOperation)
	}
	if findIdleImportantUser(DB, playersTableName, tableID, fiveSeconds, dealer) != "" {
		setOperation := "dealer = "
		gameflow.SetNextAvailablePlayerAfterThisOne(DB, tx, pokerTablesTableName, playersTableName, tableID, dealer, dealerSeatNumber, setOperation)
	}
	if findIdleImportantUser(DB, playersTableName, tableID, fiveSeconds, currentPlayerMakingMove) != "" {
		setOperation := "current_player_making_move = "
		gameflow.SetNextAvailablePlayerAfterThisOne(DB, tx, tablesTableName, playersTableName, tableID, currentPlayerMakingMove, currentPlayerMakingMoveSeatNumber, setOperation)
	}

	query := fmt.Sprintf(`UPDATE %s
	                      SET table_id = "0",
						      seat_number = "0",
							  player_state = "LEFT",
							  player_cards = "",
							  money_in_pot = 0.0
						  WHERE table_id = %s AND time_since_request < DATE_SUB(NOW(), INTERVAL %s SECOND);`, playersTableName, tableID, fiveSeconds)
	res, err := tx.Exec(query)

	if err != sql.ErrNoRows {
	    utils.CheckError(err)
	}

	numberOfUsersRemoved := utils.GetNumberOfRowsAffected(res)
	if numberOfUsersRemoved > 0 {
		fmt.Printf("%v user(s) were removed\n", numberOfUsersRemoved)
	}
}

func findIdleImportantUser(DB *mysql_db.DB, playersTableName, tableID, fiveSeconds, usernameAtRole string) (username string) {
	
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT username
	                      FROM %s
	                      WHERE table_id = %s AND 
		                        time_since_request < DATE_SUB(NOW(), INTERVAL %s SECOND) AND
		                        username = "%s" ;`, playersTableName, tableID, fiveSeconds, usernameAtRole)

	err := db.QueryRow(query).Scan(&username)
	
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return username
}