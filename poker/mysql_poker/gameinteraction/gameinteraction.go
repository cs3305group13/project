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

// takes specified amount from the player, updates their money_in_pot entry and also adds it to the game pot
func TryTakeMoneyFromPlayer(DB *mysql_db.DB, playersTableName, pokerTablesTableName, tableID, playerName, bid string) (taken bool) {
	playersFunds := gameinfo.GetPlayersFunds(DB, playersTableName, playerName)
	
	playersBid, err := strconv.ParseFloat(bid, 64)
	utils.CheckError(err)

	if playersFunds < playersBid {
		taken = false
		return taken
	}

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET funds = funds - %v,
						      money_in_pot = money_in_pot + %v
						  WHERE username = "%s";`, playersTableName, playersBid, playersBid, playerName)
	
	_, err = db.Exec(query)
	utils.CheckError(err)

	query = fmt.Sprintf(`UPDATE %s
	                     SET money_in_pot = money_in_pot + %v
						 WHERE table_id = %s;`, pokerTablesTableName, bid, tableID)

	_, err = db.Exec(query)
	utils.CheckError(err)

	taken = true
	return taken
}

// checks if this player is the player set as 'current_player_making_move' in 'tables' table
func PlayersTurn(DB *mysql_db.DB, tablesTableName, playersTableName, tableID, username string) bool {

	currentPlayerMakingMove, _ := gameinfo.GetCurrentPlayerMakingMove(DB,tablesTableName, playersTableName, tableID)

	if username == currentPlayerMakingMove {
		return true
	}

	return false
}

// sets players state to 'FOLDED' and checks what effect this has on the game
func PlayerFolded(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber string, nextPlayerFoundBool bool) {
	
	numberOfPlayersAllIn := gameinfo.GetNumberOfPlayersAllIn(DB, playersTableName, tableID)

	numberOfPlayersStillPlaying := gameinfo.GetNumberOfPlayersStillPlaying(DB, playersTableName, tableID, username, seatNumber)
	// ^ contains number of players still in game (this player who wants to fold, players still playing, and all in players)
	
	if numberOfPlayersAllIn == 0 && numberOfPlayersStillPlaying == 2 {
		nextAvailablePlayer := gameflow.NextAvailablePlayer(DB, playersTableName, tableID, username, seatNumber)
		gameshowdown.SetWinner(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, nextAvailablePlayer)
	}
	if ! nextPlayerFoundBool && numberOfPlayersStillPlaying == 1{
		gameshowdown.SetWinner(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username)
		return

	} else if ! nextPlayerFoundBool && numberOfPlayersStillPlaying > 1{
		gameshowdown.ShowDown(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)

	} else {
		// if here then there are still other players playing therefore this player can fold

		db := mysql_db.EstablishConnection(DB)
		defer db.Close()

		query := fmt.Sprintf(`
							UPDATE %s
							SET player_state = "FOLDED"
							WHERE username = "%s"`, playersTableName, username)
							
		response, err := db.Exec(query)
		utils.CheckError(err)

		numberOfRows := mysql_db.GetNumberOfRowsAffected(response)

		if numberOfRows != 1 {
			panic("One and only one row should have been changed with this operation.")
		}
	}
}

// sets player's state relative to the amount they are willing to give the pot. this also
// redirects to PlayerFolded() if the amount is less than the highest_bid for this round
func PlayerTakesAction(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, amount string) (action string) {
	amountAsFloat64 := utils.ConvertToFloat(amount)
	playersMoneyCurrentlyInPot := gameinfo.GetPlayersMoneyInPot(DB, playersTableName, username)

	playersMoneyInPot := amountAsFloat64 + playersMoneyCurrentlyInPot

	_, highestBid := gameinfo.GetHighestBidder(DB, pokerTablesTableName, tableID)
	
	playersFunds := gameinfo.GetPlayersFunds(DB, playersTableName, username)

	setOperation := "current_player_making_move = "
	successful := gameflow.SetNextAvailablePlayerAfterThisOne(DB, tablesTableName, playersTableName, tableID, username, seatNumber, setOperation)

	numberOfPlayersStillPlaying := gameinfo.GetNumberOfPlayersStillPlaying(DB, playersTableName, tableID, username, seatNumber)
	// ^ contains number of players still in game (this player who wants to fold, players still playing, and all in players)

	if ! successful && numberOfPlayersStillPlaying == 1{
		gameshowdown.SetWinner(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username)
		return

	} else if ! successful && numberOfPlayersStillPlaying > 1{
		gameshowdown.ShowDown(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)
		return
	}

	if amountAsFloat64 >= playersFunds {
		// reset amount to be taken to match all of players funds instead of
		// what the user specified.
		amount = fmt.Sprintf("%f", playersFunds)

		playerAllIn(DB, playersTableName, username)
		action = "ALL_IN"

	} else if highestBid == 0.0 && playersMoneyInPot == 0.0 {
		playerChecked(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username)
		action = "CHECKED"

		return action  // can return because no money will be taken.

	} else if playersMoneyInPot < highestBid {
		PlayerFolded(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, successful)
		action = "FOLDED"

		return action
	
	} else if playersMoneyInPot < highestBid*2 {
		// if here raise amount was not at least double of previous highest bid

		// set amount to be amount needed for user to match highestBid
		amount = fmt.Sprintf(`%f`, highestBid - playersMoneyCurrentlyInPot)  

		playerCalled(DB, playersTableName, username)
		action = "CALLED"

	} else if playersMoneyInPot >= highestBid*2 {
		playerRaised(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, amount)
		action = "RAISED"

	}

	
	successfullyTaken := TryTakeMoneyFromPlayer(DB, playersTableName, pokerTablesTableName, tableID, username, amount)

	if ! successfullyTaken {
		panic("Amount was not taken properly.")
	}

	return action
}


func playerAllIn(DB *mysql_db.DB, playersTableName, username string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "ALL_IN"
						  WHERE username = "%s"`, playersTableName, username)

	_, err := db.Exec(query)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

}

func playerRaised( DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, raiseAmount string ) bool {

	// (step 1.) make player the highest bidder + update highest bid
	setOperation := fmt.Sprintf(`highest_bidder = "%s",
	                             highest_bid = "%s"`, username, raiseAmount)

	gameflow.AssignThisPlayerToRole(DB, pokerTablesTableName, tableID, username, setOperation)

	// (step 2.) set players state as "RAISED"
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "RAISED"
						  WHERE username = "%s";`, playersTableName, username)

	_, err := db.Exec(query)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return true
}

func playerCalled( DB *mysql_db.DB, playersTableName, username string ) bool {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "CALLED"
						  WHERE username = "%s";`, playersTableName, username)

	_, err := db.Exec(query)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return true
}

func playerChecked( DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username string ) bool {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "CHECKED"
						  WHERE username = "%s";`, playersTableName, username)

	_, err := db.Exec(query)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return true
}