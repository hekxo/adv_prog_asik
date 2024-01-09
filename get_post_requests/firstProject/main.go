package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type Person struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	Address      Address   `json:"address"`
	Contacts     []Contact `json:"contacts"`
	IsStudent    bool      `json:"isStudent"`
	Grades       []int     `json:"grades"`
	RegisteredAt string    `json:"registeredAt"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Contact struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequest(w, r)
	case http.MethodPost:
		handlePostRequest(w, r)
	default:
		handleErrorResponse(w, http.StatusMethodNotAllowed, "Invalid request method")
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	person := Person{
		ID:           123,
		Name:         "John Doe",
		Age:          30,
		Address:      Address{"123 Main Street", "Anytown", "USA"},
		Contacts:     []Contact{{"email", "john.doe@example.com"}, {"phone", "+1 555-1234"}},
		IsStudent:    false,
		Grades:       []int{95, 89, 92},
		RegisteredAt: "2022-03-15T10:30:00Z",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestData Person
	err := decoder.Decode(&requestData)
	if err != nil {
		handleErrorResponse(w, http.StatusBadRequest, "Invalid JSON message")
		return
	}

	response := JsonResponse{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleErrorResponse(w http.ResponseWriter, status int, message string) {
	response := JsonResponse{
		Status:  fmt.Sprintf("%d", status),
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
