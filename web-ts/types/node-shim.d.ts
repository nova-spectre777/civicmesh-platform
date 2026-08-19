declare module "node:test" { const test: (name:string, fn:()=>void|Promise<void>)=>void; export default test; }
declare module "node:assert/strict" { const assert: { equal:(a:unknown,b:unknown,m?:string)=>void; ok:(v:unknown,m?:string)=>void; deepEqual:(a:unknown,b:unknown,m?:string)=>void; rejects:(fn:()=>Promise<unknown>)=>Promise<void> }; export default assert; }
