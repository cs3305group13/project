package poker

import (
	"html"
	"net/http"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/utils/token"

	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamecards"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinteraction"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameshowdown"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamestart"
)



func HandleUserAjaxRequest( w http.ResponseWriter, r *http.Request ) {
	valid := token.TokenValid( w, r, true )
	if valid {
		envs := env.GetEnvironmentVariables("../../../production.env")
		handleUserButtons(w, r, envs)
	}
}

func handleUserButtons( w http.ResponseWriter, r *http.Request, envs map[string]string ) {

	r.ParseForm()
	if len(r.Form) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DB := mysql_db.NewDB(envs)
	tablesTableName := envs["TABLES_TABLE"]
	playersTableName := envs["PLAYERS_TABLE"]
	pokerTablesTableName := envs["POKER_TABLES_TABLE"]

	readyButton := html.EscapeString(r.Form.Get("ready_button"))
	if readyButton != "" {
        gamestart.TryReadyUpPlayer(w, r, DB, tablesTableName, playersTableName, pokerTablesTableName)

		return
	}
	tableID := token.GetTableID(r, "token")
	inProgress := gameinfo.GameInProgress( DB, tablesTableName, tableID )
	if inProgress {
		username := token.GetUsername(r, "token")

		// checks if players turn.
		playersTurn := gameinteraction.PlayersTurn(DB, tablesTableName, playersTableName, tableID, username)
		if ! playersTurn {
			w.Write([]byte("MESSAGE:\nSorry not your turn"))
			return
		} else {

			seatNumber := token.GetSeatNumber(r, "token")

			action := html.EscapeString(r.Form.Get("action"))
			if action == "Fold" {
				FoldPlayer(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber)	
			} 

			amount := html.EscapeString(r.Form.Get("amount"))
			if ( action == "Raise" || action == "Call" || action == "Check" ) && amount != "" {
				RaiseCallCheckPlayer(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, amount)

				highestBidder, _ := gameinfo.GetHighestBidder(DB, pokerTablesTableName, tableID)
				nextAvailablePlayer := gameflow.NextAvailablePlayer(DB, playersTableName, tableID, username, seatNumber)
				if nextAvailablePlayer == highestBidder {
					EndRound(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)
				}
			}
		}
	}
}

// Creates a transaction which handles folding player.
func FoldPlayer(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber string) {

	setOperation := "current_player_making_move = "
	successful := gameflow.SetNextAvailablePlayerAfterThisOne(DB, tablesTableName, playersTableName, tableID, username, seatNumber, setOperation)

	if ! successful {
		// another non all in player still playing couldn't be found.
		gameEndedEarly := true
		gamecards.AddToCommunityCards(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, gameEndedEarly)

	}
	
	gameinteraction.PlayerFolded(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, successful)

}

// Creates a transaction which handles ( raising || calling || checking ) player.
func RaiseCallCheckPlayer(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, amount string) {

	gameinteraction.PlayerTakesAction( DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, amount )

}

// Creates a transaction which handles adding to community cards.
func EndRound(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {

	gameEndedEarly := false

	successfullyAdded := gamecards.AddToCommunityCards(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, gameEndedEarly)
	if ! successfullyAdded {
		// if a card wasn't added then that means there are already 5 cards present.
		gameshowdown.ShowDown(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)

	} else {
		// after each round clear every players money_in_pot field who managed to match the highestBid for this round
	    gameflow.ClearUsersMoneyInPot(DB, playersTableName, pokerTablesTableName, tableID)
	}

}