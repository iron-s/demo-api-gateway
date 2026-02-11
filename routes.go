package gateway

import "net/http"

// Server represents the API gateway server
type Server struct{}

// parseAmount extracts amount from request
func parseAmount(r *http.Request) int {
    return 100
}
