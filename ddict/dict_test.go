package ddict

import (
	"testing"
)

func TestAddNodeInfo(t *testing.T) {
	// 检查node是否添加成功
	d := NewDict()
	d.AddNodeInfo("test", &NodeInfo{
		ID:   1,
		Name: "test",
		Type: 1,
		IP:   "127.0.0.1",
		Port: 10000,
	})
	if len(d.NodeList()["test"]) != 1 {
		t.Fatal("node添加失败")
	}

	// 检查node是否重复添加
	d.AddNodeInfo("test", &NodeInfo{
		ID:   1,
		Name: "test",
		Type: 1,
		IP:   "127.0.0.1",
		Port: 10000,
	})
	if len(d.NodeList()["test"]) == 2 {
		t.Fatal("node重复添加")
	}
}

func TestAddRouteInfo(t *testing.T) {
	// 检查route是否添加成功
	d := NewDict()
	d.AddRouteInfo("test", &RouteInfo{
		MsgID: 1,
		Name:  "test",
	})
	if len(d.RouteList()["test"]) != 1 {
		t.Fatal("route添加失败")
	}

	// 检查route是否重复添加
	d.AddRouteInfo("test", &RouteInfo{
		MsgID: 1,
		Name:  "test",
	})
	if len(d.RouteList()["test"]) == 2 {
		t.Fatal("route重复添加")
	}
}

func TestGetRouteDicts(t *testing.T) {
	d := NewDict()
	d.AddRouteInfo("test", &RouteInfo{
		MsgID: 1,
		Name:  "test",
	})
	t.Log(d.GetRouteDicts())
	if d.GetRouteDicts()[1].MsgID != 1 {
		t.Fatal("路由字典获取失败")
	}
}
