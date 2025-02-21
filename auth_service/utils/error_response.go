package utils

import (
    "encoding/json"
    "net/http"
)

// ErrorResponse sends a structured error response in JSON format
func ErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
    // Set the status code
    w.WriteHeader(statusCode)

    // Create a response structure
    response := map[string]string{
        "error": message,
    }

    // Set the Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Write the response body as JSON
    if err := json.NewEncoder(w).Encode(response); err != nil {
        // If encoding fails, return a generic server error
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}
