package sidescroller

import (
	"html/template"
	"net/http"

	"github.com/cs3305/group13_2022/project/utils/token"
)

type User struct {
	Username string
}

func HandleSideScrollerPage(w http.ResponseWriter, r *http.Request) {
	if token.TokenValid(w, r, true) {
		user := User{Username : token.GetUsername(r, "token") }

		var tpl *template.Template
	    tpl, _ = template.ParseGlob("templates/private/sidescroller.html")
	    tpl.ExecuteTemplate(w, "sidescroller.html", user)
	}
}