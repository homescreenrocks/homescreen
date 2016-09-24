package modulemanager

import (
	"log"
	"net/http"
)

func HTTPAddPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		r.URL.Path = prefix + r.URL.Path
		log.Println(r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
