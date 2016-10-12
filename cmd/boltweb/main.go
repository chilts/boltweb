package main

import (
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/chilts/boltweb"
)

func main() {
	log.SetFlags(0)
	var (
		addr = flag.String("addr", ":9001", "bind address")
	)
	flag.Parse()

	// Validate parameters.
	var path = flag.Arg(0)
	if path == "" {
		log.Fatal("path required")
	}

	// Open the database.
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Enable logging.
	log.SetFlags(log.LstdFlags)

	// Setup the HTTP handlers.
	http.Handle("/", boltweb.NewHandler(db))

	// Start the HTTP server.
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(*addr, nil))
	}()

	log.Printf("Listening on http://localhost%s\n", *addr)
	wg.Wait()
}
