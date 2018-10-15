package protoImpl

import (
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/king"
)

type ReConnect struct {
	proto_t
}

func (L *ReConnect) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	c2s := &king.C2S_ReConnect{}
	err := proto.Unmarshal(rpcC2SProto.Data, c2s)
	if err != nil {
		seelog.Trace("proto Unmarshal error!")
		return common.NewErrorCodeRet(1061)
	}
	seelog.Info(c2s.String())

	//TODO
	//tokenInRedis := gameutil.GenToken(*c2s.PlayerId)
	//if tokenInRedis != *c2s.Token {
	//return common.NewErrorCodeRet(1067)
	//}

	playerInfoVO, err := playerInfo_db.GetPlayerInfoVOBy(*c2s.PlayerId)
	if err != nil {
		seelog.Trace("GetPlayerInfoVOBy error=", err)
		return common.NewErrorCodeRet(1063)
	}
	if playerInfoVO == nil {
		seelog.Trace(" ok! playerInfoVO is nil")
		return common.NewErrorCodeRet(1064)
	}
	s2c_reconnect := &king.S2C_ReConnect{
		PlayerId: proto.Int32(*c2s.PlayerId),
	}

	seelog.Trace(" ok!", s2c_reconnect)
	return getProtoS2CDataWithoutPush(0, "S2C_ReConnect", s2c_reconnect)
}
