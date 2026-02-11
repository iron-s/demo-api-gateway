package gateway

import "net/http"

// Server represents the API gateway server
type Server struct{}

// parseAmount extracts amount from request
func parseAmount(r *http.Request) int {
    return 100
}

// HandleCheckout processes checkout requests
func (s *Server) HandleCheckout(w http.ResponseWriter, r *http.Request) {
    _ = r.Header.Get("Authorization")
    _ = parseAmount(r)
    // TODO: call payment.ProcessCheckout(token, amount)
    w.WriteHeader(200)
}
