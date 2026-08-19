package org.civicmesh.offline

import java.time.Instant
import org.civicmesh.model.HelpRequest
import org.civicmesh.model.ResourceSite

class ResourceDirectory(private val resources: MutableMap<String, ResourceSite> = mutableMapOf()) {
    fun upsert(site: ResourceSite) {
        val old = resources[site.resourceId]
        if (old == null || site.revision > old.revision || (site.revision == old.revision && Instant.parse(site.updatedAt).isAfter(Instant.parse(old.updatedAt)))) {
            resources[site.resourceId] = site
        }
    }
    fun openInCell(cell: String): List<ResourceSite> = resources.values.filter { it.cell == cell && it.open }.sortedBy { it.name }
    fun all(): List<ResourceSite> = resources.values.sortedBy { it.resourceId }
}

class HelpQueue(private val items: MutableMap<String, HelpRequest> = mutableMapOf()) {
    fun upsert(request: HelpRequest) {
        val old = items[request.requestId]
        if (old == null || request.revision > old.revision) items[request.requestId] = request
    }
    fun active(now: Instant): List<HelpRequest> = items.values.filter { it.status != "resolved" && Instant.parse(it.expiresAt).isAfter(now) }.sortedBy { it.createdAt }
}
