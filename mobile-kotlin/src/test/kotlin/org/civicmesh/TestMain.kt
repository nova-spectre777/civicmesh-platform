package org.civicmesh

import java.security.KeyPairGenerator
import java.security.Signature
import java.util.Base64
import org.civicmesh.model.AuthorityKey
import org.civicmesh.model.SignedAlertEnvelope
import org.civicmesh.security.AlertVerifier
import org.civicmesh.security.VerificationResult
import org.civicmesh.sync.MeshRecord
import org.civicmesh.sync.MergePolicy

private fun assertTrue(value:Boolean,message:String){if(!value)error(message)}
fun main(){
    val kp=KeyPairGenerator.getInstance("Ed25519").generateKeyPair()
    val payload="{\"alert_id\":\"a1\",\"headline\":\"Evacuate\"}"
    val signer=Signature.getInstance("Ed25519");signer.initSign(kp.private);signer.update(payload.toByteArray())
    val env=SignedAlertEnvelope("env1","key1","Ed25519",payload,Base64.getEncoder().encodeToString(signer.sign()))
    val key=AuthorityKey("key1","demo-authority","Ed25519","",Base64.getEncoder().encodeToString(kp.public.encoded),"active")
    assertTrue(AlertVerifier.verify(env,key) is VerificationResult.Verified,"valid signature should verify")
    assertTrue(AlertVerifier.verify(env.copy(payloadJson="tampered"),key) is VerificationResult.Rejected,"tamper must reject")
    val a=MeshRecord("r1",1,"2026-08-19T10:00:00Z","resource","old",1)
    val b=MeshRecord("r1",2,"2026-08-19T09:00:00Z","resource","new",1)
    assertTrue(MergePolicy.winner(a,b).payload=="new","revision must win")
    val c=MeshRecord("r1",2,"2026-08-19T09:00:00Z","resource","official",5)
    assertTrue(MergePolicy.winner(b,c).payload=="official","source priority must break equal revision")
    println("Kotlin CIVICMESH tests passed")
}
