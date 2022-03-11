package find

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils/env"
)

var envs = env.GetEnvironmentVariables("../../testing.env")
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]

func TestFindRowByValue( t *testing.T ) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)

	defer tx.Rollback()
	defer db.Close()

	tableCode := "1"

	FindRowByValue(tx, testingTablesTableName, "table_id", tableCode, "table_id")
}
