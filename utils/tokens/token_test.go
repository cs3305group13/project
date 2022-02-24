package token

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestEnvKey(t *testing.T) {
	if string(jwtKey) == "" {
		t.Error("jwtKey not retrieved correctly or isn't declared in .env file.")
	}
	fmt.Println(string(jwtKey))
}

// Not a true test, only shows token is appended to a ResponseWriter.
func TestAppendTokenCookieToResponse(t *testing.T) {
	w := httptest.NewRecorder()  // NewRecorder works like ResponseWriter apparently.


	expirationTime := time.Now()
	claims := &Claims{Username: "bob",
                     StandardClaims: jwt.StandardClaims{
						 ExpiresAt: expirationTime.Unix(),
					 },}

	// fmt.Println(w.Header())  // DoNt remove these print statements.
	AppendTokenCookie(w, "token", claims, expirationTime)
	// fmt.Println(w.Header())
}

// Not a true test, only shows claims are retrieved from Request.
func TestRetrieveClaimsFromCookieInRequest(t *testing.T) {
	r := createRequestWithCookie()

	claimsJWT, err := GetClaimsFromCookie(r, "token")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(claimsJWT.Username)
}

func TestTokenValid(t *testing.T) {
	r := createRequestWithCookie()
	w := createResponseWithCookie()

	// time.Sleep(1*time.Second)  // To test this change expiration time associated with createResponseWithCookie

	accessGranted := TokenValid(w, r, true)
	if ! accessGranted {
		t.Error(w.Body)
	}
}

func TestRefreshExpirationTime(t *testing.T) {
	r := createRequestWithCookie()
	w := httptest.NewRecorder()

	requestClaims, err := GetClaimsFromCookie( r, "token" )
	if err != nil {
		t.Error("Problem retrieving claims from cookie")
	}
	err = refreshExpirationTime(w, r, "token")
	if err != nil {
		t.Error(err)
	}
	newClaimsAttachedToResponse := w.Result().Cookies()
	
	// Im not sure how to retrieve expiration time from ResponseRecorder/ResponseWriter
	fmt.Println(requestClaims.StandardClaims.ExpiresAt)
	fmt.Println(newClaimsAttachedToResponse)
}

func TestGetters(t *testing.T) {
	r := createRequestWithCookie()

	username := GetUsername( r, "token" )
	gameType := GetGameType( r, "token" )
	tableID := GetTableID( r, "token" )
	seatNumber := GetSeatNumber( r, "token" )

	if username == "" && gameType != "" && tableID != "" && seatNumber != "" {
		t.Error("Only username should contain data")
	}
}


// FUNCTIONS BELOW: Helper functions for appending cookie to Requests and Responses.
func createRequestWithCookie() *http.Request {
	w := createResponseWithCookie()

	r := httptest.NewRequest(http.MethodGet, "/login?username=hello", nil)

	// Add the cookie to the request. THIS MAY BE INCORRECT.
	r.AddCookie(w.Result().Cookies()[0])

	return r
}

func createResponseWithCookie() *httptest.ResponseRecorder {
	w := httptest.NewRecorder()  // NewRecorder works like ResponseWriter apparently.

	claims, expirationTime := NewDefaultClaims("Dave", "", "", "", "0")

	AppendTokenCookie(w, "token", claims, expirationTime)

	return w
}
// ################################################################################