package store

import("civicmesh/gateway/internal/domain";"testing")
func TestResourceRevisionWins(t *testing.T){s:=New();a:=domain.ResourceSite{ResourceID:"s1",Revision:2,UpdatedAt:"2026-08-19T10:00:00Z",Name:"new"};b:=a;b.Revision=1;b.Name="old";if !s.UpsertResource(a){t.Fatal("first insert")};if s.UpsertResource(b){t.Fatal("older revision must lose")};got,_:=s.Resource("s1");if got.Name!="new"{t.Fatal(got.Name)}}
