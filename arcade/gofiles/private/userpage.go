package private

import (
	"html/template"
	"net/http"

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
		// gameSetup := poker.HandlePokerForm(w, r)
		// if gameSetup {
		//     http.Redirect(w, r, "pokertable", http.StatusMovedPermanently)
		// }
	}
}

func attachPage( w http.ResponseWriter, htmlName string, userPage UserPage) {
	var tpl *template.Template
	tpl, _ = template.ParseGlob("templates/private/" + htmlName)
	tpl.ExecuteTemplate(w, htmlName, userPage)
}