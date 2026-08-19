package org.civicmesh.sync

import java.time.Instant

interface Revisioned { val recordId: String; val revision: Int; val updatedAt: String }
data class MeshRecord(override val recordId:String, override val revision:Int, override val updatedAt:String, val type:String, val payload:String, val sourcePriority:Int=0):Revisioned
object MergePolicy {
    fun winner(a: MeshRecord,b: MeshRecord):MeshRecord {
        require(a.recordId==b.recordId){"cannot merge different records"}
        if(a.revision!=b.revision)return if(a.revision>b.revision)a else b
        if(a.sourcePriority!=b.sourcePriority)return if(a.sourcePriority>b.sourcePriority)a else b
        val ta=Instant.parse(a.updatedAt);val tb=Instant.parse(b.updatedAt);if(ta!=tb)return if(ta.isAfter(tb))a else b
        return if(a.payload>=b.payload)a else b
    }
    fun merge(local:Map<String,MeshRecord>,incoming:Iterable<MeshRecord>):Map<String,MeshRecord>{val out=local.toMutableMap();for(record in incoming){val old=out[record.recordId];out[record.recordId]=if(old==null)record else winner(old,record)};return out.toMap()}
}
