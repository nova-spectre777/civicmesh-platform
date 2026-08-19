# Contributing to CIVICMESH

CIVICMESH is public-safety infrastructure. Contributions should optimize for clarity, resilience, testability, and failure behavior before feature count.

## High-value areas

- Bluetooth LE / Wi-Fi Direct / LAN transport adapters
- Android offline UX and accessibility
- Ed25519 key rotation/revocation and trust-bundle tooling
- geospatial shelter/resource indexing and offline maps
- anti-spam and abuse controls for community observations
- CRDT/conflict-resolution research
- low-power Raspberry Pi / embedded relay support
- multilingual and text-to-speech interfaces
- field simulation and chaos testing
- emergency-management standards interoperability

## Rules

1. Never commit private authority keys or personal emergency data.
2. Do not label community observations as official alerts.
3. Keep transport adapters separate from signed-message semantics.
4. Add deterministic tests for merge/conflict behavior.
5. Security-sensitive changes require threat-model notes.
6. Destructive admin actions must be explicit and auditable.
7. Large protocol changes should start as an RFC issue.

Run `./scripts/check_all.sh` before opening a PR.
