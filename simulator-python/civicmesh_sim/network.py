from __future__ import annotations
from dataclasses import dataclass, field
from .models import Node, Record

@dataclass
class Network:
    nodes: dict[str, Node] = field(default_factory=dict)
    links: set[tuple[str, str]] = field(default_factory=set)
    def add_node(self,node:Node)->None:self.nodes[node.node_id]=node
    def connect(self,a:str,b:str)->None:
        if a!=b:self.links.add(tuple(sorted((a,b))))
    def disconnect(self,a:str,b:str)->None:self.links.discard(tuple(sorted((a,b))))
    def partition(self,groups:list[set[str]])->None:
        self.links.clear()
        for group in groups:
            members=sorted(group)
            for i,a in enumerate(members):
                for b in members[i+1:]:self.connect(a,b)
    def heal_full_mesh(self)->None:
        ids=sorted(self.nodes);self.links={tuple(sorted((a,b))) for i,a in enumerate(ids) for b in ids[i+1:]}
    def gossip_round(self,now:int)->int:
        changed=0;snapshots={nid:list(node.active(now).values()) for nid,node in self.nodes.items() if node.online}
        for a,b in sorted(self.links):
            na,nb=self.nodes[a],self.nodes[b]
            if not na.online or not nb.online:continue
            for r in snapshots.get(a,[]):changed+=int(nb.merge(r))
            for r in snapshots.get(b,[]):changed+=int(na.merge(r))
        return changed
    def converge(self,now:int,max_rounds:int=100)->int:
        for round_no in range(1,max_rounds+1):
            if self.gossip_round(now)==0:return round_no
        raise RuntimeError("network did not converge")
    def is_converged(self,now:int)->bool:
        active=[n.active(now) for n in self.nodes.values() if n.online]
        return not active or all(x==active[0] for x in active[1:])
