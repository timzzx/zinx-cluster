package dmessage

import (
	"bytes"
	"encoding/gob"
)

// 内部消息封包解包
type NodeMessage struct {
	ConnID    uint64 // 链接ID
	NodeID    int    // nodeID
	NodeGroup string // Node组
	Data      []byte // 客户端消息data
}

func Encode(connID uint64, nodeID int, nodeGroup string, data []byte) ([]byte, error) {
	m := NodeMessage{
		ConnID:    connID,
		NodeID:    nodeID,
		NodeGroup: nodeGroup,
		Data:      data,
	}
	var buf bytes.Buffer
	encode := gob.NewEncoder(&buf)
	err := encode.Encode(m)

	return buf.Bytes(), err
}

func Decode(data []byte) (*NodeMessage, error) {
	m := &NodeMessage{}
	buf := bytes.NewBuffer(data)
	decode := gob.NewDecoder(buf)
	err := decode.Decode(m)
	return m, err
}
