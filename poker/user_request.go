package poker

import (
	"html"
	"net/http"

	"github.com/cs3305/group13_2022/project/utils"
	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils/token"

	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinteraction"
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
	pokerTablesTableName := envs["PLAYERS_TABLE"]

	readyButton := html.EscapeString(r.Form.Get("ready_button"))
	if readyButton != "" {
		gamestart.TryReadyUpPlayer(w, r, DB, tablesTableName, playersTableName, pokerTablesTableName)
        return
	}

	// checks if players turn.
	if gameinteraction.PlayersTurn(w, r, DB, tablesTableName) == false {
		return
	} else {

		foldButton := html.EscapeString(r.Form.Get("fold_button"))
		if foldButton != "" {
			gameinteraction.PlayerFolded(w, r, DB, tablesTableName, playersTableName, pokerTablesTableName)
		}
		raiseButton := html.EscapeString(r.Form.Get("raise_button"))
		raiseAmount := html.EscapeString(r.Form.Get("raise_amount"))
		if raiseButton != "" && raiseAmount != "" {
			gameinteraction.PlayerRaised(w, r, DB, tablesTableName, playersTableName, raiseAmount )
		}

		// call_Button := r.Form.Get("call_button")
		// check_Button := r.Form.Get("check_button")
	}
}