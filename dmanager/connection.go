package dmanager

// var MemberConnManager *dmanager.MemberManager
var ClientConnManager *ConnManager
var MemberConnManager *MemberManager

func init() {
	MemberConnManager = NewMemberManager()
	ClientConnManager = NewConnManager()
}
