package mysql_poker

import (
	"fmt"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils"
)

func RefreshPlayers(DB *mysql_db.DB, playersTableName string) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET time_since_request = CURRENT_TIMESTAMP();`, playersTableName)

	_, err := db.Exec(query)
	utils.CheckError(err)

}