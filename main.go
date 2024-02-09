package main

import (
 "log"
 "net/http"
 "encoding/json")

 type note struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Data  string `json:"data"`
}
var notes = []note{
    {
        ID:    "1",
        Title: "Grocery List",
        Data:  "Milk, Eggs, Bread, Butter",
    },
    {
        ID:    "2",
        Title: "Todo for Work",
        Data:  "Finish report, Call with client, Email feedback",
    },
    {
        ID:    "3",
        Title: "Birthday Gift Ideas",
        Data:  "Books, Perfume, Smartwatch, Gift Cards",
    },
    {
        ID:    "4",
        Title: "Learning Goals",
        Data:  "Learn Go, Study React, Practice DSA",
    },
}

 func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET 	", getNotes)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	   }
	   log.Println("Listening...")
	   server.ListenAndServe()
 }
 func getNotes(w http                                       .ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    jsonNotes, err := json.Marshal(notes)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    w.Write(jsonNotes)
}