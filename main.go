package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	Data []string `json:"data"`
}

type Response struct {
	IsSuccess                bool     `json:"is_success"`
	UserID                   string   `json:"user_id"`
	Email                    string   `json:"email"`
	RollNumber               string   `json:"roll_number"`
	Numbers                  []string `json:"numbers"`
	Alphabets                []string `json:"alphabets"`
	HighestLowercaseAlphabet string   `json:"highest_lowercase_alphabet"`
}

func getHighestLowercase(alphabets []string) string {
	highest := ""
	for _, alpha := range alphabets {
		if alpha == strings.ToLower(alpha) && (highest == "" || alpha > highest) {
			highest = alpha
		}
	}
	return highest
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numbers := []string{}
	alphabets := []string{}
	for _, item := range req.Data {
		if _, err := fmt.Sscanf(item, "%d", new(int)); err == nil {
			numbers = append(numbers, item)
		} else {
			alphabets = append(alphabets, item)
		}
	}

	response := Response{
		IsSuccess:                true,
		UserID:                   "Nitish_21BCE3520",         // Change as needed
		Email:                    "nitish.devwork@gmail.com", // Change as needed
		RollNumber:               "21BCE3520",                // Change as needed
		Numbers:                  numbers,
		Alphabets:                alphabets,
		HighestLowercaseAlphabet: getHighestLowercase(alphabets),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	response := map[string]int{
		"operation_code": 1,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/bfhl", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlePostRequest(w, r)
		} else if r.Method == http.MethodGet {
			handleGetRequest(w, r)
		}
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", withCORS(mux))
}
