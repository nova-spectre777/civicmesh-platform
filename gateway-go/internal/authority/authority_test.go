package authority

import "testing"

func TestSignVerifyAndTamper(t *testing.T) {
	s, err := Generate("key-1", "odisha-demo")
	if err != nil { t.Fatal(err) }
	env := s.SignPayload("env-1", `{"alert_id":"a1","headline":"Evacuate"}`)
	if err := Verify(env, s.PublicRecord()); err != nil { t.Fatalf("verify: %v", err) }
	env.PayloadJSON = `{"alert_id":"a1","headline":"Fake"}`
	if err := Verify(env, s.PublicRecord()); err == nil { t.Fatal("tampering should fail") }
}
