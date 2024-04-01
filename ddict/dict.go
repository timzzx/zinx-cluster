package ddict

type RouteInfo struct {
	MsgID uint32 // 消息id
	Name  string // 路由名称
}
type RouteDict struct {
	MsgID     uint32 // 消息id
	Name      string // 路由名称
	GroupName GroupName
}
type NodeInfo struct {
	ID   int    // node id
	Name string // node 名称
	Type int    // node类型 1.gateway 2.backend
	IP   string // IP
	Port int    // 端口号
}
type NodeDict struct {
	ID        int    // node id
	Name      string // node 名称
	Type      int    // node类型 1.gateway 2.backend
	IP        string // IP
	Port      int    // 端口号
	GroupName GroupName
}
type NodeSlice = []*NodeInfo

type RouteSlice = []*RouteInfo

type GroupName = string

// node集合
type NodeGroupMap = map[GroupName]NodeSlice
type NodeDicts = map[int]*NodeDict

// route集合
type RouteGroupMap = map[GroupName]RouteSlice
type RouteDicts = map[uint32]*RouteDict

// 字典
type Dict struct {
	NodeGroup  NodeGroupMap
	RouteGroup RouteGroupMap
}

func NewDict() *Dict {
	return &Dict{
		NodeGroup:  make(NodeGroupMap),
		RouteGroup: make(RouteGroupMap),
	}
}

// 增加node信息到字典中
func (d *Dict) AddNodeInfo(n GroupName, info *NodeInfo) {
	d.DelNodeInfo(n, info.ID)
	if value, ok := d.NodeGroup[n]; ok {
		d.NodeGroup[n] = append(value, info)
	} else {
		d.NodeGroup[n] = NodeSlice{info}
	}
}

// 增加router信息到字典中
func (d *Dict) AddRouteInfo(n GroupName, info *RouteInfo) {
	d.DelRouteInfo(n, info.MsgID)
	if value, ok := d.RouteGroup[n]; ok {
		d.RouteGroup[n] = append(value, info)
	} else {
		d.RouteGroup[n] = RouteSlice{info}
	}
}

// 删除字典中node信息
func (d *Dict) DelNodeInfo(n GroupName, NodeID int) {
	if value, ok := d.NodeGroup[n]; ok {
		j := 0
		for _, v := range value {
			if v.ID != NodeID {
				value[j] = v
				j++
			}
		}
		d.NodeGroup[n] = value[:j]
	}
}

// 删除字典中router信息
func (d *Dict) DelRouteInfo(n GroupName, MsgID uint32) {
	if value, ok := d.RouteGroup[n]; ok {
		j := 0
		for _, v := range value {
			if v.MsgID != MsgID {
				value[j] = v
				j++
			}
		}
		d.RouteGroup[n] = value[:j]
	}
}

// 展示node列表
func (d *Dict) NodeList() NodeGroupMap {
	return d.NodeGroup
}

// 展示Route列表
func (d *Dict) RouteList() RouteGroupMap {
	return d.RouteGroup
}

// Route字典
func (d *Dict) GetRouteDicts() RouteDicts {
	r := make(RouteDicts)

	for groupName, group := range d.RouteGroup {
		for _, v := range group {
			r[v.MsgID] = &RouteDict{
				MsgID:     v.MsgID,
				Name:      v.Name,
				GroupName: groupName,
			}
		}
	}
	return r
}

// node字典
func (d *Dict) GetNodeDicts() NodeDicts {
	r := make(NodeDicts)
	for groupName, group := range d.NodeGroup {
		for _, v := range group {
			r[v.ID] = &NodeDict{
				ID:        v.ID,
				Name:      v.Name,
				Type:      v.Type,
				IP:        v.IP,
				Port:      v.Port,
				GroupName: groupName,
			}

		}
	}
	return r
}
