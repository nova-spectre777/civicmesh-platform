package httpapi

import (
	"bytes"
	"civicmesh/gateway/internal/authority"
	"civicmesh/gateway/internal/store"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestAlertEndpointRejectsTamper(t *testing.T){st:=store.New();signer,_:=authority.Generate("k1","demo");st.PutKey(signer.PublicRecord());srv:=New(st).Handler();env:=signer.SignPayload("e1",`{"alert_id":"a1"}`);env.PayloadJSON=`{"alert_id":"evil"}`;b,_:=json.Marshal(env);req:=httptest.NewRequest("POST","/v1/alerts",bytes.NewReader(b));rr:=httptest.NewRecorder();srv.ServeHTTP(rr,req);if rr.Code!=403{t.Fatalf("got %d",rr.Code)}}
