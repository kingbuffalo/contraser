package protoImpl

/*

import (
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/dbOper/battleResult_db"
	"buffalo/king/king"
)

type GetAllBattleResult struct {
	proto_t
}

func (p *GetAllBattleResult) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if p.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s king.C2S_GetAllBattleResult
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(2141)
	}
	playerId := rpcC2SProto.PlayerId
	seelog.Info(playerId, ",", c2s.String())
	brvos, err := battleResult_db.GetBattleResultVOs(playerId)
	if err != nil {
		return common.NewErrorCodeRet(2142)
	}

	if len(brvos) > 0 {
		brs := make([](*king.BriefBattleResult), len(brvos))
		for i, _ := range brs {
			brs[i] = brvos[i].GenBriefPushData()
		}
		s2c := &king.S2C_GetAllBattleResult{
			BriefBattleResults: brs,
			UslessValue:        proto.Int32(0),
		}
		return getProtoS2CDataWithoutPush(playerId, "S2C_GetAllBattleResult", s2c)
	} else {
		s2c := &king.S2C_GetAllBattleResult{
			BriefBattleResults: [](*king.BriefBattleResult){},
			UslessValue:        proto.Int32(0),
		}
		return getProtoS2CDataWithoutPush(playerId, "S2C_GetAllBattleResult", s2c)
	}
}*/
