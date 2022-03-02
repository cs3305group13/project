package gamestart

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinteraction"
	"github.com/cs3305/group13_2022/project/utils"
	"github.com/cs3305/group13_2022/project/utils/token"
)


func TryReadyUpPlayer(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	
    utils.CheckError(err)

    defer tx.Rollback()   // Defer a rollback in case anything fails.

	if ! gameinteraction.GameInProgress(w, r, tx, tablesTableName) { // aka not everyone is ready
		readyUpPlayer(w, r, DB, tx, tablesTableName, playersTableName, pokerTablesTableName)
	} else {
		w.Write([]byte("MESSAGE:\nGame is in progress."))
	}
}


func readyUpPlayer(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName string) {
	username := token.GetUsername(r, "token")

	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "READY"
						  WHERE username = "%s";`, playersTableName, username)
	_, err := tx.Exec(query)

	utils.CheckError(err)

	checkIfGameShouldStart(w, r, DB, tx, tablesTableName, playersTableName, pokerTablesTableName)
	
	err = tx.Commit()

	fmt.Println(err) // TODO: sql Deadlock Bug here encountered when I ran this the first time, stack trace pointed to mysql_poker/gamecontent.go most likely because startGame transaction was in progress.
}
func checkIfGameShouldStart(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName string) {
	
	tableID := token.GetTableID(r, "token")
	query := fmt.Sprintf(`SELECT COUNT(*)
	                      FROM %s
						  WHERE player_state = "READY" AND table_id = "%s"`, playersTableName, tableID)
	
	var numberOfReadyPlayersAtTable int
    err := tx.QueryRow(query).Scan(&numberOfReadyPlayersAtTable)
	utils.CheckError(err)
	
	query = fmt.Sprintf(`SELECT COUNT(*)
	                     FROM %s
	                     WHERE table_id = "%s";`, playersTableName, tableID)

	var totalNumberOfPlayersAtTable int
	err = tx.QueryRow(query).Scan(&totalNumberOfPlayersAtTable)
	utils.CheckError(err)

	if totalNumberOfPlayersAtTable > 1 && numberOfReadyPlayersAtTable == totalNumberOfPlayersAtTable {
		startGame(r, tx, tablesTableName, playersTableName, pokerTablesTableName)
		beginGame(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)
	} else if totalNumberOfPlayersAtTable == 1 {
		w.Write([]byte("PROBLEM:\nTo start atleast two players need to be present."))
	} else {
		w.Write([]byte("MESSAGE:\nSome players are not ready yet."))
	}
}
func startGame(r *http.Request, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName string) {

	tableID := token.GetTableID(r, "token")

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
	                     SET player_state = "PLAYING"
						 WHERE player_state = "READY" AND table_id = %s;`, playersTableName, tableID)

	_, err = tx.Exec(query)
	utils.CheckError(err)
}
func beginGame(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {
	currentPlayerMakingMove, seatNumber := gameinfo.GetCurrentPlayerMakingMove(DB, tablesTableName, playersTableName, tableID)
	playersAfter, playersBefore := gameflow.NextAvailablePlayers(DB, playersTableName, tableID, currentPlayerMakingMove, seatNumber)

	var playerName string
	var playerState string

	var smallBigCurrentPlayers []string

	// the 2 for loops 
	for playersAfter.Next() && len(smallBigCurrentPlayers) != 3 {
		err := playersAfter.Scan( &playerName, &playerState )
		utils.CheckError(err)

		smallBigCurrentPlayers = append(smallBigCurrentPlayers, playerName)
	}
	if len(smallBigCurrentPlayers) != 3 {
		for playersBefore.Next() {
			err := playersAfter.Scan( &playerName, &playerState )
		    utils.CheckError(err)

		    smallBigCurrentPlayers = append(smallBigCurrentPlayers, playerName)
		}
	}
	
	smallBlind := smallBigCurrentPlayers[0]
	bigBlind := smallBigCurrentPlayers[1]
	newCurrentPlayerMakingMove := smallBigCurrentPlayers[2]

	smallBlindAmount := "1.0"
	bigBlindAmount := "2.0"
	gameinteraction.TakeMoneyFromPlayer(DB, playersTableName, pokerTablesTableName, tableID, smallBlind, smallBlindAmount)
	gameinteraction.TakeMoneyFromPlayer(DB, playersTableName, pokerTablesTableName, tableID, bigBlind, bigBlindAmount)

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	// set the big blind as the highest bidder
	query := fmt.Sprintf(`UPDATE %s
						  SET highest_bidder = "%s"
						  WHERE table_id = %s;`, pokerTablesTableName, bigBlind, tableID)

	res, err := db.Exec(query)
	utils.CheckError(err)
	if utils.GetNumberOfRowsAffected(res) != 1 {
		panic("Exactly one row should have been affected here.")
	}

	// set the current player as the player after the big blind
    query = fmt.Sprintf(`UPDATE %s
	                     SET current_player_making_move = "%s"
						 WHERE table_id = %s;`, tablesTableName, newCurrentPlayerMakingMove, tableID)

	res, err = db.Exec(query)
	utils.CheckError(err)
	if utils.GetNumberOfRowsAffected(res) != 1 {
		panic("Exactly one row should have been affected here.")
	}
}