package logger

import (
	"log"
	"net/http"
	"time"
)

type SystemLog struct {
	Id      int
	Name    string
	CrateAt *time.Time
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//counter = counter + 1

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
			//counter,
		)
	})
}
