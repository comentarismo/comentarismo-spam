package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/pat"
	"github.com/facebookgo/grace/gracehttp"
	"comentarismo-spam/spamc"
)

var (
	router *pat.Router
	Client *spamc.Client
)

type WebError struct {
	Error string
}


func init() {
	Client = spamc.New("127.0.0.1:783", 10)
}

//NewServer return pointer to new created server object
func NewServer(Port string) *http.Server {
	router = InitRouting()
	return &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}
}

//StartServer start and listen @server
func StartServer(Port string) {
	log.Println("Starting server")
	s := NewServer(Port)
	fmt.Println("Server starting --> " + Port)
	err := gracehttp.Serve(
		s,
	)
	if err != nil {
		log.Fatalln("Error: %v", err)
		os.Exit(0)
	}

}

func InitRouting() *pat.Router {

	r := pat.New()

	/** CREATE NEW COMMENT **/
	r.Post("/spam", SpamHandler)
	r.Post("/revoke", RevokeSpamHandler)
	r.Post("/report", ReportSpamHandler)

	return r
}
