import {rm,mkdir,cp,readdir,copyFile} from "node:fs/promises";
const build=new URL("../build/",import.meta.url);const dist=new URL("../dist/src/",import.meta.url);const pub=new URL("../public/",import.meta.url);
await rm(build,{recursive:true,force:true});await mkdir(build,{recursive:true});
for(const name of await readdir(dist)){if(name.endsWith(".js"))await copyFile(new URL(name,dist),new URL(name,build));}
await cp(pub,build,{recursive:true});console.log("PWA packaged at web-ts/build");
