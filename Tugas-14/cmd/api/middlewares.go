package main

// import (
// 	"errors"
// 	"net/http"
// )

// var username = "admin"
// var password = "admin"

// func Authenticate(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		user, pass, ok := r.BasicAuth()
// 		if !ok {
// 			WriteResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", errors.New("Username atau Password tidak boleh kosong").Error(), nil)
// 			return
// 		}
// 		if user != username || pass != password {
// 			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
// 			WriteResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", errors.New("You are not allowed to access this source").Error(), nil)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	}
// }