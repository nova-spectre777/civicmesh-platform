from __future__ import annotations
from dataclasses import dataclass, field
from typing import Any

@dataclass(frozen=True)
class Record:
    record_id: str
    record_type: str
    revision: int
    updated_at: int
    source_priority: int
    body: dict[str, Any]
    expires_at: int = 0

    def wins_over(self, other: "Record") -> bool:
        if self.record_id != other.record_id:
            raise ValueError("cannot compare different records")
        return (self.revision, self.source_priority, self.updated_at, repr(sorted(self.body.items()))) > (other.revision, other.source_priority, other.updated_at, repr(sorted(other.body.items())))

@dataclass
class Node:
    node_id: str
    records: dict[str, Record] = field(default_factory=dict)
    online: bool = True

    def merge(self, record: Record) -> bool:
        old = self.records.get(record.record_id)
        if old is None or record.wins_over(old):
            self.records[record.record_id] = record
            return True
        return False

    def active(self, now: int) -> dict[str, Record]:
        return {k:v for k,v in self.records.items() if v.expires_at == 0 or v.expires_at > now}
