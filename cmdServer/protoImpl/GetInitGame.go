package protoImpl

import (
	"buffalo/king/common"
	"buffalo/king/dbOper/army_db"
	"buffalo/king/dbOper/build_db"
	"buffalo/king/dbOper/hero_db"
	"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
)

type GetInitGame struct {
	proto_t
}

func (I *GetInitGame) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	seelog.Trace("GetInitGame")
	if I.bClose {
		return common.NewErrorCodeRet(107)
	}
	c2s_getInitGame := &king.C2S_GetInitGame{}
	err := proto.Unmarshal(rpcC2SProto.Data, c2s_getInitGame)
	if err != nil {
		seelog.Trace("proto Unmarshal error!")
		return common.NewErrorCodeRet(2001)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s_getInitGame.String())
	playerId := rpcC2SProto.PlayerId

	heroVOs, err := hero_db.GetHeroVOs(rpcC2SProto.PlayerId)
	seelog.Trace(len(heroVOs))
	if err != nil {
		seelog.Trace("GetHeros error")
		return common.NewErrorCodeRet(2004)
	}

	playerInfoVO, err := playerInfo_db.GetPlayerInfoVOBy(rpcC2SProto.PlayerId)
	if err != nil {
		seelog.Trace("GetPlayerInfoVOBy error")
		return common.NewErrorCodeRet(2005)
	}

	armyVOs, err := army_db.GetArmyVOs(rpcC2SProto.PlayerId)
	if err != nil {
		seelog.Trace("GetArmyVOs error")
		return common.NewErrorCodeRet(2006)
	}

	buildVO, err := build_db.GetBuildingVO(rpcC2SProto.PlayerId)
	if err != nil {
		seelog.Trace("GetBuildVO error")
		return common.NewErrorCodeRet(2007)
	}

	pheros := [](*king.Hero){}
	for _, h := range heroVOs {
		seelog.Trace(h)
		phero := &king.Hero{
			Id:           proto.Int32(h.Id),
			Exp:          proto.Int32(h.Exp),
			Level:        proto.Int32(h.Level),
			Star:         proto.Int32(h.Star),
			ArmyLv:       proto.Int32(h.ArmyLv),
			ArmySkill1Lv: proto.Int32(h.ArmySkill1Lv),
			ArmySkill2Lv: proto.Int32(h.ArmySkill2Lv),
			ArmySkill3Lv: proto.Int32(h.ArmySkill3Lv),
			Chip:         proto.Int32(h.Chip),
		}
		pheros = append(pheros, phero)
	}

	parmys := [](*king.Army){}
	for _, a := range armyVOs {
		pArmy := &king.Army{
			ArmyId:  proto.Int32(a.ArmyId),
			HeroIds: a.GetHeroIds(),
		}
		parmys = append(parmys, pArmy)
	}
	build := &king.Building{
		Farm1Lv: proto.Int32(buildVO.Farm1Lv),
		Farm2Lv: proto.Int32(buildVO.Farm2Lv),
	}

	pplayerInfo := &king.PlayerInfo{
		Coin:    proto.Int32(playerInfoVO.Coin),
		Diamond: proto.Int32(playerInfoVO.Diamond),
		Level:   proto.Int32(playerInfoVO.Level),
		Wood:    proto.Int32(playerInfoVO.Wood),
		Mineral: proto.Int32(playerInfoVO.Mineral),
		Grain:   proto.Int32(playerInfoVO.Grain),
	}

	s2c := &king.S2C_GetInitGame{
		Heros:      pheros,
		Armys:      parmys,
		PlayerInfo: pplayerInfo,
		Building:   build,
	}

	return getProtoS2CDataWithoutPush(playerId, "S2C_GetInitGame", s2c)
}
