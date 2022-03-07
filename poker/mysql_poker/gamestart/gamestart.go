package gamestart

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamecards"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinteraction"

	"github.com/cs3305/group13_2022/project/utils"
	"github.com/cs3305/group13_2022/project/utils/token"
)


func TryReadyUpPlayer(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName string) {

	tableID := token.GetTableID(r, "token")

	if gameinfo.GameInProgress(DB, tablesTableName, tableID) { // aka not everyone is ready
		w.Write([]byte("MESSAGE:\nGame is in progress."))
		return 

	} else {
		db := mysql_db.EstablishConnection(DB)
		tx := mysql_db.NewTransaction(db)
		defer tx.Rollback()   // Defer a rollback in case anything fails.
		defer db.Close()

		tableID := token.GetTableID(r, "token")
		username := token.GetUsername(r, "token")

		begin := readyUpPlayer(w, DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username)
		err := tx.Commit()
		utils.CheckError(err)

		if begin {
			beginGame(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)
		}
		
		return
	}
}


func readyUpPlayer(w http.ResponseWriter, DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username string) (beginGame bool) {
	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "READY"
						  WHERE username = "%s";`, playersTableName, username)
	_, err := tx.Exec(query)

	utils.CheckError(err)

	beginGame = checkIfGameShouldStart(w, DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID)
	return beginGame
}
func checkIfGameShouldStart(w http.ResponseWriter, DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID string) (beginGame bool) {

	query := fmt.Sprintf(`SELECT COUNT(*)
	                      FROM %s
						  WHERE player_state = "READY" AND table_id = "%s"`, playersTableName, tableID)
	
	var numberOfReadyPlayersAtTable int
    err := tx.QueryRow(query).Scan(&numberOfReadyPlayersAtTable)
	utils.CheckError(err)
	
	query = fmt.Sprintf(`SELECT COUNT(*)
	                     FROM %s
	                     WHERE table_id = %s;`, playersTableName, tableID)

	var totalNumberOfPlayersAtTable int
	err = tx.QueryRow(query).Scan(&totalNumberOfPlayersAtTable)
	utils.CheckError(err)

	if totalNumberOfPlayersAtTable > 1 && numberOfReadyPlayersAtTable == totalNumberOfPlayersAtTable {
		startGame(tx, tablesTableName, playersTableName, pokerTablesTableName, tableID)
		beginGame = true
		return beginGame
	} else if totalNumberOfPlayersAtTable == 1 {
		w.Write([]byte("PROBLEM:\nTo start at least two players need to be present."))
		beginGame = false
		return beginGame
	} else {
		w.Write([]byte("MESSAGE:\nSome players are not ready yet."))
		beginGame = false
		return beginGame
	}
}
func startGame(tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID string) (beginGame bool) {

	deck := cards.NewDeck(1)
	cards.Shuffle(deck)

	deckString := cards.DeckString(deck)
	cardsNotInDeckString := ""

	query := fmt.Sprintf(`UPDATE %s
	                      SET game_in_progress = True,
						      deck = "%s",
							  cards_not_in_deck = "%s"
						  WHERE table_id = %s;`, tablesTableName, deckString, cardsNotInDeckString, tableID)
	_, err := tx.Exec(query)
	utils.CheckError(err)

	query = fmt.Sprintf(`UPDATE %s
	                     SET player_state = "PLAYING",
						     money_in_pot = "0.0"
						 WHERE player_state = "READY" AND table_id = %s;`, playersTableName, tableID)

	_, err = tx.Exec(query)
	utils.CheckError(err)

	return true
}
func beginGame(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {
	currentPlayerMakingMove, seatNumber := gameinfo.GetCurrentPlayerMakingMove(DB, tablesTableName, playersTableName, tableID)
	
	smallBlind, bigBlind, newCurrentPlayerMakingMove := findWhoShouldBeSmallAndBigBlind(DB, playersTableName, tableID, currentPlayerMakingMove, seatNumber)

	smallBlindAmount := "1.0"
	bigBlindAmount := "2.0"

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	// initialize the big blind as the highest bidder and reset money_in_pot
	query := fmt.Sprintf(`UPDATE %s
						  SET community_cards = "",
						      highest_bidder = "%s",
						      highest_bid = "%s",
							  money_in_pot = "0.0"
						  WHERE table_id = %s;`, pokerTablesTableName, bigBlind, bigBlindAmount, tableID)

	_, err := db.Exec(query)  // result is ignored because TryTakeMoneyFromPlayers updates highestBidder
	utils.CheckError(err)

	_ = gameinteraction.TryTakeMoneyFromPlayer(DB, playersTableName, pokerTablesTableName, tableID, smallBlind, smallBlindAmount)
	_ = gameinteraction.TryTakeMoneyFromPlayer(DB, playersTableName, pokerTablesTableName, tableID, bigBlind, bigBlindAmount)

	// set the current player as the player after the big blind
    query = fmt.Sprintf(`UPDATE %s
	                     SET current_player_making_move = "%s"
						 WHERE table_id = %s;`, tablesTableName, newCurrentPlayerMakingMove, tableID)

	res, err := db.Exec(query)
	utils.CheckError(err)
	if mysql_db.GetNumberOfRowsAffected(res) > 1 {
		panic("Exactly one row should have been affected here.")
	}

	gamecards.GivePlayersTheirCards(DB, tablesTableName, playersTableName, tableID)
}


func findWhoShouldBeSmallAndBigBlind(DB *mysql_db.DB, playersTableName, tableID, currentPlayerMakingMove, theirSeatNumber string) (small, big, newCurrentPlayer string) {
	playerNames := gameflow.NextAvailablePlayers(DB, playersTableName, tableID, currentPlayerMakingMove, theirSeatNumber)

	small = playerNames[0]
	big = playerNames[1]
	if len(playerNames) == 2 {
	    newCurrentPlayer = playerNames[0]
	} else {
		newCurrentPlayer = playerNames[2]
	}

	return small, big, newCurrentPlayer
}