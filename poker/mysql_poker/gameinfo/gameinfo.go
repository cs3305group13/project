package gameinfo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils"
)

func GetCurrentPlayerMakingMove(DB *mysql_db.DB, tablesTableName, playersTableName, tableID string) (currentPlayerMakingMove, seatNumber string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT current_player_making_move
	                      FROM %s
						  WHERE table_id = %s;`, tablesTableName, tableID)

	err := db.QueryRow(query).Scan(&currentPlayerMakingMove)
	utils.CheckError(err)

	query = fmt.Sprintf(`SELECT seat_number
	                     FROM %s
						 WHERE table_id = %s AND username = "%s";`, playersTableName, tableID, currentPlayerMakingMove)
	err = db.QueryRow(query).Scan(&seatNumber)
	utils.CheckError(err)


	return currentPlayerMakingMove, seatNumber
}

func GetDealerAndHighestBidder(DB *mysql_db.DB, playersTableName, pokerTablesTableName, tableID string) (highestBidder, highestBidderSeatNumber, dealer, dealerSeatNumber string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT highest_bidder, dealer
	                      FROM %s
						  WHERE table_id = %s`, pokerTablesTableName, tableID)
	
	err := db.QueryRow(query).Scan(&highestBidder, &dealer)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}
						  
	query = fmt.Sprintf(`SELECT seat_number
	                     FROM %s
						 WHERE username = "%s";`, playersTableName, highestBidder)
	
	err = db.QueryRow(query).Scan(&highestBidderSeatNumber)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	query = fmt.Sprintf(`SELECT seat_number
	                     FROM %s
						 WHERE username = "%s";`, playersTableName, dealer)

	err = db.QueryRow(query).Scan(&dealerSeatNumber)
	if err != sql.ErrNoRows {
		utils.CheckError(err)
	}

	return highestBidder, highestBidderSeatNumber, dealer, dealerSeatNumber
}

func GetHighestBidder(DB *mysql_db.DB, pokerTablesTableName, tableID string) (bidder string, bid float64) {
	
	query := fmt.Sprintf(`SELECT highest_bidder, highest_bid
	                      FROM %s
						  WHERE table_id = %s`, pokerTablesTableName, tableID)

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	err := db.QueryRow(query).Scan(&bidder, &bid)

	utils.CheckError(err)

	return bidder, bid
}

// Gets total number of players regardless of playerState
func GetNumberOfPlayersAtTable( DB *mysql_db.DB, playersTableName, tableID string ) (numOfRows int) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	var query = fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE table_id = %s;`, playersTableName, tableID)

	err := db.QueryRow(query).Scan(&numOfRows)
	
	if err != nil {
		log.Fatal(err)
	}

	return
}
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	var query = fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE table_id = %s;`, playersTableName, tableCode)

	err := db.QueryRow(query).Scan(&numOfRows)
	
	if err != nil {
		log.Fatal(err)
	}

	return
}

func GetNumberOfPlayersAllIn(DB *mysql_db.DB, playersTableName, tableID string) int {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT COUNT(*)
	                      FROM %s
						  WHERE table_id = %s AND player_state = "ALL_IN"`, playersTableName, tableID)

	var numberOfAllInPlayers int
	err := db.QueryRow(query).Scan(&numberOfAllInPlayers)

	if err != sql.ErrNoRows {
	    utils.CheckError(err)
	}

	return numberOfAllInPlayers
}

func GetNextAvailableSeat(DB *mysql_db.DB, playersTableName, tableID string) (nextAvailableSeat string, seatFound bool) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT seat_number
	                      FROM %s
						  WHERE table_id = %s
						  ORDER BY seat_number ASC;`, playersTableName, tableID)
	rows, err := db.Query(query)

	utils.CheckError(err)

	var takenSeats []string
	availableSeats := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	var seatNumber int
	for rows.Next() {
		err = rows.Scan(&seatNumber)
		seatNumberIndex := seatNumber - 1

		takenSeats = append(takenSeats, availableSeats[seatNumberIndex])
	}
	if len(takenSeats) == 8 {
		return "", false
	} else {
	    for i:=0; i<8; i++ {
			if ! utils.ArrayContains(takenSeats, availableSeats[i]) {
				return availableSeats[i], true
			}
		}
		panic("This shouldn't have happened")
	}
}

func GetPlayersFunds(DB *mysql_db.DB, playersTableName, username string) (funds float64) {
	
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()
	
	query := fmt.Sprintf(`SELECT funds
	                      FROM %s
						  WHERE username = "%s"`, playersTableName, username)

	err := db.QueryRow(query).Scan(&funds)
	utils.CheckError(err)

	return funds
}