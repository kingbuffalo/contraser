package protoImpl

import (
	"buffalo/king/cmdServer/protoImpl/comFunc"
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/army_db"
	"buffalo/king/dbOper/hero_db"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
)

type EnQueueSingle struct {
	proto_t
}

func (p *EnQueueSingle) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if p.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s king.C2S_EnQueueSingle
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(4071)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())
	playerId := rpcC2SProto.PlayerId

	heroId := *c2s.HeroId
	pos := *c2s.Pos
	armyId := *c2s.ArmyId

	posInt := int(pos)

	if posInt != enumErrCode.ARMY_POS_1 && posInt != enumErrCode.ARMY_POS_2 && posInt != enumErrCode.ARMY_POS_3 && posInt != enumErrCode.ARMY_POS_4 && posInt != enumErrCode.ARMY_POS_5 {
		return common.NewErrorCodeRet(4072)
	}

	heroVOs := [](*hero_db.HeroVO){}
	armyVOs := [](*army_db.ArmyVO){}
	errInt := comFunc.EnQueueSetPos(&armyVOs, &heroVOs, playerId, heroId, armyId, posInt, 4073, 4074)
	if errInt != 0 {
		return common.NewErrorCodeRet(errInt)
	}

	errInt = comFunc.EnQueueSaveHeroArmy(&armyVOs, &heroVOs, 4075, 4076)
	if errInt != 0 {
		return common.NewErrorCodeRet(errInt)
	}

	s2c := &king.S2C_EnQueueSingle{
		HeroId: proto.Int32(heroId),
		Pos:    proto.Int32(pos),
		ArmyId: proto.Int32(armyId),
	}

	return getProtoS2CDataWithoutPush(playerId, "S2C_EnQueueSingle", s2c)
}
