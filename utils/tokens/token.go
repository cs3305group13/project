package token

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/cs3305/group13_2022/project/utils/env"
)

var envs = env.GetEnvironmentVariables("../../../../../.env")
var jwtKey = []byte(envs["TOKEN_KEY"])

var TIME_UNTIL_EXPIRY, _ = strconv.Atoi(envs["TOKEN_LIFE"])

type Claims struct {
	Username   string `json:"username"`
	GameType   string `json:"gameType"`
	TableID    string `json:"tableID"`
	SeatNumber string `json:"seatNumber"`
	Funds      string `json:"funds"`
	jwt.StandardClaims
}

// Creates default claims with additional username field and expiration time in minutes.
func NewDefaultClaims(username, gameType, tableID, seatNumber, funds string) (*Claims, time.Time) {

	timeToLive := time.Duration(TIME_UNTIL_EXPIRY) * time.Minute
	expirationTime := time.Now().Add(timeToLive)

	claims := &Claims{
		Username:   username,
		GameType:   gameType,
		TableID:    tableID,
		SeatNumber: seatNumber,
		Funds:      funds,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	return claims, expirationTime
}

// Appends token cookie to http.ResponseWriter
func AppendTokenCookie(w http.ResponseWriter, tokenName string, claims *Claims, expirationTime time.Time) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(jwtKey)

	// Cookie may be dropped (ex. Something was wrong in tokenString implementation)
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    tokenString,
		Expires:  expirationTime,
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

	if tokenTimeToLive > 0 { // tokenTimeToLive returns a number below or equal to zero if the token hasn't expired yet.
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
func refreshExpirationTime(w http.ResponseWriter, r *http.Request, tokenName string) (err error) {
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

	AppendTokenCookie(w, "token", claims, newExpirationTime) // REMINDER: Cookie may be silently dropped.
	return nil
}

// Retrieves all claims stored in the token.
func GetClaimsFromCookie(r *http.Request, cookieName string) (*Claims, error) {
	c, err := r.Cookie(cookieName)

	if err != nil {
		return nil, err
	}

	tokenString := c.Value

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func GetUsername(r *http.Request, tokenName string) string {
	claims, _ := GetClaimsFromCookie(r, tokenName)

	return claims.Username
}
func GetGameType(r *http.Request, tokenName string) string {
	claims, _ := GetClaimsFromCookie(r, tokenName)

	return claims.GameType
}
func GetTableID(r *http.Request, tokenName string) string {
	claims, _ := GetClaimsFromCookie(r, tokenName)

	return claims.TableID
}
func GetSeatNumber(r *http.Request, tokenName string) string {
	claims, _ := GetClaimsFromCookie(r, tokenName)

	return claims.SeatNumber
}
func GetFunds(r *http.Request, tokenName string) string {
	claims, _ := GetClaimsFromCookie(r, tokenName)

	return claims.Funds
}
