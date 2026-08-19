package httpapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"civicmesh/gateway/internal/authority"
	"civicmesh/gateway/internal/domain"
	"civicmesh/gateway/internal/store"
	syncer "civicmesh/gateway/internal/sync"
)

type Server struct{ Store *store.Store }
func New(s *store.Store) *Server { return &Server{Store:s} }
func writeJSON(w http.ResponseWriter, code int, v any){ w.Header().Set("Content-Type","application/json"); w.WriteHeader(code); _=json.NewEncoder(w).Encode(v) }
func decode(r *http.Request,v any) error { return json.NewDecoder(io.LimitReader(r.Body,1<<20)).Decode(v) }
func (s *Server) Handler() http.Handler {
	mux:=http.NewServeMux()
	mux.HandleFunc("GET /health",func(w http.ResponseWriter,r *http.Request){writeJSON(w,200,map[string]any{"ok":true,"service":"civicmesh-gateway"})})
	mux.HandleFunc("GET /v1/alerts",func(w http.ResponseWriter,r *http.Request){writeJSON(w,200,s.Store.Alerts())})
	mux.HandleFunc("GET /v1/snapshot",func(w http.ResponseWriter,r *http.Request){writeJSON(w,200,map[string]any{"alerts":s.Store.Alerts(),"resources":s.Store.Resources(),"help_requests":s.Store.Requests(),"observations":s.Store.Observations()})})
	mux.HandleFunc("POST /v1/alerts",func(w http.ResponseWriter,r *http.Request){var env domain.SignedAlertEnvelope;if err:=decode(r,&env);err!=nil{writeJSON(w,400,map[string]string{"error":err.Error()});return};key,ok:=s.Store.Key(env.KeyID);if !ok{writeJSON(w,403,map[string]string{"error":"unknown authority key"});return};if err:=authority.Verify(env,key);err!=nil{writeJSON(w,403,map[string]string{"error":err.Error()});return};s.Store.PutAlert(env);writeJSON(w,201,map[string]any{"accepted":true,"verified":true})})
	mux.HandleFunc("POST /v1/sync",func(w http.ResponseWriter,r *http.Request){var b domain.SyncBundle;if err:=decode(r,&b);err!=nil{writeJSON(w,400,map[string]string{"error":err.Error()});return};writeJSON(w,200,syncer.MergeBundle(s.Store,b))})
	mux.HandleFunc("GET /v1/keys/",func(w http.ResponseWriter,r *http.Request){id:=strings.TrimPrefix(r.URL.Path,"/v1/keys/");k,ok:=s.Store.Key(id);if !ok{writeJSON(w,404,map[string]string{"error":"not found"});return};writeJSON(w,200,k)})
	return mux
}
