// Code generated by liverpcgen, DO NOT EDIT.
// source: *.proto files under this directory
// If you want to change this file, Please see README in go-common/app/tool/liverpc/protoc-gen-liverpc/
package liverpc

import (
	"go-common/app/service/live/fans_medal/api/liverpc/v1"
	"go-common/app/service/live/fans_medal/api/liverpc/v2"
	"go-common/library/net/rpc/liverpc"
)

// Client that represents a liverpc fans_medal service api
type Client struct {
	cli *liverpc.Client
	// V1FansMedal presents the controller in liverpc
	V1FansMedal v1.FansMedal
	// V1Medal presents the controller in liverpc
	V1Medal v1.Medal
	// V2Anchor presents the controller in liverpc
	V2Anchor v2.Anchor
	// V2HighQps presents the controller in liverpc
	V2HighQps v2.HighQps
}

// DiscoveryAppId the discovery id is not the tree name
var DiscoveryAppId = "live.fansmedal"

// New a Client that represents a liverpc live.fansmedal service api
// conf can be empty, and it will use discovery to find service by default
// conf.AppID will be overwrite by a fixed value DiscoveryAppId
// therefore is no need to set
func New(conf *liverpc.ClientConfig) *Client {
	if conf == nil {
		conf = &liverpc.ClientConfig{}
	}
	conf.AppID = DiscoveryAppId
	var realCli = liverpc.NewClient(conf)
	cli := &Client{cli: realCli}
	cli.clientInit(realCli)
	return cli
}

func (cli *Client) GetRawCli() *liverpc.Client {
	return cli.cli
}

func (cli *Client) clientInit(realCli *liverpc.Client) {
	cli.V1FansMedal = v1.NewFansMedalRpcClient(realCli)
	cli.V1Medal = v1.NewMedalRpcClient(realCli)
	cli.V2Anchor = v2.NewAnchorRpcClient(realCli)
	cli.V2HighQps = v2.NewHighQpsRpcClient(realCli)
}
