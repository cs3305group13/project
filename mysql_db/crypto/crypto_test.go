package crypto

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/testing/mysql_poker"
	"github.com/cs3305/group13_2022/project/utils/env"
)

var envs = env.GetEnvironmentVariables("../../testing.env")
var DB = mysql_db.NewDB(envs)

var testingUserCredentialsTableName = envs["TESTING_USER_CREDENTIALS_TABLE"]

func TestAddUser(t *testing.T) {

	mysql_poker.RefreshUsersCredentails( DB)

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)

	defer tx.Rollback()
	defer db.Close()

    AddUser(tx, testingUserCredentialsTableName, "john", "password")
}