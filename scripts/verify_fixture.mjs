import {readFile} from "node:fs/promises";
import {verifyAlert} from "../web-ts/dist/src/signature.js";
const file=process.argv[2];if(!file)throw new Error("fixture path required");
const data=JSON.parse(await readFile(file,"utf8"));
if(!(await verifyAlert(data.envelope,data.key))) throw new Error("Go-signed fixture failed TypeScript verification");
if(await verifyAlert({...data.envelope,payload_json:data.envelope.payload_json+" "},data.key)) throw new Error("tampered cross-language fixture verified");
console.log("Go -> TypeScript Ed25519 interoperability passed");
