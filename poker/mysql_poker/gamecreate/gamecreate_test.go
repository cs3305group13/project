package gamecreate

import (
	"fmt"
	"testing"

	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/mysql_db"
)

var envs = env.GetEnvironmentVariables("../../../testing.env")
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTablesTableName = envs["TESTING_POKER_TABLES_TABLE"]


func TestInsertNewPokerTable(t *testing.T) {
	
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
	
	var tableID = "2"
	var username = "barry"

	inserted := insertNewPokerTable(tx, testingPokerTablesTableName, tableID, username)

	if inserted == false {
		t.Error("poker table wasn't inserted")
	}
}

func TestInsertNewTable(t *testing.T) {
	
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
	
	var username = "barry"

	tableID, inserted := insertNewTable(tx, testingTablesTableName, username)
	if ! inserted {
		t.Error("Entry wasn't inserted")
	}
	fmt.Println(tableID)
}

func TestChangeOldPokerTable(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	var tableID = "2"
	differentUsername := "mary"

	changed := changeOldPokerTable(tx, testingPokerTablesTableName, tableID, differentUsername)
	if ! changed {
		t.Error("poker table was not changed")
	}
}

func TestChangeOldTable(t *testing.T) {
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	differentUsername := "bob"

	// For a change to occur here a table most be older than 24hrs
	tableID, _ := changeOldTable(tx, testingTablesTableName, differentUsername)
	
	fmt.Println(tableID)
}

func TestAssignNewTable(t *testing.T) {
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	differentUsername := "lucifer"

	AssignNewTable(tx, testingTablesTableName, testingPokerTablesTableName, differentUsername)
}