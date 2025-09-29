// logging.go logs every http req and res that goes through the server 
// it skips health chekpoints and formats logs depending if its in dev or production using logger pkg 


package middleware

import (
	"feast-friends-api/pkg/logger"
	"net/http"
	"time"

	"github.com/google/uuid"

)

type responseWriter struct{ //this allows us to log the statues code 
	http.ResponseWriter

	status int
}
 
// WriteHeader overrides the default WriteHeader method of http.ResponseWriter.
// It captures the status code being written to the response so it can be logged later.
// This is necessary because the standard http.ResponseWriter does not expose the status code after writing.
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code // Store the status code for logging
	rw.ResponseWriter.WriteHeader(code) // Call the underlying ResponseWriter's WriteHeader method
}

func Logs(next http.Handler) http.Handler { //returns a http handler that we can use in request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/health" { //if the request is to /health we dont log it call neextServe straight waway 
			next.ServeHTTP(w, r)
			return
		}

		start := time.Now() //will record the current time 

		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK} // custom wrapper o track http stat code because the standard response writer dosnt expose it after writing 

		reqID := r.Header.Get("X-Request-ID") //checks if the client sent a request id . if not we generate one using uuid
		if reqID == "" {
			reqID = uuid.NewString()
		} 

		rw.Header().Set("X-Request-ID",reqID) // adds the request id to the response header so client can see it 

		duration := time.Since(start).Milliseconds() // tracks how long request took by subing start from current time 

		
		next.ServeHTTP(rw, r)
	
		logger.Log.WithFields(map[string]interface{}{
			"method":    r.Method,
			"url":       r.URL.String(),
			"status":    rw.status,
			"duration":  duration,
			"requestID": reqID,
		}).Info("Http request completed")
	})
}