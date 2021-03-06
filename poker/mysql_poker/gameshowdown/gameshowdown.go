package gameshowdown

import (
	"database/sql"
	"fmt"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamecards"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/utils"

	"github.com/chehsunliu/poker"
)

func ShowDown(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {
	
	players := gameinfo.GetPlayersAndCards(DB, playersTableName, tableID)
	
	pokerCommunityCards := getEndOfGameCommunityCards(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)
    
	
	for i:=0; i<len(players); i++ {
		var playersCards []poker.Card
		extractedPlayerCards := cards.ExtractDeck(players[i].Cards)
		card1 := (*extractedPlayerCards)[0]
		card2 := (*extractedPlayerCards)[1]

		playersCards = append(playersCards, poker.NewCard( card1 ) )
		playersCards = append(playersCards, poker.NewCard( card2 ) )



		playerCardsAndCommunityCards := append(pokerCommunityCards, playersCards[0], playersCards[1])

		cardsScore := poker.Evaluate( playerCardsAndCommunityCards )

		players[i].Score = int(cardsScore)
	}

	// decide winner
	var winner string
	bestScore := 10000  // the lower the number the better the hand
	for i:=0; i<len(players); i++ {
		if bestScore > players[i].Score {
		    bestScore = players[i].Score
			winner = players[i].Username
		}

	}

	SetWinner(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, winner)
}


func getEndOfGameCommunityCards(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) []poker.Card {
	communityCards := gameinfo.GetCommunityCards(DB, pokerTablesTableName, tableID)
	for len( *communityCards ) < 5 {

		gameEndedEarly := true
		gamecards.AddToCommunityCards(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, gameEndedEarly)

		communityCards = gameinfo.GetCommunityCards(DB, pokerTablesTableName, tableID)
	}

	// translate community cards to match chehsunliu implementation
	var pokerCommunityCards []poker.Card
	for i:=0; i<len(*communityCards); i++ {

		communityCard := (*communityCards)[i]
		pokerCommunityCards = append( pokerCommunityCards, poker.NewCard( communityCard ))
	}

	return pokerCommunityCards
}


func SetWinner(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT money_in_pot
	                      FROM %s
						  WHERE table_id = %s;`, pokerTablesTableName, tableID)

	var moneyInPot string
	err := db.QueryRow(query).Scan(&moneyInPot)

	utils.CheckError(err)
	
	query = fmt.Sprintf(`UPDATE %s
	                     SET player_state = "WINNER",
						     funds = funds + %s
						 WHERE table_id = %s AND username = "%s";`, playersTableName, moneyInPot, tableID, username)

	_, err = db.Exec(query)

	if err != sql.ErrNoRows {
	    utils.CheckError(err)
	}

	resetGameState(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)
}


// method called if game state is to be reset .ie game ended
func resetGameState(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	_, _, dealer, dealerSeatNumber := gameinfo.GetDealerAndHighestBidder(DB, playersTableName, pokerTablesTableName, tableID)

	setOperation := "current_player_making_move = "
	gameflow.SetNextAvailablePlayerAfterThisOne(DB, tablesTableName, playersTableName, tableID, dealer, dealerSeatNumber, setOperation)

	query := fmt.Sprintf(`UPDATE %s
	                      SET game_in_progress = false
						  WHERE table_id = %s`, tablesTableName, tableID)


	_, err := db.Exec(query)

	utils.CheckError(err)

}