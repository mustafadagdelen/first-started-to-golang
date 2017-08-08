package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/search", searchFunc).Methods("GET")
	http.ListenAndServe(":8088", router)
}

func searchFunc(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	client, error := elastic.NewClient(elastic.SetURL("http://127.0.0.1:32769"))
	if error != nil {
		log.Fatal(error)
	}


	queryStringQuery := elastic.NewQueryStringQuery("mustafa")
	searchResult, err := client.Search("contentindex").
						Index("contentindex").
						Query(queryStringQuery).
						From(0).
						Size(10).
						Do(ctx)

	if err!=nil{
		panic(err)
	}

	fmt.Fprint(res, searchResult.Hits)
}
