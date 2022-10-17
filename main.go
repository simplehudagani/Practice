package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func facrorial(w http.ResponseWriter, r *http.Request) {
	var fact int = 4
	var facto int = 1
	// fmt.Printf("Enter the number ")
	// fmt.Scanln(&fact)

	for i := 1; i <= fact; i++ {
		facto = facto * i
		fmt.Println(facto)
	}

	json.NewEncoder(w).Encode(facto)
}

func main() {
	http.HandleFunc("/", facrorial)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
