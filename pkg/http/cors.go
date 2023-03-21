package http

import "net/http"

func CorsEnabler(h http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-type,X-CSRF-Token,Authorization")
				return
			} else {
				h.ServeHTTP(w, r)
			}
		})
}

//For CORS error while writing API to front-end web application, we have to deal with cors.
//This example allows us to enable cors
