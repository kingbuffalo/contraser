package protoImpl

import (
	"buffalo/king/cfgmgr/cfgHeroRecruitMgr"
	"buffalo/king/cmdServer/protoImpl/comFunc"
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/hero_db"
	"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
)

type HeroRecruit struct {
	proto_t
}

func bHeroExist(heroVOs *[]*hero_db.HeroVO, heroId int32) bool {
	for _, v := range *heroVOs {
		if v.Id == heroId {
			return true
		}
	}
	return false
}

func (p *HeroRecruit) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if p.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s king.C2S_HeroRecruit
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(4131)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())
	playerId := rpcC2SProto.PlayerId

	rtype := *c2s.Type
	cfg := cfgHeroRecruitMgr.GetCfgHeroRecruit(rtype)
	if cfg == nil {
		return common.NewErrorCodeRet(4132)
	}

	heroVOs, err := hero_db.GetHeroVOs(playerId)
	if err != nil {
		seelog.WarnWE("getherovos=", err)
		return common.NewErrorCodeRet(4134)
	}

	playerInfoVO, err := playerInfo_db.GetPlayerInfoVOBy(playerId)
	if err != nil {
		seelog.WarnWE("getPlayerInfoKey=", err)
		return common.NewErrorCodeRet(4133)
	}
	if !playerInfoVO.BHasEnought(cfg.MoneyType, cfg.MoneyNum) {
		return common.NewErrorCodeRet(4135)
	}
	playerInfoVO.AddRes(cfg.MoneyType, -cfg.MoneyNum)
	addlocidcountCfg, accTimes, tenTimes := cfgHeroRecruitMgr.RandHeroId(rtype, playerInfoVO.FiveStarHeroRecruitTimes, playerInfoVO.TenTimesHeroRecuritTimes, cfg.MoneyType == enumErrCode.SPEC_RESCOURCE_DIAMOND)

	heroIds := make([]int32, len(addlocidcountCfg))
	toChipHeroIds := make([]int32, 0)
	for i, v := range addlocidcountCfg {
		heroIds[i] = v[1]
		if bHeroExist(&heroVOs, v[1]) {
			toChipHeroIds = append(toChipHeroIds, v[1])
		}
	}
	seelog.Trace("HeroRecruit---------------->", heroVOs)
	alic := comFunc.NewAddLocIdCount(playerId, playerInfoVO, heroVOs)
	if err := alic.AddCfg(addlocidcountCfg); err != nil {
		seelog.WarnWE("alic.AddCfg err=", err)
		return common.NewErrorCodeRet(4134)
	}
	if err := alic.DiffSave(); err != nil {
		seelog.WarnWE("alic.DiffSave err=", err)
		return common.NewErrorCodeRet(4135)
	}

	playerInfoVO.TenTimesHeroRecuritTimes = tenTimes
	playerInfoVO.FiveStarHeroRecruitTimes = accTimes

	tenthTimesLeft := enumErrCode.HERO_RECRUIT_SPEC_TIMES - playerInfoVO.TenTimesHeroRecuritTimes

	s2c := &king.S2C_HeroRecruit{
		HeroIds:        heroIds,
		ToChipHeroIds:  toChipHeroIds,
		FinalChip:      proto.Int32(0),
		Chip:           proto.Int32(0),
		TenthTimesLeft: proto.Int32(tenthTimesLeft),
	}
	return getProtoS2CDataWithoutPush(playerId, "S2C_HeroRecruit", s2c)
}
