package main

import (
	
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	)
	import u "urlshort/utility"



func main(){
	router:=mux.NewRouter()
	const port string =":1235"
	router.HandleFunc("/",func(resp http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(resp,"UPP and running")

	})

	router.HandleFunc("/create/",u.CreateEndPoint)//here html is name of string key that will contain value passed in request the value can be anything with or without quotes
	router.HandleFunc("/expand/",u.ExpandEndPoint).Methods("GET")
 	router.HandleFunc("/{id}",u.RootEndPoint).Methods("GET")

   log.Println("server listening on port",port)
   log.Fatalln(http.ListenAndServe(port,router))

}