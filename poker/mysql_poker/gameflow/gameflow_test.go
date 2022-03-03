package gameflow

import (
	"fmt"
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/testing/mysql_poker"
	"github.com/cs3305/group13_2022/project/utils/env"
)


var envs = env.GetEnvironmentVariables("../../../testing.env")
	
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTableName = envs["TESTING_POKER_TABLES_TABLE"]

func TestNextAvailablePlayer(t *testing.T) {

	mysql_poker.RefreshPlayers(DB, testingPlayersTableName)

	tableID := "1"
	
	NextAvailablePlayer(DB, testingPlayersTableName, tableID, "derek", "1")
}


func TestNextAvailablePlayers(t *testing.T) {

	mysql_poker.RefreshPlayers(DB, testingPlayersTableName)

	tableID := "1"
	playerNames := NextAvailablePlayers(DB, testingPlayersTableName, tableID, "derek", "1")

	// No output because 5 seconds in where clause
	for i:=0; i<len(playerNames); i++ {
		fmt.Println(playerNames[i])
	}
}

func TestAssignThisPlayerToRole(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	username := "derek"
	setOperation := fmt.Sprintf("current_player_making_move = '%s'", username)

	AssignThisPlayerToRole(tx, testingTablesTableName, tableID, username, setOperation)
}

func TestSetNextAvailablePlayerAfterThisOne(t *testing.T) {
	mysql_poker.RefreshPlayers(DB, testingPlayersTableName)

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	
	SetNextAvailablePlayerAfterThisOne(DB, tx, testingPokerTableName, testingPlayersTableName, tableID, "derek", "1", "dealer = ")
}