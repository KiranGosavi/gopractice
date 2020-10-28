package cors

import "net/http"

func MiddlewareHandler(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length,Accept,Accept-Encoding")

		handler.ServeHTTP(
			w,
			r,
		)
	})
}
