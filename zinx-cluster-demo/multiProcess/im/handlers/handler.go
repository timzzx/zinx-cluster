package handlers

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/timzzx/zinx-cluster/dmanager"
	"github.com/timzzx/zinx-cluster/dmessage"
)

func Send(request ziface.IRequest) {
	fmt.Println("服务器名称:", request.GetConnection().GetName())
	fmt.Println("接收消息id:", request.GetMsgID())
	// decode, _ := dmessage.Decode(request.GetData())
	data, _ := dmessage.Encode(uint64(2), 0, "", 2, []byte("Send..."))
	request.GetConnection().SendMsg(3, data)
}

func FullSend(request ziface.IRequest) {
	fmt.Println("服务器名称:", request.GetConnection().GetName())
	fmt.Println("接收消息id:", request.GetMsgID())
	// decode, _ := dmessage.Decode(request.GetData())
	members := dmanager.MemberConnManager.GetAllConnID()
	for _, v := range members {
		data, _ := dmessage.Encode(v, 1, "gate", 2, []byte("FullSend..."))
		request.GetConnection().SendMsg(4, data)
	}
}
