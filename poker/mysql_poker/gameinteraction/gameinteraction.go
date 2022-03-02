package gameinteraction

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameshowdown"
	"github.com/cs3305/group13_2022/project/utils"
)

func GameInProgress(w http.ResponseWriter, DB *mysql_db.DB, tablesTableName, tableID string) bool {
	
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT game_in_progress
	                      FROM %s
						  WHERE table_id = "%s"`, tablesTableName, tableID)

	var gameState bool
	err := db.QueryRow(query).Scan(&gameState)
	
	utils.CheckError(err)

	if gameState {
		w.Write([]byte("MESSAGE:\nSorry game is in progress, should be over soon."))
	}

	return gameState
}

	if highestBid < playersBid {
		highestBidder = playerName
		highestBid = playersBid
		query := fmt.Sprintf(`UPDATE %s 
	                          SET highest_bidder = "%s",
						          highest_bid = %v
						      WHERE table_id = %s;`, pokerTablesTableName, highestBidder, highestBid, tableID)
		
func TryTakeMoneyFromPlayer(DB *mysql_db.DB, tx *sql.Tx, playersTableName, tableID, playerName, bid string) (taken bool) {
	playersFunds := gameinfo.GetPlayersFunds(DB, playersTableName, playerName)
	
	playersBid, err := strconv.ParseFloat(bid, 64)
		utils.CheckError(err)

	if playersFunds < playersBid {
		taken = false
		return taken
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

func PlayersTurn(DB *mysql_db.DB, tablesTableName, playersTableName, tableID, username string) bool {

	currentPlayerMakingMove, _ := gameinfo.GetCurrentPlayerMakingMove(DB,tablesTableName, playersTableName, tableID)

	if username == currentPlayerMakingMove {
		return true
	}

	return false
}


func PlayerFolded(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber string) {
	
	setOperation := "current_player_making_move = "
	successful := gameflow.SetNextAvailablePlayerAfterThisOne(DB, tx, tablesTableName, playersTableName, tableID, username, seatNumber, setOperation)

	numberOfPlayersStillPlaying := gameinfo.GetNumberOfPlayersStillPlaying(DB, playersTableName, tableID, username, seatNumber)
	// ^ contains number of players still in game (this player who wants to fold, players still playing, and all in players)

	if ! successful && numberOfPlayersStillPlaying == 1{
		// no one is all in and this is last player not folded so cannot let them fold
		// give player pot
		return
	} else if ! successful && numberOfPlayersStillPlaying > 1{
		gameshowdown.ShowDown(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)
	}else {
		// if here then there are still other players playing therefore this player can fold

		query := fmt.Sprintf(`
							UPDATE %s
							SET player_state = "FOLDED"
							WHERE username = "%s"`, playersTableName, username)
							
		response, err := tx.Exec(query)
		utils.CheckError(err)

		numberOfRows := utils.GetNumberOfRowsAffected(response)

		if numberOfRows != 1 {
			panic("One and only one row should have been changed with this operation.")
		}
	} 
}


}


func PlayerRaised(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName, playersTableName, raiseAmount string) {
	
}