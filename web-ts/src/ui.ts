import type {MeshRecord} from "./types.js";
export function severityClass(severity:string):string{return ["extreme","severe","moderate","minor","info"].includes(severity)?`severity-${severity}`:"severity-info"}
export function summarize(records:MeshRecord[]):{resources:number;help:number;observations:number}{return{resources:records.filter(r=>r.record_type==="resource").length,help:records.filter(r=>r.record_type==="help_request").length,observations:records.filter(r=>r.record_type==="observation").length}}
