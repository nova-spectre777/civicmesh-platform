package org.civicmesh.model

data class AuthorityKey(
    val keyId: String,
    val authorityId: String,
    val algorithm: String,
    val publicKeyB64: String,
    val publicKeySpkiB64: String,
    val status: String,
)

data class SignedAlertEnvelope(
    val envelopeId: String,
    val keyId: String,
    val algorithm: String,
    val payloadJson: String,
    val signatureB64: String,
)

data class ResourceSite(
    val resourceId: String,
    val kind: String,
    val name: String,
    val cell: String,
    val revision: Int,
    val updatedAt: String,
    val open: Boolean = true,
    val water: Boolean = false,
    val power: Boolean = false,
    val capacity: Int? = null,
    val occupied: Int? = null,
)

data class HelpRequest(
    val requestId: String,
    val kind: String,
    val cell: String,
    val details: String,
    val createdAt: String,
    val expiresAt: String,
    val status: String,
    val revision: Int,
)

data class Observation(
    val observationId: String,
    val kind: String,
    val cell: String,
    val value: String,
    val createdAt: String,
    val expiresAt: String,
    val revision: Int,
    val confidence: Double,
    val sourceClass: String,
)
