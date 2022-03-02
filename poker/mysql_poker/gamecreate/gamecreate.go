package gamecreate

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project/mysql_db/change"
	"github.com/cs3305/group13_2022/project/mysql_db/insert"
)

func AssignNewTable( tx *sql.Tx, tablesTableName, pokerTablesTableName, username string ) (tableID string) {
	tableID, found := changeOldTable(tx, tablesTableName, username)

	if found {
		changed := changeOldPokerTable(tx, pokerTablesTableName, tableID, username)

		if ! changed {
			insertNewPokerTable(tx, pokerTablesTableName, tableID, username)
		}

		return tableID
	} else {
		// A stale entry wasn't found then we just insert a new row.
		tableID, inserted := insertNewTable(tx, tablesTableName, username)
		if ! inserted {
			// Add a w.Writer([]byte(message)) here?
			panic("Table entry was not inserted.")
		} else {
			inserted = insertNewPokerTable(tx, pokerTablesTableName, tableID, username)
			if ! inserted {
				// Add a w.Writer([]byte(message)) here?
				panic("Poker table entry was not inserted.")
			}
			
			return tableID
		}
	}
}


func changeOldTable(tx *sql.Tx, tablesTableName, username string) (tableID string, changed bool){
	deck := cards.NewDeck(1)
	emptyDeck := cards.NewDeck(0)
	gameInProgress := false

	setOperation := fmt.Sprintf(`time_since_last_move = CURRENT_TIMESTAMP,
	                             current_player_making_move = "%s",
								 deck = "%s",
								 cards_not_in_deck = "%s",
								 game_in_progress = %t`, username, deck, emptyDeck, gameInProgress)


	// Lets try and reuse a stale game table entry.
	MAX_INACTIVITY_OF_TABLE_IN_HOURS := "24"
	id, changed := change.CreateRowFromRowOlderThan(tx, tablesTableName, "table_id", "time_since_last_move", MAX_INACTIVITY_OF_TABLE_IN_HOURS, setOperation)

	tableID = strconv.FormatInt(id, 10)
	return tableID, changed
}
func changeOldPokerTable(tx *sql.Tx, pokerTablesTableName, tableID, dealer string) (changed bool) {
	columnNames := "table_id, community_cards, highest_bidder, highest_bid, dealer"
	values := fmt.Sprintf( "%s, '', '', 0.0, '%s'", tableID, dealer )
    
	numOfRowsAffected, inserted := insert.InsertTableEntry(tx, pokerTablesTableName, columnNames, values)
	if ! inserted {
		return false
	} else {
		if numOfRowsAffected > 1 {
			panic("Only one row should have been changed in table.")
		} else {
		    return true
		}
	}
}

func insertNewTable(tx *sql.Tx, tablesTableName, username string) (tableID string, inserted bool) {
	deck := cards.NewDeck(1)
	emptyDeck := cards.NewDeck(0)
	gameInProgress := false

	columnNames := "table_id, time_since_last_move, current_player_making_move, deck, cards_not_in_deck, game_in_progress"
	values := fmt.Sprintf(`DEFAULT, CURRENT_TIMESTAMP, "%s", "%s", "%s", %t`, username, deck, emptyDeck, gameInProgress)

	id, inserted := insert.InsertTableEntry(tx, tablesTableName, columnNames, values)
	
	if ! inserted {
		panic("Table was not inserted.")
	}
	tableID = strconv.FormatInt(id, 10)

	return tableID, inserted
}
func insertNewPokerTable(tx *sql.Tx, pokerTablesTableName, tableID, username string) (inserted bool) {
	communityCards := ""
	highestBidder := ""
	highestBid := "0.0"
	dealer := username

	columnNames := `table_id, community_cards, highest_bidder, highest_bid, dealer`
	values := fmt.Sprintf(`%s, "%s", "%s", %s, "%s"`, tableID, communityCards, highestBidder, highestBid, dealer)
	
	_, inserted = insert.InsertTableEntry(tx, pokerTablesTableName, columnNames, values)
	if ! inserted {
		panic("Poker table was not inserted.")
	}

	return inserted
}