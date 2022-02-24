package public

import (
	"html"
	"html/template"
	"net/http"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/mysql_db/crypto"

	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/utils/token"
)

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	envs := env.GetEnvironmentVariables("../../../.env")
	if loginFormOk( w, r, envs ) {
		username := html.EscapeString(r.Form.Get("username"))

		claims, expirationTime := token.NewDefaultClaims(username, "", "", "", "0.0")
		
		tokenName := "token"
		token.AppendTokenCookie(w, tokenName, claims, expirationTime)

		http.Redirect(w, r, "/userpage", http.StatusMovedPermanently)
	}
}

type loginProblems struct {
	GeneralProblem string
}

func loginFormOk( w http.ResponseWriter, r *http.Request, envs map[string]string ) bool {
	var problems loginProblems

	r.ParseForm()
	if len(r.Form) == 0 {
		attachLoginPage( w, problems )
		return false
	} 

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	username = html.EscapeString(username)
	password = html.EscapeString(password)

	if username == "" || password == "" {
		problems.GeneralProblem = "Please fill all necessary fields."
		attachLoginPage( w, problems )
		return false
	}

	DB := mysql_db.NewDB(envs)
	usersTableName := envs["USER_CREDENTIALS_TABLE"]

	match := crypto.CredentialsMatch(DB, usersTableName, "username", username, password)

	if match == false {
		problems.GeneralProblem = "Username or password provided was incorrect."
		attachLoginPage( w, problems )
		return false
	}


	// form data fully validated, user may access protected content.
	return true
}

func attachLoginPage( w http.ResponseWriter, problems loginProblems ) {
	var tpl *template.Template

	tpl, _ = template.ParseGlob("templates/public/login.html")
	tpl.ExecuteTemplate(w, "login.html", problems)
}
