package im

import (
	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/timzzx/zinx-cluster/ddict"
	"github.com/timzzx/zinx-cluster/zinx-cluster-demo/multiProcess/im/handlers"
)

func App(n *ddict.NodeInfo, groupName ddict.GroupName) ziface.IServer {
	// 设置zinx配置
	config := &zconf.Config{
		Name:    n.Name,
		Host:    n.IP,
		TCPPort: n.Port,
		WsPort:  n.Port + 1,
		// Mode:             "websocket",
		RouterSlicesMode:  true,
		LogIsolationLevel: 2,
	}
	s := znet.NewUserConfServer(config)

	s.AddRouterSlices(3, handlers.Send)
	s.AddRouterSlices(4, handlers.FullSend)
	return s
}
