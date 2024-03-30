package dconf

import (
	"testing"
)

func TestConf(t *testing.T) {
	if Dicts.NodeList()["im"][0].ID != 2 {
		t.Fatal("配置出错")
	}
}
