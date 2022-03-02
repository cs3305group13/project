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

func GetNumberOfPlayersAtTable( DB *mysql_db.DB, playersTableName, tableCode string ) (numOfRows int) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	var query = fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE table_id = %s;`, playersTableName, tableCode)

	err := db.QueryRow(query).Scan(&numOfRows)
	
	if err != nil {
		log.Fatal(err)
	}

	return
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

func GetPlayersFunds(tx *sql.Tx, playersTableName, username string) (funds string) {
	query := fmt.Sprintf(`SELECT funds
	                      FROM %s
						  WHERE username = "%s"`, playersTableName, username)

	err := tx.QueryRow(query).Scan(&funds)
	utils.CheckError(err)

	return funds
}package token

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/cs3305/group13_2022/project/utils/env"
)

var envs = env.GetEnvironmentVariables("../../../../../production.env")
var jwtKey = []byte(envs["TOKEN_KEY"])

var TIME_UNTIL_EXPIRY, _ = strconv.Atoi(envs["TOKEN_LIFE"])

type Claims struct {
	Username string `json:"username"`
	GameType string `json:"gameType"`
	TableID string `json:"tableID"`
	SeatNumber string `json:"seatNumber"`
	Funds string `json:"funds"`
	jwt.StandardClaims
}


// Creates default claims with additional username field and expiration time in minutes.
func NewDefaultClaims(username, gameType, tableID, seatNumber, funds string) (*Claims, time.Time) {

	timeToLive := time.Duration(TIME_UNTIL_EXPIRY) * time.Minute
	expirationTime := time.Now().Add(timeToLive)

	claims := &Claims{
		              Username: username,
					  GameType: gameType,
					  TableID: tableID,
					  SeatNumber: seatNumber,
					  Funds: funds,
                      StandardClaims: jwt.StandardClaims{ 
						                                ExpiresAt: expirationTime.Unix(),	
								                        },
					}
	return claims, expirationTime
}

// Appends token cookie to http.ResponseWriter
func AppendTokenCookie( w http.ResponseWriter, tokenName string, claims *Claims, expirationTime time.Time ) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(jwtKey)

	// Cookie may be dropped (ex. Something was wrong in tokenString implementation)
	http.SetCookie(w, &http.Cookie{
		Name:    tokenName,
		Value:   tokenString,
		Expires: expirationTime,
		HttpOnly: true,
	})
}

// Validates that the token is readable and that it hasn't expired yet. AND attaches refreshed token to response. If a problem is encountered the message is writen to response writer.
func TokenValid(w http.ResponseWriter, r *http.Request, refresh bool) bool {

	claims, err := GetClaimsFromCookie(r, "token")
	if err != nil {
		w.Write([]byte("Sorry you do not have access, please log in or sign up."))
		return false
	}

	currentTime := time.Now().Unix()
	tokenExpirationTime := claims.StandardClaims.ExpiresAt

	tokenTimeToLive := currentTime - tokenExpirationTime

	
	if tokenTimeToLive > 0 {  // tokenTimeToLive returns a number below or equal to zero if the token hasn't expired yet.
	    w.Write([]byte("You've been logged out, please sign in."))
		return false
	}

	if refresh == true {
		refreshExpirationTime(w, r, "token")
	}
	
	// End of validation, All statements should have been skipped if valid.
	return true
}

// Updates tokens expiration date with the current time.
func refreshExpirationTime( w http.ResponseWriter, r *http.Request, tokenName string ) (err error) {
	claims, err := GetClaimsFromCookie(r, tokenName)
	if err != nil {
		fmt.Println("REMINDER: This is the error that came up: ", err)
		return
	}
	newExpirationTime := (time.Now().Add(time.Duration(TIME_UNTIL_EXPIRY) * time.Minute))
	fmt.Println("Time to live of refresh --> ", newExpirationTime)

	// Updates both just to reduce bugs later.
	claims.ExpiresAt = newExpirationTime.Unix()
	claims.StandardClaims.ExpiresAt = newExpirationTime.Unix()

	AppendTokenCookie( w, "token", claims, newExpirationTime )  // REMINDER: Cookie may be silently dropped.
	return nil
}

// Retrieves all claims stored in the token.
func GetClaimsFromCookie( r *http.Request, cookieName string ) (*Claims, error) {
	c, err := r.Cookie(cookieName)

	if err != nil {
		return nil, err
	}

	tokenString := c.Value

	token, err := jwt.ParseWithClaims( tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}


func GetUsername( r *http.Request, tokenName string ) string {
	claims, _ := GetClaimsFromCookie( r, tokenName )

	return claims.Username
}
func GetGameType( r *http.Request, tokenName string ) string {
	claims, _ := GetClaimsFromCookie( r, tokenName )

	return claims.GameType
}
func GetTableID( r *http.Request, tokenName string ) string {
	claims, _ := GetClaimsFromCookie( r, tokenName )

	return claims.TableID
}
func GetSeatNumber( r *http.Request, tokenName string ) string {
	claims, _ := GetClaimsFromCookie( r, tokenName )

	return claims.SeatNumber
}
func GetFunds( r *http.Request, tokenName string ) string {
	claims, _ := GetClaimsFromCookie( r, tokenName )

	return claims.Funds
}