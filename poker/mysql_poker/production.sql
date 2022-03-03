
-----------------------------------------------------------------------
--- PRODUCTION BELOW ------
---------------------------

DROP TABLE tables;

CREATE TABLE tables (
    table_id INT NOT NULL AUTO_INCREMENT,
    time_since_last_move TIMESTAMP NOT NULL,
    current_player_making_move VARCHAR(255) NOT NULL,
    deck VARCHAR(255),
    cards_not_in_deck VARCHAR(255),
    game_in_progress BOOLEAN,
    PRIMARY KEY (table_id)
);

DROP TABLE poker_tables;

CREATE TABLE poker_tables (
    table_id INT NOT NULL,
    community_cards VARCHAR(32),
    highest_bidder VARCHAR(255),
    highest_bid DECIMAL(15,2),
    dealer VARCHAR(255),
    PRIMARY KEY (table_id)
);


DROP TABLE players;

CREATE TABLE players (
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

---------------------------
--- END OF PRODUCTION -----
-----------------------------------------------------------------------
SELECT *
FROM user_credentials;

SELECT *
FROM tables;

SELECT *
FROM cs2208_jr30.players;

SELECT *
FROM cs2208_jr30.poker_tables;