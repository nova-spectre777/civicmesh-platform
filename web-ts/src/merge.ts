import type {MeshRecord} from "./types.js";
export function winner(a:MeshRecord,b:MeshRecord):MeshRecord{
  if(a.record_id!==b.record_id) throw new Error("record id mismatch");
  if(a.revision!==b.revision) return a.revision>b.revision?a:b;
  if(a.source_priority!==b.source_priority) return a.source_priority>b.source_priority?a:b;
  const ta=Date.parse(a.updated_at),tb=Date.parse(b.updated_at); if(ta!==tb)return ta>tb?a:b;
  return JSON.stringify(a.body)>=JSON.stringify(b.body)?a:b;
}
export function mergeRecords(local:Map<string,MeshRecord>,incoming:Iterable<MeshRecord>):Map<string,MeshRecord>{const out=new Map(local);for(const r of incoming){const old=out.get(r.record_id);out.set(r.record_id,old?winner(old,r):r)}return out}
export function activeRecords(records:Iterable<MeshRecord>,now=new Date()):MeshRecord[]{return [...records].filter(r=>!r.expires_at||Date.parse(r.expires_at)>now.getTime()).sort((a,b)=>a.record_id.localeCompare(b.record_id))}
