package gamecontent

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils"
)

type gameDetails struct {
	Players *[]player
	TableDetails *tableDetails
}

func JSONGameDetails(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username string) []byte {

	tableDetails := getTableDetails(DB, tablesTableName, pokerTablesTableName, tableID)

	players := getPlayerDetails(DB, playersTableName, tableID, username)
	
	details := gameDetails{players, tableDetails}

	d, _ := json.MarshalIndent(details, "", " ")

	return d
}


type player struct {
	Username string
	Funds string
	SeatNumber string
	PlayerState string
	MoneyInPot string
	Cards string
}

func getPlayerDetails( DB *mysql_db.DB, playerTableName, tableID, username string ) *[]player {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	var query = fmt.Sprintf(`SELECT username, funds, seat_number, player_state, money_in_pot
	                                FROM %s 
	                                WHERE table_id = %s
									ORDER BY seat_number ASC;`, playerTableName, tableID)

	rows, err := db.Query(query)
	utils.CheckError(err)

	defer rows.Close()

	var players []player

	MAX_NUMBER_OF_PLAYERS := 8
	for i:=1; i<=MAX_NUMBER_OF_PLAYERS; i++ {
		players = append(players, player{})
	}

    for rows.Next() {
		p := player{}
		err := rows.Scan(&p.Username, &p.Funds, &p.SeatNumber, &p.PlayerState, &p.MoneyInPot)
		utils.CheckError(err)

		thisPlayersSeatNumber, err := strconv.Atoi(p.SeatNumber)
		utils.CheckError(err)

		if p.Username == username {
			p.Cards = getThisUsersCards(DB, playerTableName, username)
		}

		players[ thisPlayersSeatNumber - 1 ] = p
	}

	return &players
}

func getThisUsersCards(DB *mysql_db.DB, playerTableName, username string) (playerCards string) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()
	
	query :=fmt.Sprintf(`SELECT player_cards
	                     FROM %s
						 WHERE username = "%s";`, playerTableName, username)

	err := db.QueryRow(query).Scan(&playerCards)
	utils.CheckError(err)


	return playerCards
}

type tableDetails struct {
	CurrentPlayerMakingMove string
	GameState string
	CommunityCards string
	MoneyInPot string
}

func getTableDetails( DB *mysql_db.DB, tablesTableName, pokerTablesTableName, tableID string ) *tableDetails {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()
	
	var query = fmt.Sprintf(string(`SELECT current_player_making_move, game_in_progress
	                                FROM %s 
				                    WHERE table_id = %s;`), tablesTableName, tableID)
	
	tableDetails := tableDetails{}
	err := db.QueryRow(query).Scan(&tableDetails.CurrentPlayerMakingMove,
	                               &tableDetails.GameState)
	utils.CheckError(err)

	query = fmt.Sprintf(`SELECT community_cards, money_in_pot
	                     FROM %s
						 WHERE table_id = %s`, pokerTablesTableName, tableID)

	err = db.QueryRow(query).Scan(&tableDetails.CommunityCards, &tableDetails.MoneyInPot)
	
	utils.CheckError(err)

    return &tableDetails
}
