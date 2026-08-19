# Threat Model

Major risks include fake evacuation alerts, stolen authority keys, malicious community reports, replayed old alerts, spammed help requests, location/privacy leakage, Sybil nodes, partition manipulation, compromised relay devices, poisoned resource-capacity reports, denial of service, and unsafe UI wording.

v0.1 mitigations include Ed25519 signatures for official alerts, explicit key status, expiration fields, message-size limits, deterministic merges, no remote execution, separation of official and community state, and coarse location cells.

Not yet solved: signed trust-bundle distribution, compromised-key revocation propagation, Sybil resistance, authenticated operator roles, encrypted sensitive help-request transport, metadata privacy, radio-layer abuse, and formal protocol review.
