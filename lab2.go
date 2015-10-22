package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Greet struct {
	Greeting string `json:"Greetings"`
}

type Name struct {
	Name string `json:"name"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var args Name
	err := json.NewDecoder(req.Body).Decode(&args)
	if err != nil {
		panic(err)
	}
	var myGreeting Greet
	myGreeting.Greeting = "Hello," + args.Name + "!"
	err1 := json.NewEncoder(rw).Encode(myGreeting)
	if err1 != nil {
		panic(err1)
	}
}

func main() {
	mux := httprouter.New()
	mux.POST("/hello/", hello)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
