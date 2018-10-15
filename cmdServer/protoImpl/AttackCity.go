package protoImpl

/*
import (
	"buffalo/king/common"
	"buffalo/king/dbOper/army_db"
	"buffalo/king/dbOper/city_db"
	"buffalo/king/dbOper/mapCity_db"
	"buffalo/king/dbOper/mapCity_db/astar"
	"buffalo/king/king"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"time"
)

type AttackCity struct {
	proto_t
}

func (p *AttackCity) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	var c2s king.C2S_AttackCity
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(2081)
	}
	seelog.Info(c2s.String())

	mapCityIdx := *c2s.MapCityIdx
	armyId := *c2s.ArmyId
	playerId := rpcC2SProto.PlayerId

	mapCityVO, err := mapCity_db.GetMapCityVO(playerId, mapCityIdx)
	if err != nil {
		seelog.WarnWE("GetMapCityVO,err=", err)
		return common.NewErrorCodeRet(2082)
	}
	if mapCityVO == nil {
		return common.NewErrorCodeRet(2083)
	}

	armyVO, err := army_db.GetArmyVO(playerId, armyId)
	if err != nil {
		seelog.WarnWE("GetArmyVO,err=", err)
		return common.NewErrorCodeRet(2084)
	}
	if armyVO == nil {
		return common.NewErrorCodeRet(2085)
	}
	armyVO.UpdateMapCityIdx()
	if !armyVO.BFreeInCity() {
		return common.NewErrorCodeRet(2086)
	}
	// 2018年07月14日 星期六 15时20分11秒 TODO comment it for debug
	//if armyVO.BHasDeadHero() {
	//return common.NewErrorCodeRet(2088)
	//}

	cityVO, err := city_db.GetCityVO(playerId, armyVO.CityId)
	if err != nil {
		return common.NewErrorCodeRet(2089)
	}
	if cityVO == nil {
		return common.NewErrorCodeRet(2090)
	}
	seelog.Tracef("fx=%d,fy=%d,tx=%d,ty=%d", cityVO.X, cityVO.Y, mapCityVO.X, mapCityVO.Y)
	distance, bFound := astar.Distance(cityVO.X, cityVO.Y, mapCityVO.X, mapCityVO.Y)
	if !bFound {
		return common.NewErrorCodeRet(2091)
	}
	needTime := astar.NeedTime(distance)
	seelog.Trace("distance=", distance)

	ct := time.Now().Unix()
	armyVO.TargetTimestamp = int32(ct) + needTime
	armyVO.SourceTimestamp = int32(ct)
	armyVO.MapCityIdx = mapCityIdx
	if err = armyVO.Save(true); err != nil {
		seelog.WarnWE("SaveArmyVO,err=", err)
		return common.NewErrorCodeRet(2087)
	}

	march.AddMarch(playerId, armyId, armyVO.TargetTimestamp, mapCityIdx, needTime)

	var s2c king.S2C_AttackCity
	s2c.ArmyId = proto.Int32(armyId)
	s2c.AttackTimestamp = proto.Int32(armyVO.TargetTimestamp)
	s2c.MarchTimestamp = proto.Int32(armyVO.SourceTimestamp)
	s2c.MapCityIdx = proto.Int32(mapCityIdx)
	return getProtoS2CDataWithoutPush(playerId, "S2C_AttackCity", &s2c)
}*/
