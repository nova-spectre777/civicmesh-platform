import test from "node:test";import assert from "node:assert/strict";import {winner,activeRecords} from "../src/merge.js";
const base={record_type:"resource",record_id:"r1",revision:1,updated_at:"2026-08-19T10:00:00Z",source_priority:1,body:{name:"old"}};
test("higher revision wins",()=>{assert.equal(winner(base,{...base,revision:2,body:{name:"new"}}).revision,2)});
test("source priority wins equal revision",()=>{assert.equal(winner({...base,revision:2},{...base,revision:2,source_priority:9}).source_priority,9)});
test("expired records filtered",()=>{assert.equal(activeRecords([{...base,expires_at:"2026-08-19T09:00:00Z"}],new Date("2026-08-19T10:00:00Z")).length,0)});
