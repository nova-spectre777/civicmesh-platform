<div align="center">

# CIVICMESH
### Offline-First Public Resilience Network

**Signed emergency alerts, local mesh synchronization, shelter/resource discovery, and help coordination that can keep working when normal connectivity degrades.**

</div>

CIVICMESH is an experimental open-source public-good platform for resilient local communication during floods, cyclones, earthquakes, fires, blackouts, and other emergencies. The v0.1 foundation focuses on **trust, offline operation, inspectable state, privacy, and interoperability** rather than pretending a new app can replace official emergency systems.

> **Status:** v0.1 engineering foundation. It is not a certified emergency-alerting system and must not be presented as one. Real deployments require local authorities, emergency-management review, accessibility testing, legal review, security audits, and field exercises.

## Core capabilities

- **Signed authority alerts** using Ed25519 and key IDs.
- **Offline verification** in Go, Kotlin/JVM, and TypeScript/WebCrypto.
- **Local mesh state** for alerts, shelters, resources, road reports, and help requests.
- **Deterministic conflict resolution** using revision, authority priority, timestamp, and stable IDs.
- **Gateway federation model** for synchronizing community nodes when connectivity returns.
- **Privacy-aware help requests** with expiry and coarse location support.
- **Offline PWA foundation** with a service worker and local state store.
- **Edge relay engine** in C++20 for lightweight community nodes.
- **Network simulator** in Python for partition/reconnection scenarios.

## Polyglot architecture

| Layer | Language | Responsibility |
|---|---|---|
| `gateway-go/` | Go | authority signing, API, federation store, sync/merge |
| `mobile-kotlin/` | Kotlin/JVM | offline client domain, alert verification, local merge |
| `edge-cpp/` | C++20 | lightweight relay/node state and gossip planning |
| `web-ts/` | TypeScript | offline PWA logic, WebCrypto verification, operator UI |
| `simulator-python/` | Python | partition, delay, reconnection and convergence simulation |
| `contracts/` | JSON Schema | language-neutral message contracts |

## Trust model

```text
Emergency authority
      │ Ed25519 private key
      ▼
Signed alert envelope
      │
      ├── Internet gateway
      ├── Community relay
      ├── Bluetooth/Wi-Fi bridge (future transport adapter)
      └── QR/file transfer
              │
              ▼
Client verifies with pinned/trusted public key
              │
              ▼
        VERIFIED / UNTRUSTED
```

No client should display an alert as official merely because it came from the CIVICMESH network. Verification is based on cryptographic signatures and an explicitly trusted authority-key registry.

## Quick verification

```bash
./scripts/check_all.sh
```

The script compiles/tests C++20, Go, Kotlin/JVM, TypeScript and Python components without requiring paid services.

## Example workflow

```text
1. Authority signs an evacuation alert.
2. Gateway accepts and verifies the envelope.
3. Community relays synchronize the signed record.
4. Phones verify the signature offline.
5. A road becomes flooded; a local report is created with limited lifetime.
6. A shelter updates capacity and water/power availability.
7. Network partitions.
8. Nearby nodes continue exchanging state locally.
9. Connectivity returns.
10. Federation merges converge deterministically and retain provenance.
```

## Safety boundaries

- No anonymous message becomes an “official alert”.
- No remote command execution exists in the core.
- Precise location is optional; coarse grid/cell location is supported.
- Help requests expire and can be redacted from normal public views.
- Authority private keys never belong in the repository or gateway database.
- Transport adapters must preserve the signed envelope unchanged.
- Local observations are clearly distinct from authority messages.

## Repository layout

```text
contracts/          protocol schemas
gateway-go/         gateway + authority + sync server
mobile-kotlin/      offline client domain layer
edge-cpp/           lightweight relay engine
web-ts/             PWA/domain + WebCrypto verifier
simulator-python/   disaster-network simulator
examples/           fixtures and demo data
docs/               architecture, privacy, threat model, roadmap
scripts/            full verification + demo helpers
```

## Contributing

See [`CONTRIBUTING.md`](CONTRIBUTING.md). CIVICMESH needs networking, cryptography, Android, accessibility, mapping, distributed-systems, emergency-management, security, UX, and field-testing contributors.
