package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"tugas14/config"
	"tugas14/data"
)

var webPort = "8080"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
	conn := config.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to database!")
	}

	app := Config{
		DB: conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(), 
	}

	fmt.Println(fmt.Sprintf("server running at http://localhost:%s", webPort))
	
	err:= srv.ListenAndServe()
	if err != nil {log.Panic(err)}
}

