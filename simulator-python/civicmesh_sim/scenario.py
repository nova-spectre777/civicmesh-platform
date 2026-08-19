from __future__ import annotations
from .models import Node, Record
from .network import Network

def cyclone_partition_demo() -> dict[str, object]:
    now=1000
    net=Network()
    for nid in ["authority-gateway","school-relay","clinic-relay","phone-a","phone-b"]: net.add_node(Node(nid))
    net.heal_full_mesh()
    alert=Record("alert-1","signed_alert",1,900,100,{"headline":"Cyclone evacuation","verified":True},2000)
    shelter=Record("shelter-1","resource",1,910,50,{"name":"Government High School","capacity":300,"occupied":120,"water":True,"power":True},0)
    net.nodes["authority-gateway"].merge(alert); net.nodes["school-relay"].merge(shelter); net.converge(now)
    net.partition([{"authority-gateway"},{"school-relay","phone-a"},{"clinic-relay","phone-b"}])
    road=Record("road-7","observation",1,1010,10,{"status":"flooded","confidence":0.8},1300)
    help_req=Record("help-1","help_request",1,1015,5,{"kind":"medical","cell":"KDP-17"},1400)
    net.nodes["phone-a"].merge(road); net.nodes["phone-b"].merge(help_req)
    net.converge(1020); partition_counts={nid:len(n.records) for nid,n in net.nodes.items()}
    net.heal_full_mesh(); rounds=net.converge(1030)
    return {"converged":net.is_converged(1030),"rounds_after_reconnect":rounds,"records_after_reconnect":len(net.nodes["phone-a"].active(1030)),"partition_counts":partition_counts}
