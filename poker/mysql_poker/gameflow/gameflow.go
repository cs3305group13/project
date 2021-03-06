package gameflow

import (
	"fmt"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/utils"
)

// function refreshes users time_since_request and checks/removes any players are idle
func UpdateUsersTimeSinceRequest(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, username, tableID, seatNumber string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET time_since_request = CURRENT_TIMESTAMP()
						  WHERE username = "%s";`, playersTableName, username)
	db.Exec(query)
}

// Method used to update next player who holds the responsibility.
// 
// setOperation := "highest_bidder = "
func SetNextAvailablePlayerAfterThisOne(DB *mysql_db.DB, tableName, playersTableName, tableID, username, seatNumber, setOperation string) (successful bool) {
	playerName := NextAvailablePlayer(DB, playersTableName, tableID, username, seatNumber)
	
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	setOperation += fmt.Sprintf(`"%s"`, playerName)

	query := fmt.Sprintf(`UPDATE %s
	                      SET %s
						  WHERE table_id = %s;`, tableName, setOperation, tableID)


	res, err := db.Exec(query)
	utils.CheckError(err)

	numOfRowsAffected := mysql_db.GetNumberOfRowsAffected(res)

	if numOfRowsAffected == 0 {
		// if here player was reassigned his role therefore no change was encountered.
		return false
		
	} else {
	    return true
	}
}

// function used to assign player as new current_player_making_move, dealer or highest_bidder in either poker tables or tables.
func AssignThisPlayerToRole(DB *mysql_db.DB, tableName, tableID, username, setOperation string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET %s
						  WHERE table_id = %s;`, tableName, setOperation, tableID)

	res, err := db.Exec(query)

	utils.CheckError(err)

	rowsAffected := mysql_db.GetNumberOfRowsAffected(res)
	if rowsAffected > 1 {
		panic("A change should have been caused unless method is used for wrong intention")
	}
}

// return next available players who are not idle nor in 'NOT_READY', 'LEFT', 'FOLDED', and 'ALL_IN' state.
func NextAvailablePlayers(DB *mysql_db.DB, playersTableName, tableID, username, seatNumber string) (playerNames []string) {

	// TODO: Make "5" seconds into an env variable.
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()
	
	// time_since_request > DATE_SUB(NOW(), INTERVAL %s SECOND) is used here to prevent selecting idle players
	query := fmt.Sprintf(`SELECT username
						  FROM %s
						  WHERE table_id = %s AND
						        seat_number > %s AND
								player_state != "NOT_READY" AND
								player_state != "LEFT" AND
								player_state != "FOLDED" AND
								player_state != "ALL_IN" AND
								time_since_request > DATE_SUB(NOW(), INTERVAL %s SECOND)
						  ORDER BY seat_number ASC;`, playersTableName, tableID, seatNumber, "5")
	playersAfter, err := db.Query(query)
	utils.CheckError(err)

	var playerName string

	for playersAfter.Next() {
		err := playersAfter.Scan(&playerName)
		utils.CheckError(err)
		playerNames = append(playerNames, playerName)
	}
	playersAfter.Close()

	query = fmt.Sprintf(`SELECT username
						 FROM %s
						 WHERE table_id = %s AND
						       seat_number <= %s AND
						       player_state != "NOT_READY" AND
							   player_state != "LEFT" AND
							   player_state != "FOLDED" AND
							   player_state != "ALL_IN" AND
							   time_since_request > DATE_SUB(NOW(), INTERVAL %s SECOND)
						 ORDER BY seat_number ASC;`, playersTableName, tableID, seatNumber, "5")  // seat_number used to be `<` it is `<=` because of dealer, big, small blinds overlap.
	
	playersBefore, err := db.Query(query)
	utils.CheckError(err)

	for playersBefore.Next() {
		err := playersBefore.Scan(&playerName)
		utils.CheckError(err)
		playerNames = append(playerNames, playerName)
	}
	playersBefore.Close()

	return playerNames
}

// return next available player who is not idle nor in 'NOT_READY', 'LEFT', 'FOLDED', and 'ALL_IN' state.
func NextAvailablePlayer(DB *mysql_db.DB, playersTableName, tableID, username, seatNumber string) (playerName string) {
	
	playerNames := NextAvailablePlayers(DB, playersTableName, tableID, username, seatNumber)

	if len(playerNames) != 0 {
		playerName = playerNames[0]
	}

	return playerName
}

// function clears users money in pot if they matched the highest bidder(highestBidder may have checked.)
func ClearUsersMoneyInPot(DB *mysql_db.DB, playersTableName, pokerTablesTableName, tableID string) {
	_, highestBid := gameinfo.GetHighestBidder(DB, pokerTablesTableName, tableID)

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s
	                      SET money_in_pot = 0
						  WHERE table_id = %s AND money_in_pot = %v;`, playersTableName, tableID, highestBid)

	_, err := db.Exec(query)

	utils.CheckError(err)

    query = fmt.Sprintf(`UPDATE %s
	                     SET highest_bid = 0
						 WHERE table_id = %s`, pokerTablesTableName, tableID)

	_, err = db.Exec(query)

	utils.CheckError(err)
}