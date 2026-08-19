package domain

import "time"

type AlertPayload struct { AlertID string `json:"alert_id"`; AuthorityID string `json:"authority_id"`; Severity string `json:"severity"`; Category string `json:"category"`; Headline string `json:"headline"`; Body string `json:"body"`; IssuedAt string `json:"issued_at"`; ExpiresAt string `json:"expires_at"`; Revision int `json:"revision"`; Areas []string `json:"areas"`; Instructions []string `json:"instructions,omitempty"`; Languages map[string]any `json:"languages,omitempty"` }
type SignedAlertEnvelope struct { EnvelopeID string `json:"envelope_id"`; KeyID string `json:"key_id"`; Algorithm string `json:"algorithm"`; PayloadJSON string `json:"payload_json"`; SignatureB64 string `json:"signature_b64"` }
type AuthorityKey struct { KeyID string `json:"key_id"`; AuthorityID string `json:"authority_id"`; Algorithm string `json:"algorithm"`; PublicKeyB64 string `json:"public_key_b64"`; PublicKeySPKIB64 string `json:"public_key_spki_b64,omitempty"`; Status string `json:"status"` }
type ResourceSite struct { ResourceID string `json:"resource_id"`; Kind string `json:"kind"`; Name string `json:"name"`; Cell string `json:"cell"`; Capacity *int `json:"capacity,omitempty"`; Occupied *int `json:"occupied,omitempty"`; Water bool `json:"water"`; Power bool `json:"power"`; Open bool `json:"open"`; Revision int `json:"revision"`; UpdatedAt string `json:"updated_at"`; Source string `json:"source"` }
type HelpRequest struct { RequestID string `json:"request_id"`; Kind string `json:"kind"`; Cell string `json:"cell"`; Details string `json:"details"`; CreatedAt string `json:"created_at"`; ExpiresAt string `json:"expires_at"`; Status string `json:"status"`; Revision int `json:"revision"`; ContactToken *string `json:"contact_token,omitempty"` }
type Observation struct { ObservationID string `json:"observation_id"`; Kind string `json:"kind"`; Cell string `json:"cell"`; Value string `json:"value"`; CreatedAt string `json:"created_at"`; ExpiresAt string `json:"expires_at"`; Revision int `json:"revision"`; Confidence float64 `json:"confidence"`; SourceClass string `json:"source_class"` }
type Record struct { RecordType string `json:"record_type"`; RecordID string `json:"record_id"`; Revision int `json:"revision"`; UpdatedAt string `json:"updated_at"`; Body map[string]any `json:"body"` }
type SyncBundle struct { BundleID string `json:"bundle_id"`; NodeID string `json:"node_id"`; CreatedAt string `json:"created_at"`; Records []Record `json:"records"` }
func ParseTime(v string) time.Time { t,_:=time.Parse(time.RFC3339,v); return t }
