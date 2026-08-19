#pragma once
#include <cstdint>
#include <map>
#include <optional>
#include <set>
#include <string>
#include <vector>

namespace civicmesh {

struct Record {
    std::string id;
    std::string type;
    std::uint64_t revision{};
    std::int64_t updated_epoch{};
    int source_priority{};
    std::string payload;
    std::int64_t expires_epoch{};
};

struct PeerSummary {
    std::string node_id;
    std::map<std::string, std::uint64_t> revisions;
};

class MeshState {
public:
    bool upsert(const Record& record);
    std::optional<Record> get(const std::string& id) const;
    std::vector<Record> active(std::int64_t now_epoch) const;
    PeerSummary summary(const std::string& node_id) const;
    std::vector<Record> delta_for(const PeerSummary& peer, std::int64_t now_epoch) const;
    std::size_t purge_expired(std::int64_t now_epoch);
    std::size_t size() const noexcept { return records_.size(); }
private:
    static bool incoming_wins(const Record& current, const Record& incoming);
    std::map<std::string, Record> records_;
};

class SeenBundles {
public:
    explicit SeenBundles(std::size_t max_entries = 1024) : max_entries_(max_entries) {}
    bool first_time(const std::string& bundle_id);
    std::size_t size() const noexcept { return order_.size(); }
private:
    std::size_t max_entries_;
    std::vector<std::string> order_;
    std::set<std::string> seen_;
};

} // namespace civicmesh
