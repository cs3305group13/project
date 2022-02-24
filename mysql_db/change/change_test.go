package change

import (
	"fmt"
	"testing"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/mysql_db"
)

var envs = env.GetEnvironmentVariables("../../testing.env")
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTablesTableName = envs["TESTING_POKER_TABLES_TABLE"]

func TestCreateRowFromRowOlderThan(t *testing.T) {
	
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	username := "johnny"
	deck := cards.NewDeck(1)
	deckString := cards.DeckString(deck)
	emptyDeckString := ""
	gameInProgress := false

	setOperation := fmt.Sprintf(`time_since_last_move = CURRENT_TIMESTAMP,
	                             current_player_making_move = "%s",
								 deck = "%s",
								 cards_not_in_deck = "%s",
								 game_in_progress = %t`, username, deckString, emptyDeckString, gameInProgress)
	
    hours := "24"
	rowID, changed := CreateRowFromRowOlderThan(tx, testingTablesTableName, "table_id", "time_since_last_move", hours, setOperation)

	if ! changed {
		t.Errorf("Table entry older than %s hours couldn't be found", hours)
	}
	fmt.Println(rowID)
}

func TestChangeValueInRow(t *testing.T) {
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	setOperation := "player_state = 'PLAYING'"

	ChangeValueInRows(tx, testingPlayersTableName, "table_id", tableID, setOperation)
}