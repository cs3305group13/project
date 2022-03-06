package gameshowdown

import (
	"database/sql"
	"fmt"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/utils"

	"github.com/chehsunliu/poker"
)

<<<<<<< HEAD
func ShowDown(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {
	
	players := gameinfo.GetPlayersAndCards(DB, playersTableName, tableID)
	
	pokerCommunityCards := getEndOfGameCommunityCards(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)
=======
func ShowDown(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {
	
	players := GetPlayersAndCards(DB, playersTableName, tableID)
	
	pokerCommunityCards := getEndOfGameCommunityCards(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
    
	
	for i:=0; i<len(players); i++ {
		var playersCards []poker.Card
<<<<<<< HEAD
		extractedPlayerCards := cards.ExtractDeck(players[i].Cards)
=======
		extractedPlayerCards := cards.ExtractDeck(players[i].cards)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
		card1 := (*extractedPlayerCards)[0]
		card2 := (*extractedPlayerCards)[1]

		playersCards = append(playersCards, poker.NewCard( card1 ) )
		playersCards = append(playersCards, poker.NewCard( card2 ) )



		playerCardsAndCommunityCards := append(pokerCommunityCards, playersCards[0], playersCards[1])

		cardsScore := poker.Evaluate( playerCardsAndCommunityCards )

<<<<<<< HEAD
		players[i].Score = int(cardsScore)
=======
		players[i].score = int(cardsScore)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
	}

	// decide winner
	var winner string
	bestScore := 10000  // the lower the number the better the hand
	for i:=0; i<len(players); i++ {
<<<<<<< HEAD
		if bestScore > players[i].Score {
		    bestScore = players[i].Score
			winner = players[i].Username
=======
		if bestScore > players[i].score {
		    bestScore = players[i].score
			winner = players[i].username
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
		}

	}

<<<<<<< HEAD
	SetWinner(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, winner)
}


func getEndOfGameCommunityCards(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) []poker.Card {
=======
	SetWinner(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, winner)
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
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
	communityCards := gameinfo.GetCommunityCards(DB, pokerTablesTableName, tableID)


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

	// reset the poker table pot
	query = fmt.Sprintf(`UPDATE %s
		SET money_in_pot = 0.0
		WHERE table_id = %s;`, pokerTablesTableName, tableID))

		_, err = db.Exec(query)

		if err != sql.ErrNoRows {
			utils.CheckError(err)
		}
	
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
<<<<<<< HEAD
func resetGameState(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {
=======
func resetGameState(DB *mysql_db.DB, tx *sql.Tx, tablesTableName, playersTableName, pokerTablesTableName, tableID string) {
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	_, _, dealer, dealerSeatNumber := gameinfo.GetDealerAndHighestBidder(DB, playersTableName, pokerTablesTableName, tableID)

	setOperation := "current_player_making_move = "
<<<<<<< HEAD
	gameflow.SetNextAvailablePlayerAfterThisOne(DB, tablesTableName, playersTableName, tableID, dealer, dealerSeatNumber, setOperation)
=======
	gameflow.SetNextAvailablePlayerAfterThisOne(DB, tx, tablesTableName, playersTableName, tableID, dealer, dealerSeatNumber, setOperation)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	query := fmt.Sprintf(`UPDATE %s
	                      SET game_in_progress = false
						  WHERE table_id = %s`, tablesTableName, tableID)


	_, err := db.Exec(query)

	utils.CheckError(err)
<<<<<<< HEAD

=======
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
}