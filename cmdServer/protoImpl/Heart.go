package protoImpl

import (
	"github.com/golang/protobuf/proto"
	//"github.com/kingbuffalo/seelog"
	"time"
	"buffalo/king/common"
	"buffalo/king/king"
)

type Heart struct {
	proto_t
}

func (g *Heart) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	c2s_heart := &king.C2S_Heart{}
	err := proto.Unmarshal(rpcC2SProto.Data, c2s_heart)
	if err != nil {
		return common.NewErrorCodeRet(1041)
	}

	//seelog.Info(c2s_heart.String())
	ct := time.Now().Unix()
	s2c_heart := &king.S2C_Heart{
		ServerTimestamp: proto.Int32(int32(ct)),
	}
	return getProtoS2CDataWithoutPush(0, "S2C_Heart", s2c_heart)
}
