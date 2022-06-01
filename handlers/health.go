package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func Health(mux chi.Router) {
	mux.Get("/health", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received health ping")
	})
}
