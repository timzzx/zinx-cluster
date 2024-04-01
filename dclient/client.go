package dclient

import (
	"fmt"
	"math/rand"

	"github.com/aceld/zinx/zdecoder"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/timzzx/zinx-cluster/dconf"
	"github.com/timzzx/zinx-cluster/ddict"
	"github.com/timzzx/zinx-cluster/dmanager"
	"github.com/timzzx/zinx-cluster/dmessage"
)

type ClientInterceptor struct {
}

func (m *ClientInterceptor) Intercept(chain ziface.IChain) ziface.IcResp {
	request := chain.Request() //从责任链中获取当前拦截器的输入数据
	// 这一层是自定义拦截器处理逻辑，这里只是简单打印输入
	iRequest := request.(ziface.IRequest) //注意：由于Zinx的Request类型，这里需要做一下断言转换

	deData, err := dmessage.Decode(iRequest.GetData())
	if err != nil {
		iRequest.Abort()
		return chain.Proceed(chain.Request()) //进入并执行下一个拦截器
	}
	if deData.Type == 1 {
		// 内部消息转发
		// 获取连接
		client := Client(deData.NodeGroup, deData.NodeID)
		err = client.SendMsg(iRequest.GetMsgID(), iRequest.GetData())
		if err != nil {
			fmt.Println("发送数据错误")
			iRequest.Abort()
		}
	} else {
		//发送给玩家
		c, err := dmanager.MemberConnManager.Get(deData.ConnID)
		if err != nil {
			fmt.Println("玩家链接获取失败：", err.Error())
		}
		err = c.SendMsg(iRequest.GetMsgID(), deData.Data)

		if err != nil {
			fmt.Println("发送数据错误")
			iRequest.Abort()
		}
	}
	return chain.Proceed(chain.Request()) //进入并执行下一个拦截器

}

// client
func Client(group ddict.GroupName, NodeId int) ziface.IConnection {
	// 获取随机配置 TODO：可以扩展，不同的负载均衡的方式
	l := dconf.Dicts.NodeList()[group]
	n := rand.Intn(len(l))
	var c *ddict.NodeInfo
	for _, v := range l {
		if v.ID == NodeId {
			c = v
		}
	}
	// 获取配置
	if c == nil {
		c = l[n]
	}

	conn, _ := dmanager.ClientConnManager.Get(c.ID)
	if conn != nil {
		return conn
	}
	// 启动client
	connChannl := make(chan ziface.IConnection, 1)
	go clientSatrt(c.IP, c.Port, connChannl, c.ID)
	dmanager.ClientConnManager.Add(c.ID, <-connChannl)
	conn, _ = dmanager.ClientConnManager.Get(c.ID)
	return conn
}

func clientSatrt(ip string, port int, connChannl chan ziface.IConnection, nodeId int) ziface.IConnection {
	// 启动client
	client := znet.NewClient(ip, port)
	client.SetOnConnStart(func(i ziface.IConnection) {
		connChannl <- i
	})
	// client.StartHeartBeat(3 * time.Second)
	client.AddInterceptor(&zdecoder.TLVDecoder{})
	client.AddInterceptor(&ClientInterceptor{})
	client.SetDecoder(nil)
	// 删除失效链接
	client.SetOnConnStop(func(conn ziface.IConnection) {
		fmt.Println("内部客户端删除链接")
		dmanager.ClientConnManager.Remove(nodeId)
	})
	client.Start()
	select {}
}
