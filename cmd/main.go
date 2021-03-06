package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	model "github.com/matthausen/n2yo/src/model"
)

var (
	baseURL = "https://api.n2yo.com/rest/v1/satellite/positions/25544/51.508/-0.125/0/2/&apiKey="
	apiKey  = os.Getenv("API_KEY")
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	satInfo := getSatInfo()
	jsonData, err := json.Marshal(satInfo)
	if err != nil {
		log.Fatalf("Could not marhsla response: %v", err)
	}
	w.Write(jsonData)
}

func getSatInfo() model.Response {
	resp, err := http.Get(baseURL + apiKey)
	if err != nil {
		log.Fatalf("Could not read from endpoint %v", err)
	}

	var satInfo model.Response
	if err := json.NewDecoder(resp.Body).Decode(&satInfo); err != nil {
		log.Println("Could not decode body:", err)
	}

	return satInfo
}

func main() {
	http.HandleFunc("/", index)
	getSatInfo()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
