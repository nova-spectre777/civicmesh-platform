# CIVICMESH v0.1 Verification

Verified on 2026-08-19 with the repository-wide `./scripts/check_all.sh` workflow.

## Results

- C++20 edge relay: compiled and CTest passed.
- Go gateway: all package tests passed, including Ed25519 tamper rejection and store revision behavior.
- Kotlin/JVM offline client: signature verification/tamper tests and deterministic merge tests passed.
- TypeScript: strict compilation succeeded; 4 Node tests passed, including WebCrypto Ed25519 verification and tamper rejection.
- Python simulator: 3 tests passed, including partition/reconnect convergence.
- Cross-language interoperability: a Go-generated signed alert verified in the TypeScript client and a modified payload was rejected.
- PWA packaging: static offline web build generated successfully.

## Resilience scenario

The built-in cyclone partition scenario creates five nodes, partitions them into three disconnected groups, creates a road-flood observation and medical help request while isolated, reconnects the network, and converges all active nodes to four active records in two gossip rounds.

## Scope note

This is an engineering foundation, not a certified emergency warning system. Bluetooth/Wi-Fi Direct transports, production key rotation/revocation, offline maps, encrypted sensitive help channels, and field validation remain roadmap work.
