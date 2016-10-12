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

	// pretend the path is currently just "/"
	bucket := ""

	if bucket == "" {
		err := h.db.View(func(tx *bolt.Tx) error {
			c := tx.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Fprintf(w, "<li>%s : %s</li>\n", k, v)
			}
			return nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Fprintf(w, "Deeper buckets are not yet implemented (%s)\n", r.URL.Path)
	}
}
