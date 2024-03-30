package dconf

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/timzzx/zinx-cluster/ddict"
)

const (
	EnvDefaultConfigFilePath = "/conf/dinx.json"
	RouteFilePath            = "/conf/route.json"
)

var Dicts *ddict.Dict

func Reload(d *ddict.Dict) {
	// 配置
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFilePath := filepath.Join(pwd, EnvDefaultConfigFilePath)

	// 不存在配置文件直接退出
	if confFileExists, _ := PathExists(configFilePath); !confFileExists {
		return
	}

	// 读取配置文件
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &d.NodeGroup)
	if err != nil {
		panic(err)
	}
	// 路由
	routeFilePath := filepath.Join(pwd, RouteFilePath)

	// 不存在路由直接退出
	if routeFileExists, _ := PathExists(routeFilePath); !routeFileExists {
		return
	}

	// 读取配置文件
	routeData, err := os.ReadFile(routeFilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(routeData, &d.RouteGroup)
	if err != nil {
		panic(err)
	}
}

// PathExists Check if a file exists.(判断一个文件是否存在)
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func init() {
	Dicts = ddict.NewDict()
	// 默认增加gate配置
	Dicts.AddNodeInfo("gate", &ddict.NodeInfo{
		ID:   1,
		Name: "gate-1",
		Type: 1,
		IP:   "0.0.0.0",
		Port: 10000,
	})
	Reload(Dicts)
}
