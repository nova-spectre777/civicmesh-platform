# Architecture

CIVICMESH separates **trust**, **state**, and **transport** so a new Bluetooth, Wi-Fi, LAN, QR, satellite, or internet transport cannot silently redefine what counts as an official alert.

```text
Authority signer ──signed envelope──► Gateway / Relay / File / QR
                                           │
                                           ▼
                                  transport-independent state
                                           │
             ┌─────────────────────────────┼─────────────────────────┐
             ▼                             ▼                         ▼
        Kotlin client                 TypeScript PWA              C++ edge relay
             │                             │                         │
             └─────────────────────────────┴─────────────────────────┘
                                           │
                                    deterministic merge
                                           │
                                    local useful state
```

The Go gateway provides federation and signing primitives. The C++ edge engine holds a compact record set and computes deltas for peers. Kotlin and TypeScript provide offline verification/merge behavior. Python simulates partitions and convergence.
