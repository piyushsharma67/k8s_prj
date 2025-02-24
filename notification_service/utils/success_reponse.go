package utils

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse sends a JSON response with http.StatusOK by default
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	code := http.StatusOK // Default status code

	// If a custom status code is provided, use it
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Convert data to JSON
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    data,
	})
}
