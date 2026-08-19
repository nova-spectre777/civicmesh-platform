#include "civicmesh/mesh.hpp"
#include <cassert>
#include <iostream>
using namespace civicmesh;
int main(){
    MeshState a,b;
    assert(a.upsert({"r1","resource",1,100,1,"old",0}));
    assert(a.upsert({"r1","resource",2,90,1,"new",0}));
    assert(!a.upsert({"r1","resource",1,200,9,"stale",0}));
    assert(a.get("r1")->payload=="new");
    b.upsert({"r2","observation",1,100,0,"road flooded",150});
    auto delta=b.delta_for(a.summary("a"),120); assert(delta.size()==1); for(auto &r:delta)a.upsert(r);
    assert(a.size()==2); assert(a.purge_expired(151)==1); assert(a.size()==1);
    SeenBundles seen(2); assert(seen.first_time("b1")); assert(!seen.first_time("b1")); assert(seen.first_time("b2")); assert(seen.first_time("b3")); assert(seen.first_time("b1"));
    std::cout << "C++ CIVICMESH edge tests passed\n";
}
