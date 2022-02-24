package blackjack

import (
	"html/template"
	"net/http"

	"github.com/cs3305/group13_2022/project/utils/token"
)

type User struct {
	Username string
}

func HandleBlackjackPage(w http.ResponseWriter, r *http.Request) {

	if token.TokenValid(w, r, true) {
		user := User{Username : token.GetUsername(r, "token") }

	    var tpl *template.Template
	    tpl, _ = template.ParseGlob("templates/private/blackjack.html")
	    tpl.ExecuteTemplate(w, "blackjack.html", user)
	}
	
}