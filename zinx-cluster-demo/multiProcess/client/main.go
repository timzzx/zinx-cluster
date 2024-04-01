package main

import (
	"fmt"
	"time"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// 客户端自定义业务
func loginLoop(conn ziface.IConnection) {

	err := conn.SendMsg(1, []byte("登录[FromClient]"))
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1 * time.Second)
	err = conn.SendMsg(2, []byte("退出[FromClient]"))
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1 * time.Second)
	err = conn.SendMsg(3, []byte("单发[FromClient]"))
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1 * time.Second)
	err = conn.SendMsg(4, []byte("群发[FromClient]"))
	if err != nil {
		fmt.Println(err)
		return
	}

}

// 创建连接的时候执行
func onClientStart(conn ziface.IConnection) {
	fmt.Println("onClientStart is Called ... ")
	go loginLoop(conn)
}

type LoginRouter struct {
	znet.BaseRouter
}

func (this *LoginRouter) Handle(request ziface.IRequest) {
	fmt.Println("登录成功")
}

type LogoutRouter struct {
	znet.BaseRouter
}

func (this *LogoutRouter) Handle(request ziface.IRequest) {
	fmt.Println("退出成功")
}

type SendRouter struct {
	znet.BaseRouter
}

func (this *SendRouter) Handle(request ziface.IRequest) {
	fmt.Println("单发成功")
}

type FullSendRouter struct {
	znet.BaseRouter
}

func (this *FullSendRouter) Handle(request ziface.IRequest) {
	fmt.Println("全发成功")
}
func main() {
	//创建Client客户端
	client := znet.NewClient("127.0.0.1", 10000)
	//设置链接建立成功后的钩子函数
	client.SetOnConnStart(onClientStart)
	client.AddRouter(1, &LoginRouter{})
	client.AddRouter(2, &LogoutRouter{})
	client.AddRouter(3, &SendRouter{})
	client.AddRouter(4, &FullSendRouter{})
	//启动客户端
	client.Start()

	//防止进程退出，等待中断信号
	select {}
}
