package insert

import (
	"testing"

	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/mysql_db"
)

var envs = env.GetEnvironmentVariables("../../testing.env")
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPokerTablesTableName = envs["TESTING_POKER_TABLES_TABLE"]

func TestInsertTableEntry(t *testing.T) {
    db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	pokerColumnNames := "table_id, community_cards, highest_bidder, highest_bid, dealer"
	pokerValues := "2, '', 'joe', 0.0, 'joe'"

	InsertTableEntry(tx, testingPokerTablesTableName, pokerColumnNames, pokerValues)
}