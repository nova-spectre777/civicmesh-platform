import unittest
from civicmesh_sim.models import Node, Record
from civicmesh_sim.network import Network
from civicmesh_sim.scenario import cyclone_partition_demo

class NetworkTests(unittest.TestCase):
    def test_revision_converges(self):
        net=Network();net.add_node(Node("a"));net.add_node(Node("b"));net.connect("a","b")
        net.nodes["a"].merge(Record("r","resource",1,1,1,{"v":"old"}));net.nodes["b"].merge(Record("r","resource",2,1,1,{"v":"new"}))
        net.converge(10);self.assertEqual(net.nodes["a"].records["r"].body["v"],"new");self.assertTrue(net.is_converged(10))
    def test_expiry(self):
        n=Node("a");n.merge(Record("o","observation",1,1,1,{"x":1},5));self.assertEqual(len(n.active(4)),1);self.assertEqual(len(n.active(5)),0)
    def test_partition_reconnect_scenario(self):
        out=cyclone_partition_demo();self.assertTrue(out["converged"]);self.assertEqual(out["records_after_reconnect"],4)

if __name__=="__main__":unittest.main()
