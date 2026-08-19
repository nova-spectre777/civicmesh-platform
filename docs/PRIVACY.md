# Privacy

CIVICMESH should collect less data than conventional cloud-first emergency applications.

- Default public records use coarse `cell` identifiers rather than precise GPS coordinates.
- Help requests can carry a short-lived opaque contact token instead of a phone number.
- Help requests and observations have explicit expiry.
- Relays should avoid retaining historical sensitive records unless required by an authorized operator policy.
- Analytics must be aggregate-first and opt-in where individual linkage is possible.
- A future encrypted help channel should separate public routing metadata from private contact details.
