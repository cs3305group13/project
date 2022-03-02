SELECT *
FROM cs2208_jr30.dummy_tables;

SELECT *
FROM cs2208_jr30.dummy_poker_tables;

SELECT *
FROM cs2208_jr30.dummy_players
ORDER BY seat_number;
-----------------------------------------------------------------------
--- TESTING BELOW ---------
---------------------------
DROP TABLE cs2208_jr30.dummy_tables;

CREATE TABLE cs2208_jr30.dummy_tables (
    table_id INT NOT NULL AUTO_INCREMENT,
    time_since_last_move TIMESTAMP NOT NULL,
    current_player_making_move VARCHAR(255) NOT NULL,
    deck VARCHAR(255),
    cards_not_in_deck VARCHAR(255),
    game_in_progress BOOLEAN,
    PRIMARY KEY (table_id)
);

DROP TABLE cs2208_jr30.dummy_poker_tables;

CREATE TABLE cs2208_jr30.dummy_poker_tables (
    table_id INT NOT NULL,
    community_cards VARCHAR(32),
    highest_bidder VARCHAR(255),
    highest_bid DECIMAL(15,2),
    dealer VARCHAR(255),
    PRIMARY KEY (table_id)
);

DROP TABLE cs2208_jr30.dummy_players;

CREATE TABLE cs2208_jr30.dummy_players (
    username VARCHAR(255),
    funds DECIMAL(15, 2),
    table_id INT,
    seat_number INT,
    player_state VARCHAR(25),
    player_cards VARCHAR(10),
    money_in_pot DECIMAL(15, 2),
    time_since_request TIMESTAMP NOT NULL,
    PRIMARY KEY (username)
);


-- TestCheckTableExists(),
-- TestGetNumberOfPlayersAtTable()
DELETE FROM cs2208_jr30.dummy_players;

INSERT INTO cs2208_jr30.dummy_players 
VALUES 
       ("derek", 30.0, 1, 1, "PLAYING", "", 0.0, CURRENT_TIMESTAMP()),
       ("jason", 30.0, 1, 2, "PLAYING", "", 0.0, CURRENT_TIMESTAMP()),
       ("john", 30.0, 1, 3, "PLAYING", "", 0.0, CURRENT_TIMESTAMP()),
       ("barry", 30.0, 1, 4, "PLAYING", "", 0.0, CURRENT_TIMESTAMP()),
       ("ahmed", 30.0, 1, 5, "PLAYING", "", 0.0, CURRENT_TIMESTAMP()),
       ("laura", 30.0, 1, 6, "PLAYING", "", 0.0, CURRENT_TIMESTAMP()),
       ("alejandro", 30.0, 1, 7, "PLAYING", "", 0.0, CURRENT_TIMESTAMP()),
       ("dan", 30.0, 1, 8, "PLAYING", "", 0.0, CURRENT_TIMESTAMP());

 

DELETE FROM cs2208_jr30.dummy_tables;
INSERT INTO cs2208_jr30.dummy_tables 
VALUES (1, DATE_SUB(NOW(), INTERVAL 48 HOUR), "barry", "AH2H3H4H5H6H7H8H9H10HJHQHKHAD2D3D4D5D6D7D8D9D10DJDQDKDAS2S3S4S5S6S7S8S9S10SJSQSKSAC2C3C4C5C6C7C8C9C10CJCQCKC", "", false);


DELETE FROM cs2208_jr30.dummy_poker_tables;
INSERT INTO cs2208_jr30.dummy_poker_tables
VALUES (1, "", "john", 1.0, "derek");



SELECT highest_bidder, dealer
FROM cs2208_jr30.dummy_poker_tables
WHERE table_id = 1;

SELECT current_player_making_move
FROM cs2208_jr30.dummy_tables
WHERE table_id = 1;


----------- BIN -------------
SELECT COUNT(*)
FROM cs2208_jr30.dummy_players
WHERE table_id = 1 AND
	  player_state IN ( "PLAYING", "ALL_IN", "RAISED", "CALLED", "CHECKED");
