package gamecontent

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils"
)

type gameDetails struct {
	Players *[]player
	TableDetails *tableDetails
}

func JSONGameDetails(DB *mysql_db.DB, playersTableName, tablesTableName, pokerTablesTableName, tableID string) []byte {

	players := getPlayerDetails(DB, playersTableName, tableID)

	tableDetails := getTableDetails(DB, tablesTableName, pokerTablesTableName, tableID)
	
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

func getPlayerDetails( DB *mysql_db.DB, tableName, tableID string ) *[]player {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	var query = fmt.Sprintf(string(`SELECT username, funds, seat_number, player_state, money_in_pot, player_cards
	                                FROM %s 
	                                WHERE table_id = %s
									ORDER BY seat_number ASC;`), tableName, tableID)

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var players []player

	MAX_NUMBER_OF_PLAYERS := 8
	for i:=1; i<=MAX_NUMBER_OF_PLAYERS; i++ {
		players = append(players, player{})
	}

    for rows.Next() {
		p := player{}
		err := rows.Scan(&p.Username, &p.Funds, &p.SeatNumber, &p.PlayerState, &p.MoneyInPot, &p.Cards)
		utils.CheckError(err)

		thisPlayersSeatNumber, err := strconv.Atoi(p.SeatNumber)
		utils.CheckError(err)

		players[ thisPlayersSeatNumber - 1 ] = p
	}

	return &players
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
