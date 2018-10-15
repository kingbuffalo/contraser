package protoImpl

import (
	//"buffalo/king/battle"
	"buffalo/king/battle"
	"buffalo/king/cfgmgr/cfgPVELevelMgr"
	"buffalo/king/common"
	"buffalo/king/dbOper/army_db"
	//"buffalo/king/dbOper/battleResult_db"
	"buffalo/king/dbOper/hero_db"
	"buffalo/king/dbOper/pveLevel_db"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	//"time"
)

type FightPVE struct {
	proto_t
}

func doFightPVE(heroVOs []*hero_db.HeroVO, pveLevelVO *pveLevel_db.PVELevelVO, playerId int32, cfgPVELevel *cfgPVELevelMgr.CfgPVELevel) ([]*king.Round, int) {
	b := battle.NewBattlePVELevel(heroVOs, cfgPVELevel)
	b.SetPlayerId(playerId)
	r := b.StartFight()
	b.Save()
	return r, 0
}

func (p *FightPVE) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if p.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s king.C2S_FightPVE
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(7011)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())
	playerId := rpcC2SProto.PlayerId

	levelId := *c2s.LevelId
	armyId := *c2s.ArmyId
	armyVO, err := army_db.GetArmyVO(playerId, armyId)
	if err != nil {
		return common.NewErrorCodeRet(7012)
	}

	if armyVO.BEmpty() {
		return common.NewErrorCodeRet(7014)
	}

	pveLevelVO, err := pveLevel_db.GetPVELevelVO(playerId)
	if err != nil {
		return common.NewErrorCodeRet(7013)
	}
	cfgPVELevel := cfgPVELevelMgr.GetCfgPVELevel(levelId)
	if cfgPVELevel == nil {
		return common.NewErrorCodeRet(7015)
	}

	if pveLevelVO.LevelId == 0 {
		if levelId != cfgPVELevelMgr.GetFirstLevelId() {
			return common.NewErrorCodeRet(7018)
		}
	} else {
		maxLv_cfg := cfgPVELevelMgr.GetCfgPVELevel(pveLevelVO.LevelId)
		if !maxLv_cfg.IsValid(levelId) {
			return common.NewErrorCodeRet(7016)
		}
	}

	heroVOs, err := armyVO.GetHeroVOs()
	if err != nil {
		seelog.WarnWE("armyVO.GetHeroVOs err=", err)
		return common.NewErrorCodeRet(7017)
	}

	r, errInt := doFightPVE(heroVOs, pveLevelVO, playerId, cfgPVELevel)
	if errInt != 0 {
		return common.NewErrorCodeRet(errInt)
	}

	s2c := &king.S2C_FightPVE{
		Rounds:  r,
		Useless: proto.Int32(0),
	}
	return getProtoS2CDataWithoutPush(playerId, "S2C_FightPVE", s2c)
}
