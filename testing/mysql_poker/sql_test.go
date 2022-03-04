package mysql_poker

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils/env"
)


var envs = env.GetEnvironmentVariables("../../testing.env")

var DB = mysql_db.NewDB(envs)

func TestRefreshTablesTable(t *testing.T) {
	RefreshTablesTable(DB)
}

func TestRefreshPlayersTable(t *testing.T) {
	RefreshPlayerTable(DB)
}

func TestRefreshPokerTablesTable(t *testing.T) {
	RefreshPokerTable(DB)
}