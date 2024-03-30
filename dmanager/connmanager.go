package dmanager

// 内部链接管理
import (
	"errors"
	"strconv"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/zutils"
)

type ConnManager struct {
	connections zutils.ShardLockMaps
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: zutils.NewShardLockMaps(),
	}
}

func (connMgr *ConnManager) Add(Id int, conn ziface.IConnection) {
	strConnId := strconv.Itoa(Id)
	connMgr.connections.Set(strConnId, conn) // 将conn连接添加到ConnManager中

	zlog.Ins().InfoF("connection add to ConnManager successfully: conn num = %d", connMgr.Len())
}

func (connMgr *ConnManager) Remove(Id int) {
	strConnId := strconv.Itoa(Id)
	connMgr.connections.Remove(strConnId) // 删除连接信息

	zlog.Ins().InfoF("connection Remove ConnID=%d successfully: conn num = %d", Id, connMgr.Len())
}

func (connMgr *ConnManager) Get(Id int) (ziface.IConnection, error) {

	strConnId := strconv.Itoa(Id)
	if conn, ok := connMgr.connections.Get(strConnId); ok {
		return conn.(ziface.IConnection), nil
	}

	return nil, errors.New("connection not found")
}

func (connMgr *ConnManager) Len() int {

	length := connMgr.connections.Count()

	return length
}

func (connMgr *ConnManager) ClearConn() {

	// Stop and delete all connection information
	for item := range connMgr.connections.IterBuffered() {
		val := item.Val
		if conn, ok := val.(ziface.IConnection); ok {
			// stop will eventually trigger the deletion of the connection,
			// no additional deletion is required
			conn.Stop()
		}
	}

	zlog.Ins().InfoF("Clear All Connections successfully: conn num = %d", connMgr.Len())
}

func (connMgr *ConnManager) GetAllConnID() []int {

	strConnIdList := connMgr.connections.Keys()
	ids := make([]int, 0, len(strConnIdList))

	for _, strId := range strConnIdList {
		connId, err := strconv.Atoi(strId)
		if err == nil {
			ids = append(ids, connId)
		} else {
			zlog.Ins().InfoF("GetAllConnID Id: %d, error: %v", connId, err)
		}
	}

	return ids
}

func (connMgr *ConnManager) GetAllConnIdStr() []string {
	return connMgr.connections.Keys()
}
