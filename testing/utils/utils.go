package utils

import (
	"net/http"
	"net/http/httptest"

	"github.com/cs3305/group13_2022/project/utils/token"
)

// FUNCTIONS BELOW: Helper functions for appending cookie to Requests and Responses.


func CreateRequestWithPokerCookie() *http.Request {
	w := CreateResponseWithCookie("bob", "poker", "1", "1", "30.0")

	request := "/"
	r := httptest.NewRequest(http.MethodGet, request, nil)

	// Add the cookie to the request. THIS MAY BE INCORRECT.
	r.AddCookie(w.Result().Cookies()[0])

	return r
}

// Creates a request with a plain token with username `Dave`
// 
// param: A submitted form request.
// 
// param example:
//     request := "/pokertable?poker=online&tableCode=17"
//     r := CreateRequestWithCookie( request )
func CreateRequestWithCookie(request string) *http.Request {
	w := CreateResponseWithCookie("Dave", "", "", "", "0")

	r := httptest.NewRequest(http.MethodGet, request, nil)

	// Add the cookie to the request. THIS MAY BE INCORRECT.
	r.AddCookie(w.Result().Cookies()[0])

	return r
}

func CreateResponseWithCookie(username, gameType, tableID, seatNumber, funds string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()  // NewRecorder works like ResponseWriter apparently.

	claims, expirationTime := token.NewDefaultClaims(username, gameType, tableID, seatNumber, funds)

	token.AppendTokenCookie(w, "token", claims, expirationTime)

	return w
}

func CreateRegularResponse() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

// ################################################################################