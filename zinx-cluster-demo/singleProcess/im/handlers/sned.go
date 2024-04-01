package handlers

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/timzzx/zinx-cluster/dmessage"
)

func Send(request ziface.IRequest) {
	fmt.Println("服务器名称:", request.GetConnection().GetName())
	fmt.Println("接收消息id:", request.GetMsgID())
	decode, _ := dmessage.Decode(request.GetData())
	data, _ := dmessage.Encode(decode.ConnID, 0, "", 2, []byte("Pong...Pong...Pong..."))
	request.GetConnection().SendMsg(2, data)
}

func Forward(request ziface.IRequest) {
	fmt.Println("服务器名称:", request.GetConnection().GetName())
	fmt.Println("接收消息id:", request.GetMsgID())
	decode, _ := dmessage.Decode(request.GetData())
	data, _ := dmessage.Encode(decode.ConnID, 1, "gate", 1, []byte("Forward...Forward...Forward..."))
	request.GetConnection().SendMsg(3, data)
}
