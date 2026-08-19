import type {MeshRecord,SignedAlertEnvelope,AuthorityKey} from "./types.js";
export interface StateSnapshot { records:MeshRecord[]; alerts:SignedAlertEnvelope[]; keys:AuthorityKey[]; updatedAt:string }
export class MemoryStore {
  private records=new Map<string,MeshRecord>(); private alerts=new Map<string,SignedAlertEnvelope>(); private keys=new Map<string,AuthorityKey>();
  putKey(k:AuthorityKey){this.keys.set(k.key_id,k)}; key(id:string){return this.keys.get(id)}
  putAlert(a:SignedAlertEnvelope){this.alerts.set(a.envelope_id,a)}
  putRecord(r:MeshRecord){this.records.set(r.record_id,r)}
  snapshot():StateSnapshot{return{records:[...this.records.values()],alerts:[...this.alerts.values()],keys:[...this.keys.values()],updatedAt:new Date().toISOString()}}
}
