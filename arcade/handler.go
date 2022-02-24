package main

import (
	"fmt"
	"net/http"

	"github.com/cs3305/group13_2022/project/arcade/gofiles/private"
	"github.com/cs3305/group13_2022/project/arcade/gofiles/public"
	"github.com/cs3305/group13_2022/project/blackjack"
	"github.com/cs3305/group13_2022/project/poker"
	"github.com/cs3305/group13_2022/project/sidescroller"
)

func main() {
	http.HandleFunc("/", public.HandleLoginPage)
	http.HandleFunc("/signup", public.HandleSignupPage)

	http.HandleFunc("/userpage", private.HandleUserPage)

	pokerHandlers()
	sidescrollerHandlers()
	blackjackHandlers()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    fmt.Print("\033[H\033[2J")  // Tidies up terminal 
	
	fmt.Println("")
	fmt.Println("##################################################")
    fmt.Println("Server running on localhost: http://localhost:9000")
    fmt.Println("##################################################")

	http.ListenAndServe("localhost:9000", nil)
}

func pokerHandlers() {
	http.HandleFunc("/pokertable", poker.HandlePokerTableRequest)
	http.HandleFunc("/content_request", poker.HandleContentAjaxRequest)
}
func sidescrollerHandlers() {
	http.HandleFunc("/adventure_game", sidescroller.HandleSideScrollerPage)
}
func blackjackHandlers() {
	http.HandleFunc("/blackjack", blackjack.HandleBlackjackPage)
}
