package main

import (
	"civicmesh/gateway/internal/authority"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	s, err := authority.Generate("cross-key", "cross-authority")
	if err != nil { log.Fatal(err) }
	payload := `{"alert_id":"cross-1","authority_id":"cross-authority","severity":"extreme","category":"cyclone","headline":"Move to shelter","body":"Demonstration fixture","issued_at":"2026-08-19T16:00:00Z","expires_at":"2026-08-20T16:00:00Z","revision":1,"areas":["DEMO-1"]}`
	out := map[string]any{"key": s.PublicRecord(), "envelope": s.SignPayload("cross-envelope", payload)}
	raw, _ := json.Marshal(out)
	fmt.Println(string(raw))
}
