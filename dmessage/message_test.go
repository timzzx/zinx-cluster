package dmessage

import (
	"testing"

	"github.com/aceld/zinx/zpack"
)

func TestEncode(t *testing.T) {
	_, err := Encode(1, 1, "gate", 1, []byte("test"))
	if err != nil {
		t.Error("解析错误:" + err.Error())
	}
}

func TestDecode(t *testing.T) {
	data, err := Encode(1, 1, "gate", 2, []byte("test"))
	if err != nil {
		t.Error("解析错误:" + err.Error())
		return
	}
	m, err := Decode(data)
	if err != nil {
		t.Error("解析错误:" + err.Error())
		return
	}
	if m.ConnID != 1 {
		t.Error("解析错误: 数据不对")
		return
	}
}

func TestZinxMessageDecode(t *testing.T) {

	dp := zpack.NewDataPack()
	msg, _ := dp.Pack(zpack.NewMsgPackage(1, []byte("test")))
	_, err := Decode(msg)
	if err.Error() != "EOF" {
		t.Error("解析错误:" + err.Error())
		return
	}
}
