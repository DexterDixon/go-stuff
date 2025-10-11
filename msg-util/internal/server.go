package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println("Thank you come again :)")
		fmt.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func CreateProducers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Producers endpoint")
}

func GetProducers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Producers endpoint")
}

func UpdateProducers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Producers endpoint")
}

func DeleteProducers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Producers endpoint")
}

func CreateConsumers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Consumers endpoint")
}

func GetConsumers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Consumers endpoint")
}

func UpdateConsumers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Consumers endpoint")
}

func AddConsumerTopic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add Consumers endpoint")
}

func DeleteConsumers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Consumers endpoint")
}

func GetHealt(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func WriteJSON(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("json encode error: %v", err)
	}
}
