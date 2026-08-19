#include "civicmesh/mesh.hpp"
#include <iostream>
int main(){
    civicmesh::MeshState state;
    state.upsert({"shelter-1","resource",1,1,2,"{\"open\":true}",0});
    auto summary=state.summary("edge-demo");
    std::cout << "CIVICMESH edge node records=" << summary.revisions.size() << "\n";
}
