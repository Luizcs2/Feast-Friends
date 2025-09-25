//cors.go this file tells the server which frontend is allowed to make requests to it 
//allows specific http methods,headers and cookies can be send 
//handles preflight options incase the server wants to know what is can request 


package middleware
import (
	"net/http"
	"feast-friends-api/internal/config"
)

// CROS adds CORS headers to HTTP responses.
func CROS(next http.Handler) http.Handler {
	Origin := config.Get().Server.Frontend

	// Return a handler that sets CORS headers
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Use localhost if no origin is set
		if Origin == "" {
			Origin = "https://localhost:3000"
		}

		// Allow requests from the frontend
		w.Header().Set("Access-Control-Allow-Origin", Origin)
		// Allow these HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// Allow these headers from the frontend
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Allow cookies and credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// If it's a preflight request, just return yes or no
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}