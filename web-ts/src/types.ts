export interface AuthorityKey { key_id:string; authority_id:string; algorithm:"Ed25519"; public_key_b64:string; public_key_spki_b64?:string; status:"active"|"revoked"|"expired" }
export interface SignedAlertEnvelope { envelope_id:string; key_id:string; algorithm:"Ed25519"; payload_json:string; signature_b64:string }
export interface MeshRecord { record_type:string; record_id:string; revision:number; updated_at:string; source_priority:number; body:Record<string,unknown>; expires_at?:string }
export interface ResourceSite { resource_id:string; kind:string; name:string; cell:string; revision:number; updated_at:string; open:boolean; water:boolean; power:boolean; capacity?:number|null; occupied?:number|null }
export interface HelpRequest { request_id:string; kind:string; cell:string; details:string; created_at:string; expires_at:string; status:string; revision:number }
