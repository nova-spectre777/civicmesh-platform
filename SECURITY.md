# Security

CIVICMESH handles safety-critical information. Treat security reports seriously and avoid publishing weaponized bypasses before maintainers have had an opportunity to fix them.

The v0.1 security boundary is intentionally conservative: signed authority alerts, explicit trusted-key registries, immutable signature payloads, no remote code execution, no bundled private keys, short-lived local observations, and planning-only transport adapters.

A real deployment requires an independent cryptographic review, key-rotation procedures, compromised-key revocation, abuse monitoring, accessibility review, disaster-response exercises, and a documented operating authority.
