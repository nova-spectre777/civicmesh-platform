#include "civicmesh/mesh.hpp"
#include <algorithm>

namespace civicmesh {

bool MeshState::incoming_wins(const Record& current, const Record& incoming) {
    if (incoming.revision != current.revision) return incoming.revision > current.revision;
    if (incoming.source_priority != current.source_priority) return incoming.source_priority > current.source_priority;
    if (incoming.updated_epoch != current.updated_epoch) return incoming.updated_epoch > current.updated_epoch;
    return incoming.payload > current.payload;
}

bool MeshState::upsert(const Record& record) {
    auto it = records_.find(record.id);
    if (it == records_.end()) { records_.emplace(record.id, record); return true; }
    if (incoming_wins(it->second, record)) { it->second = record; return true; }
    return false;
}

std::optional<Record> MeshState::get(const std::string& id) const {
    auto it = records_.find(id); if (it == records_.end()) return std::nullopt; return it->second;
}

std::vector<Record> MeshState::active(std::int64_t now_epoch) const {
    std::vector<Record> out;
    for (const auto& [_, r] : records_) if (r.expires_epoch == 0 || r.expires_epoch > now_epoch) out.push_back(r);
    return out;
}

PeerSummary MeshState::summary(const std::string& node_id) const {
    PeerSummary p{node_id,{}}; for (const auto& [id,r] : records_) p.revisions[id]=r.revision; return p;
}

std::vector<Record> MeshState::delta_for(const PeerSummary& peer, std::int64_t now_epoch) const {
    std::vector<Record> out;
    for (const auto& [id,r] : records_) {
        if (r.expires_epoch != 0 && r.expires_epoch <= now_epoch) continue;
        auto it = peer.revisions.find(id);
        if (it == peer.revisions.end() || r.revision > it->second) out.push_back(r);
    }
    return out;
}

std::size_t MeshState::purge_expired(std::int64_t now_epoch) {
    std::size_t n=0; for (auto it=records_.begin(); it!=records_.end();) { if (it->second.expires_epoch!=0 && it->second.expires_epoch<=now_epoch){it=records_.erase(it);++n;} else ++it; } return n;
}

bool SeenBundles::first_time(const std::string& bundle_id) {
    if (seen_.contains(bundle_id)) return false;
    seen_.insert(bundle_id); order_.push_back(bundle_id);
    while (order_.size() > max_entries_) { seen_.erase(order_.front()); order_.erase(order_.begin()); }
    return true;
}

} // namespace civicmesh
