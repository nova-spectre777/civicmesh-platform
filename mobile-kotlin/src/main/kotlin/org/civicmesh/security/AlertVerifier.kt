package org.civicmesh.security

import java.security.KeyFactory
import java.security.Signature
import java.security.spec.X509EncodedKeySpec
import java.util.Base64
import org.civicmesh.model.AuthorityKey
import org.civicmesh.model.SignedAlertEnvelope

sealed class VerificationResult {
    data class Verified(val authorityId: String, val keyId: String) : VerificationResult()
    data class Rejected(val reason: String) : VerificationResult()
}

object AlertVerifier {
    fun verify(envelope: SignedAlertEnvelope, key: AuthorityKey): VerificationResult {
        if (key.status != "active") return VerificationResult.Rejected("authority key is not active")
        if (envelope.keyId != key.keyId) return VerificationResult.Rejected("key id mismatch")
        if (envelope.algorithm != "Ed25519" || key.algorithm != "Ed25519") return VerificationResult.Rejected("unsupported algorithm")
        return try {
            val encoded = Base64.getDecoder().decode(key.publicKeySpkiB64)
            val publicKey = KeyFactory.getInstance("Ed25519").generatePublic(X509EncodedKeySpec(encoded))
            val signature = Signature.getInstance("Ed25519")
            signature.initVerify(publicKey)
            signature.update(envelope.payloadJson.toByteArray(Charsets.UTF_8))
            if (signature.verify(Base64.getDecoder().decode(envelope.signatureB64))) VerificationResult.Verified(key.authorityId, key.keyId)
            else VerificationResult.Rejected("signature verification failed")
        } catch (e: Exception) { VerificationResult.Rejected("verification error: ${e.message}") }
    }
}
