package handlers

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/timzzx/zinx-cluster/dmessage"
)

// 登录
func Login(request ziface.IRequest) {
	fmt.Println("服务器名称:", request.GetConnection().GetName())
	fmt.Println("接收消息id:", request.GetMsgID(), "[登录成功]")
	decode, _ := dmessage.Decode(request.GetData())
	data, _ := dmessage.Encode(decode.ConnID, 0, "", 2, []byte("登录成功"))
	request.GetConnection().SendMsg(1, data)
}

// 退出
func Logout(request ziface.IRequest) {
	fmt.Println("服务器名称:", request.GetConnection().GetName())
	fmt.Println("接收消息id:", request.GetMsgID(), "[退出成功]")

	decode, _ := dmessage.Decode(request.GetData())
	data, _ := dmessage.Encode(decode.ConnID, 0, "", 2, []byte("退出成功"))
	request.GetConnection().SendMsg(2, data)
}
