package protoImpl

import (
	"buffalo/king/battle"
	"buffalo/king/common"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
)

type ContinueFight struct {
	proto_t
}

func (p *ContinueFight) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if p.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s king.C2S_ContinueFight
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(7051)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())
	playerId := rpcC2SProto.PlayerId

	b := battle.GetBattle(playerId)
	if b == nil {
		return common.NewErrorCodeRet(7052)
	}

	skillId1 := *c2s.SkillId1
	skillId2 := *c2s.SkillId2
	skillId3 := *c2s.SkillId3
	r := b.ConRunOneStep(skillId1, skillId2, skillId3)

	s2c := &king.S2C_ContinueFight{
		Rounds:  r,
		Useless: proto.Int32(0),
	}
	return getProtoS2CDataWithoutPush(playerId, "S2C_ContinueFight", s2c)
}
