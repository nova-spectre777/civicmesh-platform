declare const self: ServiceWorkerGlobalScope;
const CACHE="civicmesh-v0.1";const ASSETS=["/","/index.html","/app.js","/styles.css","/manifest.webmanifest"];
self.addEventListener("install",e=>{e.waitUntil(caches.open(CACHE).then(c=>c.addAll(ASSETS)))});
self.addEventListener("fetch",e=>{if(e.request.method!=="GET")return;e.respondWith(caches.match(e.request).then(hit=>hit||fetch(e.request).then(r=>{const clone=r.clone();caches.open(CACHE).then(c=>c.put(e.request,clone));return r}).catch(()=>caches.match("/index.html") as Promise<Response>)))});
