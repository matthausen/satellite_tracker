package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"

	model "github.com/matthausen/n2yo/src/model"
)

var (
	baseURL = "https://api.n2yo.com/rest/v1/satellite/positions/25544/51.508/-0.125/0/2/&apiKey="
	apiKey = os.Getenv("API_KEY")
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello")
}

func getSatInfo() {
	resp, err := http.Get(baseURL+apiKey)
	if err != nil {
		log.Fatalf("Could not read from endpoint %v", err)
	}

	var satInfo model.Response
	if err := json.NewDecoder(resp.Body).Decode(&satInfo); err != nil {
		log.Println("Could not decode body:", err)
	}

	fmt.Println(satInfo)
}

func main() {
	// http.HandleFunc("/", index)
	getSatInfo()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
