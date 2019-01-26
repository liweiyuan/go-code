package controller

import (
	"net/http"
	"log"
	"github.com/liweiyuan/go-code/model"
)

//中间层验证
func middleAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		username, err := getSessionUser(r)
		log.Println("middle:", username)
		if username != "" {

			//更新最后登陆的时间
			log.Println("Last seen:", username)
			model.UpdateLastSeen(username)
		}
		if err != nil {
			log.Println("middle get session err and redirect to login")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}

	}
}
