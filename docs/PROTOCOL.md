# Protocol v0.1

## Signed alerts

An authority signs the exact UTF-8 bytes in `payload_json` using Ed25519. Relays MUST NOT parse and reserialize the payload before verification because even semantically identical JSON may have different bytes.

`SignedAlertEnvelope` contains `envelope_id`, `key_id`, `algorithm`, `payload_json`, and `signature_b64`.

## Trust bundles

A trust bundle contains authority key records. v0.1 models key status as active/revoked/expired. Production designs need signed trust-bundle updates, key rotation, emergency revocation, expiry, and authority delegation.

## Mesh records

Non-alert state uses stable record IDs, revisions, timestamps, source priority and optional expiry. Merge precedence is:

1. higher revision
2. higher source priority
3. newer update timestamp
4. deterministic payload tie-break

This rule is deliberately simple and must be evolved carefully; arbitrary source priority must not be available to untrusted peers.
