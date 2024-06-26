package gate

import (
	"fmt"

	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/zdecoder"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/timzzx/zinx-cluster/ddict"
	"github.com/timzzx/zinx-cluster/dmanager"
	"github.com/timzzx/zinx-cluster/dnode"
	"github.com/timzzx/zinx-cluster/zinx-cluster-demo/multiProcess/gate/handlers"
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
	// 前端
	if n.Type == 1 {
		fmt.Println("开启前端拦截器")
		// 启动数据拦截
		s.AddInterceptor(&zdecoder.TLVDecoder{})
		s.AddInterceptor(&dnode.NodeInterceptor{Node: n, GroupName: groupName})
		// 关闭默认的解码器  因为提前解码获取参数，所以后续的解码拦截器要关闭，不然会重复解码报错
		s.SetDecoder(nil)
	}
	// 删除失效链接
	s.SetOnConnStop(func(conn ziface.IConnection) {
		fmt.Println("删除失效连接")
		dmanager.MemberConnManager.Remove(conn)
	})

	// handlers
	s.AddRouterSlices(1, handlers.Login)
	s.AddRouterSlices(2, handlers.Logout)

	return s
}
