package protoImpl

/*

import (
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/dbOper/army_db"
	"buffalo/king/dbOper/city_db"
	"buffalo/king/king"
)

type SetDefArmy struct {
	proto_t
}

func (p *SetDefArmy) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if p.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s king.C2S_SetDefArmy
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(4051)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())
	armyId := *c2s.ArmyId
	playerId := rpcC2SProto.PlayerId
	armyVO, err := army_db.GetArmyVO(playerId, armyId)
	if err != nil {
		seelog.WarnWE("GetArmyVO err=", err)
		return common.NewErrorCodeRet(4052)
	}
	if armyVO == nil {
		return common.NewErrorCodeRet(4053)
	}
	cityId := *c2s.CityId
	cityVO, err := city_db.GetCityVO(playerId, cityId)
	if err != nil {
		seelog.WarnWE("GetCityVO err=", err)
		return common.NewErrorCodeRet(4054)
	}
	if cityVO == nil {
		return common.NewErrorCodeRet(4055)
	}
	old_cityVO := cityVO.Clone()
	cityVO.DefenceArmyId = armyId
	if err = cityVO.DiffSave(old_cityVO); err != nil {
		seelog.WarnWE("SaveCityVO err=", err)
		return common.NewErrorCodeRet(4056)
	}

	var s2c king.S2C_SetDefArmy
	s2c.ArmyId = proto.Int32(armyId)
	s2c.CityId = proto.Int32(cityId)
	return getProtoS2CDataWithoutPush(playerId, "S2C_SetDefArmy", &s2c)
}*/
