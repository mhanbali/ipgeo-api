package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oschwald/maxminddb-golang"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Query(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// the database can be found at the following sources:
	// https://github.com/abdullahdevrel/kaggle_data_source
	// https://www.kaggle.com/datasets/ipinfo/ipinfo-country-asn
	db, err := maxminddb.Open("country_asn.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get the ip address from the url param
	ip := net.ParseIP(ps.ByName("name"))

	var record any
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Panic(err)
	}
	
	// convert the map to JSON output
	jsonString, _ := json.Marshal(record)
	if err != nil {
		fmt.Fprint(w, "Error creating the JSON output.", err)
		return
	}
	
	fmt.Fprint(w, string(jsonString))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/query/:name", Query)

	fmt.Println("Running server...")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}