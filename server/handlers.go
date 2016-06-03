package server

import (
	"net/http"
	"log"
	"encoding/json"
	"comentarismo-spam/spamc"
)

func ReportSpamHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	text := req.Form["text"]
	reply := spamc.SpamReport{}

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: ReportSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		reply.Code = 404
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("ReportSpamHandler, -->  ", text)
	spamc.Train("bad", text[0])
	reply.Code = 200

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: ReportSpamHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func RevokeSpamHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	text := req.Form["text"]
	reply := spamc.SpamReport{}

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: RevokeSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		reply.Code = 404
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("RevokeSpamHandler, -->  ", text)

	spamc.Untrain("bad",text[0])

	reply.Code = 200


	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: RevokeSpamHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}


func WhitelistSpamHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	text := req.Form["text"]
	reply := spamc.SpamReport{}

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: WhitelistSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		reply.Code = 404
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("WhitelistSpamHandler, -->  ", text)

	spamc.Train("good", text[0])

	reply.Code = 200

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: WhitelistSpamHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}


func SpamHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.
	text := req.Form["text"]
	reply := spamc.SpamReport{}

	//validate inputs
	if len(text) == 0 {
		reply.Code = 404
		errMsg := "Error: SpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("SpamHandler, -->  ", text)

	//classify spam text
	//reply, err := SpamassassinClient.Check(text[0])
	class := spamc.Classify(text[0])
	reply.Code = 200

	if class == "bad" {
		reply.IsSpam = true
	} else {
		reply.IsSpam = false
	}

	log.Println("SpamHandler, reply.Code,  ",reply.Code)
	log.Println("SpamHandler reply.IsSpam, ",reply.IsSpam)

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: SpamHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
