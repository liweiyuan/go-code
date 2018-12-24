package controller

import (
	"net/http"
	"log"
)

//中间层验证
func midleAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		username, err := getSessionUser(r)
		log.Println("middle:", username)
		if err != nil {
			log.Println("middle get session err and redirect to login")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}

	}
}
