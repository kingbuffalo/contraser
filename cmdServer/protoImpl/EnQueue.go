package protoImpl

/*
import (
	//"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/cmdServer/protoImpl/comFunc"
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/army_db"
	"buffalo/king/dbOper/hero_db"
	"buffalo/king/king"
)

type EnQueue struct {
	proto_t
}

func (p *EnQueue) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	var c2s king.C2S_EnQueue
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(4031)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())

	heroId1 := *c2s.HeroId1
	heroId2 := *c2s.HeroId2
	heroId3 := *c2s.HeroId3
	armyId := *c2s.ArmyId
	playerId := rpcC2SProto.PlayerId

	heroVOs := [](*hero_db.HeroVO){}
	armyVOs := [](*army_db.ArmyVO){}
	errInt := comFunc.EnQueueSetPos(&armyVOs, &heroVOs, playerId, heroId1, armyId, enumErrCode.ARMY_POS_HEAD, 4032, 4033)
	if errInt != 0 {
		return common.NewErrorCodeRet(errInt)
	}
	errInt = comFunc.EnQueueSetPos(&armyVOs, &heroVOs, playerId, heroId2, armyId, enumErrCode.ARMY_POS_MID, 4034, 4035)
	if errInt != 0 {
		return common.NewErrorCodeRet(errInt)
	}
	errInt = comFunc.EnQueueSetPos(&armyVOs, &heroVOs, playerId, heroId3, armyId, enumErrCode.ARMY_POS_TAIL, 4036, 4037)
	if errInt != 0 {
		return common.NewErrorCodeRet(errInt)
	}

	errInt = comFunc.EnQueueSaveHeroArmy(&armyVOs, &heroVOs, 4042, 4043)
	if errInt != 0 {
		return common.NewErrorCodeRet(errInt)
	}

	s2c := &king.S2C_EnQueue{
		HeroId1: proto.Int32(heroId1),
		HeroId2: proto.Int32(heroId2),
		HeroId3: proto.Int32(heroId3),
		ArmyId:  proto.Int32(armyId),
	}
	return getProtoS2CDataWithoutPush(playerId, "S2C_EnQueue", s2c)
}*/
