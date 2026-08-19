package store

import (
	"errors"
	"sort"
	"sync"
	"civicmesh/gateway/internal/domain"
)

type Store struct{mu sync.RWMutex;alerts map[string]domain.SignedAlertEnvelope;resources map[string]domain.ResourceSite;requests map[string]domain.HelpRequest;observations map[string]domain.Observation;keys map[string]domain.AuthorityKey}
func New()*Store{return &Store{alerts:map[string]domain.SignedAlertEnvelope{},resources:map[string]domain.ResourceSite{},requests:map[string]domain.HelpRequest{},observations:map[string]domain.Observation{},keys:map[string]domain.AuthorityKey{}}}
func(s *Store)PutKey(k domain.AuthorityKey){s.mu.Lock();defer s.mu.Unlock();s.keys[k.KeyID]=k}
func(s *Store)Key(id string)(domain.AuthorityKey,bool){s.mu.RLock();defer s.mu.RUnlock();k,ok:=s.keys[id];return k,ok}
func(s *Store)PutAlert(a domain.SignedAlertEnvelope){s.mu.Lock();defer s.mu.Unlock();s.alerts[a.EnvelopeID]=a}
func(s *Store)Alerts()[]domain.SignedAlertEnvelope{s.mu.RLock();defer s.mu.RUnlock();out:=make([]domain.SignedAlertEnvelope,0,len(s.alerts));for _,v:=range s.alerts{out=append(out,v)};sort.Slice(out,func(i,j int)bool{return out[i].EnvelopeID<out[j].EnvelopeID});return out}
func newer(revA int,timeA string,revB int,timeB string)bool{if revA!=revB{return revA>revB};return domain.ParseTime(timeA).After(domain.ParseTime(timeB))}
func(s *Store)UpsertResource(v domain.ResourceSite)bool{s.mu.Lock();defer s.mu.Unlock();old,ok:=s.resources[v.ResourceID];if !ok||newer(v.Revision,v.UpdatedAt,old.Revision,old.UpdatedAt){s.resources[v.ResourceID]=v;return true};return false}
func(s *Store)UpsertRequest(v domain.HelpRequest)bool{s.mu.Lock();defer s.mu.Unlock();old,ok:=s.requests[v.RequestID];if !ok||newer(v.Revision,v.CreatedAt,old.Revision,old.CreatedAt){s.requests[v.RequestID]=v;return true};return false}
func(s *Store)UpsertObservation(v domain.Observation)bool{s.mu.Lock();defer s.mu.Unlock();old,ok:=s.observations[v.ObservationID];if !ok||newer(v.Revision,v.CreatedAt,old.Revision,old.CreatedAt){s.observations[v.ObservationID]=v;return true};return false}
func(s *Store)Resources()[]domain.ResourceSite{s.mu.RLock();defer s.mu.RUnlock();out:=make([]domain.ResourceSite,0,len(s.resources));for _,v:=range s.resources{out=append(out,v)};sort.Slice(out,func(i,j int)bool{return out[i].ResourceID<out[j].ResourceID});return out}
func(s *Store)Requests()[]domain.HelpRequest{s.mu.RLock();defer s.mu.RUnlock();out:=make([]domain.HelpRequest,0,len(s.requests));for _,v:=range s.requests{out=append(out,v)};sort.Slice(out,func(i,j int)bool{return out[i].RequestID<out[j].RequestID});return out}
func(s *Store)Observations()[]domain.Observation{s.mu.RLock();defer s.mu.RUnlock();out:=make([]domain.Observation,0,len(s.observations));for _,v:=range s.observations{out=append(out,v)};sort.Slice(out,func(i,j int)bool{return out[i].ObservationID<out[j].ObservationID});return out}
func(s *Store)Resource(id string)(domain.ResourceSite,error){s.mu.RLock();defer s.mu.RUnlock();v,ok:=s.resources[id];if !ok{return domain.ResourceSite{},errors.New("not found")};return v,nil}
func(s *Store)Request(id string)(domain.HelpRequest,error){s.mu.RLock();defer s.mu.RUnlock();v,ok:=s.requests[id];if !ok{return domain.HelpRequest{},errors.New("not found")};return v,nil}
func(s *Store)Observation(id string)(domain.Observation,error){s.mu.RLock();defer s.mu.RUnlock();v,ok:=s.observations[id];if !ok{return domain.Observation{},errors.New("not found")};return v,nil}
