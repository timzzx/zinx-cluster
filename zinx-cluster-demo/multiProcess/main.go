package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/aceld/zinx/ziface"
	"github.com/timzzx/zinx-cluster/dconf"
	"github.com/timzzx/zinx-cluster/ddict"
	"github.com/timzzx/zinx-cluster/zinx-cluster-demo/multiProcess/gate"
	"github.com/timzzx/zinx-cluster/zinx-cluster-demo/multiProcess/im"
)

func main() {
	// 获取需要启动的NodeID
	nodeId := flag.Int("id", 0, "id")
	flag.Parse()

	// Node集合
	apps := make(map[string]func(n *ddict.NodeInfo, groupName ddict.GroupName) ziface.IServer, 0)
	apps["gate"] = gate.App
	apps["im"] = im.App

	var servers []ziface.IServer
	// 获取配置
	c := dconf.Dicts.NodeList()
	if *nodeId != 0 {
		for k, v := range c {
			for _, n := range v {
				if *nodeId == n.ID {
					s := apps[k](n, k)
					servers = append(servers, s)
					// 启动
					go s.Serve()
				}
			}
		}
	} else {
		for k, v := range c {
			for _, n := range v {
				s := apps[k](n, k)
				servers = append(servers, s)
				// 启动
				go s.Serve()
			}
		}
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt) // TODO kill -9 mac上有点问题
	sig := <-exit
	fmt.Println("===exit===", sig)
	for _, node := range servers {
		node.Stop()
	}
}
