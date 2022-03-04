package gameshowdown

import (
	"database/sql"
	"fmt"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/utils"

	"github.com/chehsunliu/poker"
)

func ShowDown(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {
	
	players := GetPlayersAndCards(DB, playersTableName, tableID)
	
	pokerCommunityCards := getEndOfGameCommunityCards(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID)
    
	
	for i:=0; i<len(players); i++ {
		var playersCards []poker.Card
		extractedPlayerCards := cards.ExtractDeck(players[i].cards)
		card1 := (*extractedPlayerCards)[0]
		card2 := (*extractedPlayerCards)[1]

		playersCards = append(playersCards, poker.NewCard( card1 ) )
		playersCards = append(playersCards, poker.NewCard( card2 ) )



		playerCardsAndCommunityCards := append(pokerCommunityCards, playersCards[0], playersCards[1])

		cardsScore := poker.Evaluate( playerCardsAndCommunityCards )

		players[i].score = int(cardsScore)
	}

	// decide winner
	var winner string
	bestScore := 10000  // the lower the number the better the hand
	for i:=0; i<len(players); i++ {
		if bestScore > players[i].score {
		    bestScore = players[i].score
			winner = players[i].username
		}

	}

	SetWinner(DB, tx, playersTableName, pokerTablesTableName, tableID, winner)
}

type player struct {
	username string
	cards string
	score int
}

func GetPlayersAndCards(DB *mysql_db.DB, playersTableName, tableID string) (players []player) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT username, player_cards
	                      FROM %s
						  WHERE table_id = %s AND 
						       player_state IN ("PLAYING", "ALL_IN", "RAISED", "CALLED", "CHECKED");`, playersTableName, tableID)

	
	rows, err := db.Query(query)
	utils.CheckError(err)
	
	var p player
	for rows.Next() {
		err := rows.Scan(&p.username, &p.cards)
		utils.CheckError(err)

		players = append(players, p)
	}

	return players
}

func getEndOfGameCommunityCards(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID string) []poker.Card {
	communityCards := gameinfo.GetCommunityCards(DB, pokerTablesTableName, tableID)


	// translate community cards to match chehsunliu implementation
	var pokerCommunityCards []poker.Card
	for i:=0; i<len(*communityCards); i++ {

		communityCard := (*communityCards)[i]
		pokerCommunityCards = append( pokerCommunityCards, poker.NewCard( communityCard ))
	}

	return pokerCommunityCards
}


func SetWinner(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username string) {

	query := fmt.Sprintf(`SELECT money_in_pot
	                      FROM %s
						  WHERE table_id = %s;`, pokerTablesTableName, tableID)

	var moneyInPot string
	err := tx.QueryRow(query).Scan(&moneyInPot)

	utils.CheckError(err)
	
	query = fmt.Sprintf(`UPDATE %s
	                     SET player_state = "WINNER",
						     funds = funds + %s
						 WHERE table_id = %s AND username = "%s";`, playersTableName, moneyInPot, tableID, username)

	_, err = tx.Exec(query)

	if err != sql.ErrNoRows {
	    utils.CheckError(err)
	}

}