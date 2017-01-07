package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func SpamassassinReportSpamHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	text := req.Form["text"]

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: ReportSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("RevokeSpamHandler, -->  ", text)

	reply, err := SpamassassinClient.Report(text[0])

	if err != nil {
		log.Println("Error: ReportSpamHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		log.Println(reply.Code)
		log.Println(reply.Message)
		log.Println(reply.Vars)
	}

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		log.Println("Error: ReportSpamHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func SpamassassinRevokeSpamHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	text := req.Form["text"]

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: RevokeSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("RevokeSpamHandler, -->  ", text)

	reply, err := SpamassassinClient.RevokeSpam(text[0])

	if err != nil {
		log.Println("Error: RevokeSpamHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		log.Println(reply.Code)
		log.Println(reply.Message)
		log.Println(reply.Vars)
	}

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		log.Println("Error: RevokeSpamHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func SpamassassinSpamHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
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
	reply, err := SpamassassinClient.Check(text[0])
	if err != nil {
		log.Println("Error: SpamHandler Marshal -> ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
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
