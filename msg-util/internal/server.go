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
		fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, time.Since(start))
	})
}

type CreateProducerReq struct {
	Name             string   `json:"name"`             // required
	BootstrapServers []string `json:"bootstrapServers"` // required (e.g. ["broker:9092"])
	Topic            string   `json:"topic"`            // required
	SecretNamespace  string   `json:"secretNamespace"`  // required
}

type CreateConsumerReq struct {
	Name             string   `json:"name"`             // required
	BootstrapServers []string `json:"bootstrapServers"` // required (e.g. ["broker:9092"])
	Topics           []string `json:"topics"`           // required
	GroupID          string   `json:"groupID"`          // required
	SecretNamespace  string   `json:"secretNamespace"`  // required
}

type ServerResp struct {
	Name    string `json:"name"`
	Success bool   `json:"success"`
}

// GET /v1/producers
func GetProducers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Producers endpoint")
}

// POST /v1/consumers
func CreateProducers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Producers endpoint")
}

// PUT /v1/consumers
func UpdateProducers(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Updatating Producer: %s\n", msg)
}

// DELETE /v1/producers?name={producerName}
func DeleteProducers(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Deleting Producer: %s\n", msg)
}

// GET /v1/consumers
func GetConsumers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Consumers endpoint")
}

// POST /v1/consumers
func CreateConsumers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Consumers endpoint")
}

// PUT /v1/consumers?name={consumerName}
func UpdateConsumers(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Updatating Consumer: %s\n", msg)
}

// PATCH /v1/consumers?name={consumerName}
func AddConsumerTopic(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Adding Consumer: %s\n", msg)
}

// DELETE /v1/consumers?name={consumerName}
func DeleteConsumers(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Deleting Consumer: %s\n", msg)
}

// GET /v1/healthz
func GetHealth(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// GET /v1/echo/{msg}
func Echo(w http.ResponseWriter, r *http.Request) {
	msg := r.PathValue("msg")
	WriteJSON(w, http.StatusOK, map[string]any{
		"msg": msg,
	})
}

func WriteJSON(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("json encode error: %v", err)
	}
}
