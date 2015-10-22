package main

import (
	"encoding/json"
	"fmt"
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
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		panic(err)
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
