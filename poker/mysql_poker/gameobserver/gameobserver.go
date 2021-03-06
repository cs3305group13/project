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
// IMPORTANT - test takes 5seconds because of condition in mysql query
func GameObserver(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {

	numOfPlayers := gameinfo.GetNumberOfPlayersAtTable( DB, playersTableName, tableID )  // check how many players are in game.
	
	for numOfPlayers != 0 {
		
		removeIdleUsers(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID )

		numOfPlayers = gameinfo.GetNumberOfPlayersAtTable( DB, playersTableName, tableID )  // refresh numOfPlayers
		
		time.Sleep( time.Second )
	}
}


func removeIdleUsers(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {

	fiveSeconds := "5"

	currentPlayerMakingMove, currentPlayerMakingMoveSeatNumber := gameinfo.GetCurrentPlayerMakingMove(DB, tablesTableName, playersTableName, tableID)
	highestBidder, highestBidderSeatNumber, dealer, dealerSeatNumber := gameinfo.GetDealerAndHighestBidder(DB, playersTableName, pokerTablesTableName, tableID)

	if findIdleImportantUser(DB, playersTableName, tableID, fiveSeconds, highestBidder) != "" {
		setOperation := "highest_bidder = "
		gameflow.SetNextAvailablePlayerAfterThisOne(DB, pokerTablesTableName, playersTableName, tableID, highestBidder, highestBidderSeatNumber, setOperation)
	}
	if findIdleImportantUser(DB, playersTableName, tableID, fiveSeconds, dealer) != "" {
		setOperation := "dealer = "
		gameflow.SetNextAvailablePlayerAfterThisOne(DB, pokerTablesTableName, playersTableName, tableID, dealer, dealerSeatNumber, setOperation)
	}
	if findIdleImportantUser(DB, playersTableName, tableID, fiveSeconds, currentPlayerMakingMove) != "" {
		setOperation := "current_player_making_move = "
		gameflow.SetNextAvailablePlayerAfterThisOne(DB, tablesTableName, playersTableName, tableID, currentPlayerMakingMove, currentPlayerMakingMoveSeatNumber, setOperation)
	}

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET table_id = "0",
						      seat_number = "0",
							  player_state = "LEFT",
							  player_cards = "",
							  money_in_pot = 0.0
						  WHERE table_id = %s AND time_since_request < DATE_SUB(NOW(), INTERVAL %s SECOND);`, playersTableName, tableID, fiveSeconds)
	db.Exec(query)
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