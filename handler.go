package boltweb

import (
	"fmt"
	"net/http"

	"github.com/boltdb/bolt"
)

// NewHandler returns a new root HTTP handler.
func NewHandler(db *bolt.DB) http.Handler {
	h := &handler{db}

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.index)

	return mux
}

type handler struct {
	db *bolt.DB
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "path = %s\n", r.URL.Path)
}
