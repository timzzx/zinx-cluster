package handlers

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
)

func Login(request ziface.IRequest) {
	fmt.Println("服务器名称:", request.GetConnection().GetName())
	fmt.Println("接收消息id:", request.GetMsgID())
}
