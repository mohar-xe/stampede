// stampede control plane entrypoint.
//
// TODO(you): wire the pieces together in this order:
//  1. a Store (in-memory first, Postgres/Neon in Phase 2)
//  2. the HTTP API (chi or stdlib mux — stdlib is enough)
//  3. the job queue (Phase 3) that dispatches harness runs
package main

import (
	"log"
	"os"

	"github.com/you/stampede/control/internal/api"
)

func main() {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	// TODO(you): store := session.NewMemoryStore()  // swap for Postgres in Phase 2
	srv := api.New(/* store */)

	log.Printf("stampede control plane listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
