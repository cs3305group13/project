package main

import (
	"net/http"

	"github.com/cs3305/group13_2022/project/arcade/gofiles/public"
)

func main() {
	http.HandleFunc("/", public.HandleLoginPage)
	http.HandleFunc("/signup", public.HandleSignupPage)

	http.HandleFunc("/userpage", public.HandleUserPage)


	http.ListenAndServe("localhost:9000", nil)
}