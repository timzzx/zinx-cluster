package dnode

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/timzzx/zinx-cluster/dclient"
	"github.com/timzzx/zinx-cluster/dconf"
	"github.com/timzzx/zinx-cluster/ddict"
	"github.com/timzzx/zinx-cluster/dmanager"
	"github.com/timzzx/zinx-cluster/dmessage"
)

type NodeInterceptor struct {
	Node      *ddict.NodeInfo
	GroupName string
}

func (m *NodeInterceptor) Intercept(chain ziface.IChain) ziface.IcResp {
	request := chain.Request() //从责任链中获取当前拦截器的输入数据

	// 这一层是自定义拦截器处理逻辑，这里只是简单打印输入
	iRequest := request.(ziface.IRequest) //注意：由于Zinx的Request类型，这里需要做一下断言转换

	// 保存用户coon
	dmanager.MemberConnManager.Add(iRequest.GetConnection())
	// 消息封包
	data, err := dmessage.Encode(iRequest.GetConnection().GetConnID(), 0, "", iRequest.GetData())
	// 消息转发
	d := dconf.Dicts.GetRouteDicts()
	if r, ok := d[iRequest.GetMsgID()]; ok {
		// 判断是否为当前gate数据
		if r.GroupName != m.GroupName {
			// 获取连接
			c := dclient.Client(r.GroupName, 0)
			if err != nil {
				fmt.Println("封包错误")
				iRequest.Abort()
			}
			err = c.SendMsg(iRequest.GetMsgID(), data)
			if err != nil {
				fmt.Println("发送数据错误")
				iRequest.Abort()
			}
		}
	}
	iRequest.GetMessage().SetData(data)
	iRequest.GetMessage().SetDataLen(uint32(len(data)))
	return chain.Proceed(chain.Request()) //进入并执行下一个拦截器
}
