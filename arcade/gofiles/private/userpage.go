package private

import (
	"fmt"
	"html"
	"html/template"
	"net/http"

	"github.com/cs3305/group13_2022/project/poker"
	"github.com/cs3305/group13_2022/project/utils/token"
)

func HandleUserPage(w http.ResponseWriter, r *http.Request) {
	if token.TokenValid(w, r, true) {
		checkForm(w, r)
	}
}


type UserPage struct {
	Username string
}

func checkForm( w http.ResponseWriter, r *http.Request ) {

	r.ParseForm()
	if len(r.Form) == 0 {
		username := token.GetUsername( r, "token" )
		attachPage( w, "userpage.html", UserPage{ Username: username })
		return
	} else {
		submitValue := html.EscapeString(r.FormValue("submit"))
		switch submitValue {
		case "Poker" :
            gameSetup := poker.HandlePokerForm(w, r)
			if gameSetup {
				http.Redirect(w, r, "pokertable", http.StatusMovedPermanently)
			}

		case "Royal Runner" :
			http.Redirect(w, r, "adventure_game", http.StatusMovedPermanently)

	    case "Blackjack" :
			http.Redirect(w, r, "blackjack", http.StatusMovedPermanently)

	    default :
		    fmt.Printf("Problem; Couldn't match users choice '%s' with any option", submitValue)
		}
	}
}

func attachPage( w http.ResponseWriter, htmlName string, userPage UserPage) {
	var tpl *template.Template
	tpl, _ = template.ParseGlob("templates/private/" + htmlName)
	tpl.ExecuteTemplate(w, htmlName, userPage)
}