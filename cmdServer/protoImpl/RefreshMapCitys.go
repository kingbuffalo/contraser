package protoImpl

/*

import (
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/city_db"
	"buffalo/king/dbOper/mapCity_db"
	"buffalo/king/dbOper/map_db"
	"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/king"
)

func refreshMapCity(playerId int32, map_id int32) ([]*mapCity_db.MapCityVO, int) {
	var staticIdMapMapCityVOs map[int32](*mapCity_db.MapCityVO)

	armyVOs, err := map_db.GetArmysInOneMap(playerId, map_id)
	if err != nil {
		return nil, 2066
	}
	mapCityIdxs := make([]int32, 0)
	for _, v := range armyVOs {
		if v.MapCityIdx > 0 {
			mapCityIdxs = append(mapCityIdxs, v.MapCityIdx)
		}
	}
	if len(mapCityIdxs) > 0 {
		staticIdMapMapCityVOs = make(map[int32](*mapCity_db.MapCityVO), len(mapCityIdxs))
		for _, v := range mapCityIdxs {
			var err error
			staticIdMapMapCityVOs[v], err = mapCity_db.GetMapCityVO(playerId, v)
			if err != nil {
				return nil, 2067
			}
		}
	}

	mapCityVOs, err := mapCity_db.RefreshMapCityVO(playerId, map_id, staticIdMapMapCityVOs)
	if err != nil {
		seelog.Trace("GetMapCityVOs error")
		return nil, 2062
	}
	seelog.Trace("xxxxxxxxxxxxxxxxxxxxxxxx")
	seelog.Trace(mapCityVOs)
	return mapCityVOs, 0
}

func wrapRefreshMapCityInfo(playerId int32, map_id int32) ([](*king.MapCity), int) {
	mapCityVOs, errCode := refreshMapCity(playerId, map_id)
	if errCode != 0 {
		return nil, errCode
	}
	mapCitys := [](*king.MapCity){}

	for _, mapCityVO := range mapCityVOs {
		//资源获取
		playerInfoVO, err := playerInfo_db.GetPlayerInfoVOBy(mapCityVO.Id)
		if err != nil {
			seelog.Trace("getRobableRes  error, errCode:", 2023)
			return nil, 2063
		}
		robableRes := getRobableRes(playerInfoVO)
		info := getLoadPlayerInfo(playerInfoVO)
		//皮肤获取
		cityVO, err := city_db.GetCityVO(mapCityVO.Id, mapCityVO.CityId)
		if err != nil {
			seelog.Trace("GetCityVO error, errCode:", 2024)
			return nil, 2064
		}
		default_skin := enumErrCode.CONF_DEFAULT_SKIN
		if cityVO != nil {
			default_skin = cityVO.SkinId
		}
		//组合信息
		tmp := &king.MapCity{
			MapCityIdx: proto.Int32(mapCityVO.MapCityIdx),
			CityId:     proto.Int32(mapCityVO.CityId),
			X:          proto.Int32(mapCityVO.X),
			Y:          proto.Int32(mapCityVO.Y),
			Id:         proto.Int32(mapCityVO.Id),
			IdType:     proto.Int32(mapCityVO.IdType),
			Res:        robableRes,
			Info:       info,
			Status:     proto.Int32(mapCityVO.Status),
			SkinId:     proto.Int32(default_skin),
			Durability: proto.Int32(mapCityVO.Durability),
		}
		mapCitys = append(mapCitys, tmp)
	}
	return mapCitys, 0
}

type RefreshMapCitys struct {
	proto_t
}

func (p *RefreshMapCitys) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	var c2s king.C2S_RefreshMapCitys
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		seelog.Trace("proto Unmarshal error!")
		return common.NewErrorCodeRet(2060)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s.String())

	playerId := rpcC2SProto.PlayerId
	map_id := *c2s.MapId

	mapCitys, errCode := wrapRefreshMapCityInfo(playerId, map_id)
	if errCode != 0 {
		return common.NewErrorCodeRet(errCode)
	}
	var s2c king.S2C_RefreshMapCitys
	s2c.MapId = proto.Int32(map_id)
	s2c.MapCitys = mapCitys

	seelog.Trace("ok", mapCitys)
	return getProtoS2CDataWithoutPush(playerId, "S2C_RefreshMapCitys", &s2c)
}*/
