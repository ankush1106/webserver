package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ank1106/webserver/src/main/controller"
	"github.com/ank1106/webserver/src/main/models"
)

type WebServer struct {
	host string
	port string
}

func (w *WebServer) Run(db *models.DBClient) {
	w.port = os.Getenv("PORT")
	w.host = os.Getenv("HOST")
	if w.port == "" {
		w.port = "3000"
	}
	if w.host == "" {
		w.host = "localhost"
	}

	mux := http.NewServeMux()
	cnt := controller.NewController(db)
	fs := http.FileServer(http.Dir("src/main/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/favicon.ico", http.StripPrefix("", fs))

	mux.HandleFunc("/", cnt.IndexHandler)
	mux.HandleFunc("/login/", cnt.LoginHandler)
	mux.HandleFunc("/register/", cnt.RegisterHandler)
	mux.HandleFunc("/todos/", cnt.GeTODOs)
	fmt.Println("Server running on " + w.host + ":" + w.port)
	http.ListenAndServe(w.host+":"+w.port, mux)
}
