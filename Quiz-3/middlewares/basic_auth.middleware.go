package middlewares

import "net/http"

func BasicAuth(r *http.Request) string {
	uname, pwd, ok := r.BasicAuth()
	if !ok {
		return "Username atau Password tidak boleh kosong"
	}

	if uname == "admin" && pwd == "password" {
		return ""
	} else if uname == "editor" && pwd == "secret" {
		return ""
	} else if uname == "trainer" && pwd == "rahasia" {
		return ""
	} else {
		return "Username atau Password tidak sesuai"
	}
}
