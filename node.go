package nodecom

import (
	"github.com/lendloan/lendproto/loginnode"
	"github.com/lendloan/lendproto/registernode"

	"github.com/lendloan/lendproto/common"

	"github.com/lendloan/lendproto/authnode"
	"github.com/lendloan/lendproto/certnode"
	"github.com/lendloan/lendproto/cloudnode"
	"github.com/lendloan/lendproto/codenode"
	"github.com/lendloan/lendproto/cronnode"
	"github.com/lendloan/lendproto/dartynode"
	"github.com/lendloan/lendproto/datanode"
	"github.com/lendloan/lendproto/limitnode"
	"github.com/lendloan/lendproto/lognode"
	"github.com/lendloan/lendproto/macipnode"
	"github.com/lendloan/lendproto/paynode"
	"github.com/lendloan/lendproto/pkgnode"
	"github.com/lendloan/lendproto/sensinode"
	"github.com/lendloan/lendproto/sharenode"
	"github.com/lendloan/lendproto/usernode"
	"github.com/lendloan/lendproto/friendnode"
	"github.com/lendloan/lendproto/loannode"
	"go-micro.dev/v4/client"
)

// 检查是否有权限
//
// @param auth 	授权参数
// @return bool
//
func Authorize(auth *common.Authorize) bool {
	if nil == auth {
		return false
	}

	return true
}

// 获取有数据节点客户端
//
// @param s2sname
//
func Datanode(s2sname string, cli client.Client) datanode.DatanodeService {
	reqcli := datanode.NewDatanodeService(s2sname, cli)

	return reqcli
}

// 获取有数据节点客户端
//
// @param s2sname
//
func Codenode(s2sname string, cli client.Client) codenode.CodenodeService {
	reqcli := codenode.NewCodenodeService(s2sname, cli)

	return reqcli
}

// 获取cloud节点客户端
//
// @param s2sname
//
func Cloudnode(s2sname string, cli client.Client) cloudnode.CloudnodeService {
	reqcli := cloudnode.NewCloudnodeService(s2sname, cli)

	return reqcli
}

// 获取register节点客户端
//
// @param s2sname
//
func Registernode(s2sname string, cli client.Client) registernode.RegisternodeService {
	reqcli := registernode.NewRegisternodeService(s2sname, cli)

	return reqcli
}

// 获取login节点客户端
//
// @param s2sname
//
func Loginnode(s2sname string, cli client.Client) loginnode.LoginnodeService {
	reqcli := loginnode.NewLoginnodeService(s2sname, cli)

	return reqcli
}

// 获取login节点客户端
//
// @param s2sname
//
func Usernode(s2sname string, cli client.Client) usernode.UsernodeService {
	reqcli := usernode.NewUsernodeService(s2sname, cli)

	return reqcli
}

// 获取login节点客户端
//
// @param s2sname
//
func Authnode(s2sname string, cli client.Client) authnode.AuthnodeService {
	reqcli := authnode.NewAuthnodeService(s2sname, cli)

	return reqcli
}

// 获取darty节点客户端
//
// @param s2sname
//
func Dartynode(s2sname string, cli client.Client) dartynode.DartynodeService {
	reqcli := dartynode.NewDartynodeService(s2sname, cli)

	return reqcli
}

// 获取sensi节点客户端
//
// @param s2sname
//
func Sensinode(s2sname string, cli client.Client) sensinode.SensinodeService {
	reqcli := sensinode.NewSensinodeService(s2sname, cli)

	return reqcli
}

// 获取limit节点客户端
//
// @param s2sname
//
func Limitnode(s2sname string, cli client.Client) limitnode.LimitnodeService {
	reqcli := limitnode.NewLimitnodeService(s2sname, cli)

	return reqcli
}

// 创建cert客户端
//
func Certnode(s2sname string, cli client.Client) certnode.CertnodeService {
	reqcli := certnode.NewCertnodeService(s2sname, cli)

	return reqcli
}

// 创建macipnode客户端
//
func Macipnode(s2sname string, cli client.Client) macipnode.MacipnodeService {
	reqcli := macipnode.NewMacipnodeService(s2sname, cli)

	return reqcli
}

// 获取Lognode客户端
//
func Lognode(s2sname string, cli client.Client) lognode.LognodeService {
	reqcli := lognode.NewLognodeService(s2sname, cli)

	return reqcli
}

// 获取Pkgnode客户端
//
func Pkgnode(s2sname string, cli client.Client) pkgnode.PkgnodeService {
	reqcli := pkgnode.NewPkgnodeService(s2sname, cli)

	return reqcli
}

// 获取Cronnode客户端
//
func Cronnode(s2sname string, cli client.Client) cronnode.CronnodeService {
	reqcli := cronnode.NewCronnodeService(s2sname, cli)

	return reqcli
}

// 获取paynode客户端
//
func Paynode(s2sname string, cli client.Client) paynode.PaynodeService {
	reqcli := paynode.NewPaynodeService(s2sname, cli)

	return reqcli
}

// 获取mesharenode客户端
//
func Sharenode(s2sname string, cli client.Client) sharenode.SharenodeService {
	reqcli := sharenode.NewSharenodeService(s2sname, cli)

	return reqcli
}

// 获取friendnode客户端
//
func Friendnode(s2sname string, cli client.Client) friendnode.FriendnodeService {
	reqcli := friendnode.NewFriendnodeService(s2sname, cli)

	return reqcli
}

// 获取loannode客户端
//
func Loannode(s2sname string, cli client.Client) loannode.FriendnodeService {
	reqcli := loannode.NewLoannodeService(s2sname, cli)

	return reqcli
}