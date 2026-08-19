# Data Model

CIVICMESH distinguishes five major record families:

- **Authority alerts**: cryptographically signed, immutable envelope payload.
- **Resource sites**: shelters, clinics, water, food, power, transport.
- **Observations**: short-lived local reports such as road flooding or network status.
- **Help requests**: expiring requests for medical, rescue, water, food, transport or shelter assistance.
- **Trust keys**: authority public keys and status.

Stable IDs and revisions allow disconnected nodes to converge without requiring a single central database during the outage.
