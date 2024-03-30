package main

import (
	"fmt"
	"time"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// 客户端自定义业务
func loginLoop(conn ziface.IConnection) {
	// for {
	err := conn.SendMsg(3, []byte("Ping...Ping...Ping...[FromClient]"))
	if err != nil {
		fmt.Println(err)
		// break
	}

	time.Sleep(1 * time.Second)
	// }
}

// 创建连接的时候执行
func onClientStart(conn ziface.IConnection) {
	fmt.Println("onClientStart is Called ... ")
	go loginLoop(conn)
}

// Ping test custom routing.
type PingRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	zlog.Debug("Call PingRouter Handle")
	zlog.Debug("recv from server : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}
func main() {
	//创建Client客户端
	client := znet.NewClient("127.0.0.1", 10000)
	//设置链接建立成功后的钩子函数
	client.SetOnConnStart(onClientStart)
	client.AddRouter(2, &PingRouter{})
	//启动客户端
	client.Start()

	//防止进程退出，等待中断信号
	select {}
}
