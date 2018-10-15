package protoImpl

/*

import (
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/dbOper/battleResult_db"
	"buffalo/king/king"
)

type GetBattleResultDetail struct {
	proto_t
}

func (p *GetBattleResultDetail) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if p.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s king.C2S_GetBattleResultDetail
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(2161)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())
	playerId := rpcC2SProto.PlayerId
	battleResultId := *c2s.BattleResultId
	br, err := battleResult_db.GetBattleResultVO(playerId, battleResultId)
	if err != nil {
		seelog.WarnWE("GetBattleResultVO,err=", err)
		return common.NewErrorCodeRet(2162)
	}

	s2c := br.GenDetailPushData()
	return getProtoS2CDataWithoutPush(playerId, "S2C_GetBattleResultDetail", s2c)
}*/
