package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cs3305/group13_2022/project/arcade/gofiles/private"
	"github.com/cs3305/group13_2022/project/arcade/gofiles/public"
	"github.com/cs3305/group13_2022/project/blackjack"
	"github.com/cs3305/group13_2022/project/poker"
	"github.com/cs3305/group13_2022/project/sidescroller"
)

func main() {
	
	// URL paths.
	http.HandleFunc("/", public.HandleLoginPage)
	http.HandleFunc("/signup", public.HandleSignupPage)

	http.HandleFunc("/userpage", private.HandleUserPage)

	http.HandleFunc("/pokertable", poker.HandlePokerTableRequest)
	http.HandleFunc("/content_request", poker.HandleContentAjaxRequest)
	http.HandleFunc("/user_request", poker.HandleUserAjaxRequest)

	http.HandleFunc("/adventure_game", sidescroller.HandleSideScrollerPage)

	http.HandleFunc("/blackjack", blackjack.HandleBlackjackPage)
	// URL paths.


	// Distributes css, js, img files when requested in html file.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    fmt.Print("\033[H\033[2J")  // clears terminal window
	
	fmt.Println("")
	fmt.Println("##################################################")
    fmt.Println("Server running on localhost: http://localhost:9000")
    fmt.Println("##################################################")

	// Listens out for requests on port 9000 localhost
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}


