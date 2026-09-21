// File: main.go
// Purpose: entry point of the web server.
// Receives: the optional -addr flag from the operating system.
// Previous stage: the operating system.
// Next stage: internal/web (Serve).
// Restrictions: no calculation, formatting, HTML or validation; it only starts the server.

package main

import (
	"flag"
	"log"

	"gauss/internal/web"
)

func main() {
	address := flag.String("addr", ":8080", "listening address")
	flag.Parse()
	log.Printf("Gauss web listening on %s", *address)
	log.Fatal(web.Serve(*address))
}
