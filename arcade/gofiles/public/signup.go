package public

import (
	"html"
	"html/template"
	"net/http"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/mysql_db/crypto"
	"github.com/cs3305/group13_2022/project/mysql_db/find"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamejoin"

	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/utils/token"
)

// If url contains `/signup` then request is directed here for handling
func HandleSignupPage(w http.ResponseWriter, r *http.Request) {
	envs := env.GetEnvironmentVariables("../../../production.env")

	if signupFormOk(w, r, envs) { 
		username := html.EscapeString(r.Form.Get("username"))

		// make the cookie token claims.
		claims, expirationTime := token.NewDefaultClaims(username, "", "", "", "0.0")

		tokenName := "token"
		token.AppendTokenCookie(w, tokenName, claims, expirationTime)

		http.Redirect(w, r, "/userpage", http.StatusMovedPermanently)
	}
}

// This struct is used within the html template to inject information about why the user wasn't given access.
type signupProblems struct {
	UsernameProblem string
	PasswordProblem string
	GeneralProblem string
}


// function validates user form and only returns false when it detected an error and wrote to http.ResponseWriter 
func signupFormOk(w http.ResponseWriter, r *http.Request, envs map[string]string) bool {
	var problems signupProblems

	r.ParseForm()
	if len(r.Form) == 0 {
		attachSignupPage(w, problems)
		return false
	}

	// Get form inputs.
	username := r.Form.Get("username")
	password_00 := r.Form.Get("password_00")
	password_01 := r.Form.Get("password_01")

	// Escape these form inputs.
	username = html.EscapeString(username)
	password_00 = html.EscapeString(password_00)
	password_01 = html.EscapeString(password_01)

	if username == "" || password_00 == "" || password_01 == "" {
		problems.GeneralProblem = "Please fill all necessary fields."
		attachSignupPage( w, problems )
		return false
	}

	if password_00 != password_01 {
		problems.PasswordProblem = "Passwords do not match."
		attachSignupPage( w, problems )
		return false
	}
	if len(password_00) < 8 {
		problems.PasswordProblem = "Password must be at least 8 characters long."
		attachSignupPage( w, problems )
		return false
	}

	DB := mysql_db.NewDB(envs)

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	usersTableName := envs["USER_CREDENTIALS_TABLE"]

	// Checks if these credentials already exist
	userFound := find.FindRowByValue(tx, usersTableName, "username", username, "username")

	if userFound != "" {  // if user found then signup page is reattached along with reason for failure

		problems.GeneralProblem = `The username provided already has an account. Login <a href="/signup">Here</a>`
		attachSignupPage( w, problems )
		return false
	}

	successfullyAdded := crypto.AddUser(tx, usersTableName, username, password_00)
	playerFunds := envs["PLAYER_FUNDS"]

	playersTableName := envs["PLAYERS_TABLE"]
	successfullyAddedPlayer := gamejoin.AddPlayer(tx, playersTableName, username, playerFunds)

	if ! successfullyAdded || ! successfullyAddedPlayer {  // if user is unsuccessfully added then signup page along with reason for failure are attached

		problems.GeneralProblem = "We experienced a problem signing you up. Try again or come back later."
		attachSignupPage( w, problems )
		return false
	}

	tx.Commit() // commit transaction

	// If at this point all went well, user granted access and also NOTHING was written to `w` responseWriter.
	return true
}

// Attaches signup.html to response along with any problems inserted into the html.
func attachSignupPage(w http.ResponseWriter, problems signupProblems) {
	var tpl *template.Template

	tpl, _ = template.ParseGlob("templates/public/signup.html")
	tpl.ExecuteTemplate(w, "signup.html", problems)
}