package gamecontent

import (
	"fmt"
	"testing"

	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/mysql_db"
)

func TestJSONGameDetails(t *testing.T) {
	envs := env.GetEnvironmentVariables("../../../testing.env")

	DB := mysql_db.NewDB(envs)
	playersTableName := envs["TESTING_PLAYERS_TABLE"]
	tablesTableName := envs["TESTING_TABLES_TABLE"]
	pokerTablesTableName := envs["TESTING_POKER_TABLES_TABLE"]

	tableID := "1"

	details := JSONGameDetails(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)

	fmt.Println(string(details))
}