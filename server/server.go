package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/pat"
	"github.com/facebookgo/grace/gracehttp"
	"encoding/json"
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

	return r
}




func SpamHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	text := req.Form["text"]

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: SpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("SpamHandler, -->  ", text)

	//classify spam text
	reply, err := Client.Check(text[0])
	if err != nil {
		log.Println("Error: SpamHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}else {
		log.Println(reply.Code)
		log.Println(reply.Message)
		log.Println(reply.Vars)
	}

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		log.Println("Error: SpamHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}