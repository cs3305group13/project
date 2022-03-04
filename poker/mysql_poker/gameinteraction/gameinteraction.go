package gameinteraction

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameshowdown"
	"github.com/cs3305/group13_2022/project/utils"
)




func TryTakeMoneyFromPlayer(DB *mysql_db.DB, tx *sql.Tx, playersTableName, pokerTablesTableName, tableID, playerName, bid string) (taken bool) {
	playersFunds := gameinfo.GetPlayersFunds(DB, playersTableName, playerName)
	
	playersBid, err := strconv.ParseFloat(bid, 64)
	utils.CheckError(err)

	if playersFunds < playersBid {
		taken = false
		return taken
	}

	query := fmt.Sprintf(`UPDATE %s
	                      SET funds = funds - %v,
						      money_in_pot = money_in_pot + %v
						  WHERE username = "%s";`, playersTableName, playersBid, playersBid, playerName)
	_, err = tx.Exec(query)
	utils.CheckError(err)

	query = fmt.Sprintf(`UPDATE %s
	                     SET money_in_pot = money_in_pot + %v
						 WHERE table_id = %s;`, pokerTablesTableName, bid, tableID)

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


func PlayerFolded(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber string, nextPlayerFoundBool bool) {
	
	numberOfPlayersStillPlaying := gameinfo.GetNumberOfPlayersStillPlaying(DB, playersTableName, tableID, username, seatNumber)
	// ^ contains number of players still in game (this player who wants to fold, players still playing, and all in players)

	if ! nextPlayerFoundBool && numberOfPlayersStillPlaying == 1{
		// no one is all in and this is last player not folded so cannot let them fold
		// give player pot
		return
	} else if ! nextPlayerFoundBool && numberOfPlayersStillPlaying > 1{
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

func PlayerTakesAction(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, amount string) (action string) {
	amountAsFloat64 := utils.ConvertToFloat(amount)
	playersMoneyCurrentlyInPot := gameinfo.GetPlayersMoneyInPot(DB, playersTableName, username)

	playersMoneyInPot := amountAsFloat64 + playersMoneyCurrentlyInPot

	_, highestBid := gameinfo.GetHighestBidder(DB, pokerTablesTableName, tableID)
	
	playersFunds := gameinfo.GetPlayersFunds(DB, playersTableName, username)

	setOperation := "current_player_making_move = "
	successful := gameflow.SetNextAvailablePlayerAfterThisOne(DB, tx, tablesTableName, playersTableName, tableID, username, seatNumber, setOperation)


	if amountAsFloat64 >= playersFunds {
		playerAllIn(tx, playersTableName, username)
		action = "ALL_IN"

	} else if highestBid == 0.0 && playersMoneyInPot == 0.0 {
		playerChecked(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username)
		action = "CHECKED"

	} else if playersMoneyInPot < highestBid {
		PlayerFolded(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, successful)
		action = "FOLDED"

		return action
	
	} else if playersMoneyInPot < highestBid*2 {
		// if here raise amount was not at least double of previous highest bid
		playerCalled(tx, playersTableName, username)
		action = "CALLED"


	} else if playersMoneyInPot >= highestBid*2 {
		playerRaised(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, amount)
		action = "RAISED"

	}

	
	successfullyTaken := TryTakeMoneyFromPlayer(DB, tx, playersTableName, pokerTablesTableName, tableID, username, amount)

	if ! successfullyTaken {
		panic("Amount was not taken properly.")
	}

	return action
}


func playerAllIn(tx *sql.Tx, playersTableName, username string) {
	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "ALL_IN"
						  WHERE username = "%s"`, playersTableName, username)

	_, err := tx.Exec(query)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

}

func playerRaised( DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, raiseAmount string ) bool {

	setOperation := fmt.Sprintf(`highest_bidder = "%s",
	                             highest_bid = "%s"`, username, raiseAmount)

	gameflow.AssignThisPlayerToRole(tx, pokerTablesTableName, tableID, username, setOperation)

	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "RAISED"
						  WHERE username = "%s";`, playersTableName, username)

	_, err := tx.Exec(query)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return true
}

func playerCalled( tx *sql.Tx, playersTableName, username string ) bool {

	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "CALLED"
						  WHERE username = "%s";`, playersTableName, username)

	_, err := tx.Exec(query)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return true
}

func playerChecked( DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username string ) bool {

	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "CHECKED"
						  WHERE username = "%s";`, playersTableName, username)

	_, err := tx.Exec(query)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return true
}