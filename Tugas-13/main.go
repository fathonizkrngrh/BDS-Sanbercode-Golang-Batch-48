package main

import (
	"fmt"
	"net/http"
)

var webPort = "8080"

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
	}

	http.HandleFunc("/nilai-mahasiswa", Authenticate(HandlerNilaiMahasiswa))

	fmt.Println(fmt.Sprintf("server running at http://localhost:%s", webPort))
	server.ListenAndServe()
}