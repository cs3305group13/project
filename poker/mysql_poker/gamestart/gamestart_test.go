package gamestart

import (
	"context"
	"testing"

	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/testing/utils"

	mysql_utils "github.com/cs3305/group13_2022/project/utils"
)



func TestReadyUpPlayer(t *testing.T) {
	envs := env.GetEnvironmentVariables("../../testing.env")
	
	w := utils.CreateRegularResponse()
	r := utils.CreateRequestWithPokerCookie()

	DB := mysql_db.NewDB(envs)

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	
    mysql_utils.CheckError(err)

    defer tx.Rollback()   // Defer a rollback in case anything fails.

	tablesTableName := envs["TESTING_TABLES_TABLE"]
	playersTableName := envs["TESTING_PLAYERS_TABLE"]
	pokerTablesTableName := envs["TESTING_POKER_TABLES_TABLE"]

	readyUpPlayer(w, r, DB, tx, tablesTableName, playersTableName, pokerTablesTableName)
}