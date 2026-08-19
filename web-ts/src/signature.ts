import type {AuthorityKey,SignedAlertEnvelope} from "./types.js";
function b64(value:string):Uint8Array { return Uint8Array.from(atob(value), c=>c.charCodeAt(0)); }
export async function verifyAlert(env:SignedAlertEnvelope,key:AuthorityKey):Promise<boolean>{
  if(key.status!=="active"||key.key_id!==env.key_id||key.algorithm!=="Ed25519"||env.algorithm!=="Ed25519") return false;
  const publicKey=await crypto.subtle.importKey("raw",b64(key.public_key_b64),{name:"Ed25519"},false,["verify"]);
  return crypto.subtle.verify({name:"Ed25519"},publicKey,b64(env.signature_b64),new TextEncoder().encode(env.payload_json));
}
