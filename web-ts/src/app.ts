import {MemoryStore} from "./store.js";
import {summarize} from "./ui.js";
const store=new MemoryStore();
const status=document.querySelector<HTMLDivElement>("#status");
const summary=document.querySelector<HTMLPreElement>("#summary");
function render(){const s=store.snapshot();if(status)status.textContent=navigator.onLine?"Internet available — local mode still active":"OFFLINE — local resilience mode";if(summary)summary.textContent=JSON.stringify(summarize(s.records),null,2)}
window.addEventListener("online",render);window.addEventListener("offline",render);render();
if("serviceWorker" in navigator) navigator.serviceWorker.register("/service-worker.js").catch(()=>{});
