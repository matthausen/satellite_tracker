package server

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/matthausen/n2yo/models"
)

var (
	baseURL = "https://api.n2yo.com/rest/v1/satellite/positions/25544/51.508/-0.125/0/2/&apiKey="
	apiKey  = os.Getenv("API_KEY")
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	satData := getSatInfo()

	tmpl, _ := template.ParseFiles("../templates/index.html",
		"../templates/map.html")
	err := tmpl.Execute(w, satData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func getSatInfo() models.Response {
	resp, err := http.Get(baseURL + apiKey)
	if err != nil {
		log.Fatalf("Could not read from endpoint %v", err)
	}

	var satInfo models.Response
	if err := json.NewDecoder(resp.Body).Decode(&satInfo); err != nil {
		log.Println("Could not decode body:", err)
	}

	return satInfo
}

func GracefullyShutDown(ctx context.Context) (err error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(index))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen:%s\n", err)
		}
	}()

	log.Printf("Server started on port 8080")

	<-ctx.Done()

	log.Printf("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server shutdown failed:%+s", err)
	}

	log.Printf("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
