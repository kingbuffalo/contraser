package protoImpl

import (
	"github.com/golang/protobuf/proto"
	//"github.com/kingbuffalo/seelog"
	//"buffalo/king/battle/march"
	"buffalo/king/common"
	"buffalo/king/king"
)

type Ping struct {
	proto_t
}

func (g *Ping) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	c2s_ping := &king.C2S_Ping{}
	err := proto.Unmarshal(rpcC2SProto.Data, c2s_ping)
	if err != nil {
		return common.NewErrorCodeRet(1031)
	}
	//seelog.Info(c2s_ping.String())

	s2c_ping := &king.S2C_Ping{
		ClientTimestamp: c2s_ping.ClientTimestamp,
	}

	return getProtoS2CDataWithoutPush(0, "S2C_Ping", s2c_ping)
}
