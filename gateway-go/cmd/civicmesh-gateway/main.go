package main

import (
	"civicmesh/gateway/internal/authority"
	"civicmesh/gateway/internal/httpapi"
	"civicmesh/gateway/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	st := store.New()
	if pub := os.Getenv("CIVICMESH_DEMO_GENERATE_KEY"); pub == "1" {
		signer, err := authority.Generate("demo-key", "demo-authority")
		if err != nil { log.Fatal(err) }
		st.PutKey(signer.PublicRecord())
		log.Printf("generated ephemeral demo authority key id=%s", signer.KeyID)
	}
	addr := os.Getenv("CIVICMESH_ADDR")
	if addr == "" { addr = ":8080" }
	log.Printf("CIVICMESH gateway listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, httpapi.New(st).Handler()))
}
