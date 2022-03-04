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
	
	communityCards := gameinfo.GetCommunityCards(DB, pokerTablesTableName, tableID)

	// translate community cards to match chehsunliu implementation
	var pokerCommunityCards []poker.Card
	for i:=0; i<len(*communityCards); i++ {

		communityCard := (*communityCards)[i]
		pokerCommunityCards = append( pokerCommunityCards, poker.NewCard( communityCard ))
	}
    
	
	for i:=0; i<len(players); i++ {
		var playersCards []poker.Card
		extractedPlayerCards := cards.ExtractDeck(players[i].cards)
		card1 := (*extractedPlayerCards)[0]
		card2 := (*extractedPlayerCards)[1]

		playersCards = append(playersCards, poker.NewCard( card1 ) )
		playersCards = append(playersCards, poker.NewCard( card2 ) )



		playerCardsAndCommunityCards := append(pokerCommunityCards, playersCards[0], playersCards[1])

		fmt.Println(playerCardsAndCommunityCards)

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


		fmt.Println( poker.RankString(int32(players[i].score)) )
	}

	setWinner(DB, tx, playersTableName, pokerTablesTableName, tableID, winner)
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


func setWinner(DB *mysql_db.DB, tx *sql.Tx, playersTableName, pokerTablesTableName, tableID, username string) {
	
	query := fmt.Sprintf(`UPDATE %s
	                      SET player_state = "WINNER"
						  WHERE table_id = %s AND username = "%s";`, playersTableName, tableID, username)

	_, err := tx.Exec(query)

	if err != sql.ErrNoRows {
	    utils.CheckError(err)
	}

}