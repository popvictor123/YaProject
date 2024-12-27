package main

import (
	"encoding/json"
	"strconv"
	"log"
	"net/http"
	"project/internal/calculator"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		s := err.Error()
		if s == "Only numbers and arithmetic operations are allowed" || s == "Division by zero" || s == "Not right parentheses" || s == "Not enough values" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(CalculateResponse{Error: "Expression is not valid. " + s + "."})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(CalculateResponse{Error: "Internal server error"})
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalculateResponse{Result: strconv.FormatFloat(result, 'f', -1, 64)})
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

