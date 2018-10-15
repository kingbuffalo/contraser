package protoImpl

/*

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/dbOper/city_db"
	"buffalo/king/dbOper/mapCity_db"
	"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/king"
)

func getRobableRes(playerInfoVO *playerInfo_db.PlayerInfoVO) *king.RobableRes {
	return &king.RobableRes{
		Wood:  proto.Int32(playerInfoVO.Wood),
		Gold:  proto.Int32(playerInfoVO.Gold),
		Grain: proto.Int32(playerInfoVO.Grain),
	}
}
func getLoadPlayerInfo(playerInfoVO *playerInfo_db.PlayerInfoVO) *king.LordPlayerInfo {
	name, err := base64.StdEncoding.DecodeString(playerInfoVO.Name)
	if err != nil {
		seelog.WarnWE("base64Decode error=", err)
	}
	return &king.LordPlayerInfo{
		Name:  proto.String(string(name)),
		Url:   proto.String(playerInfoVO.Url),
		Level: proto.Int32(playerInfoVO.Level), //TODO
	}
}

func getMapCityInfo(playerId int32, map_id int32) ([](*king.MapCity), int) {
	mapCityVOs, err := mapCity_db.GetMapCityVOs(playerId, map_id)
	if err != nil {
		seelog.Trace("GetMapCityVOs error")
		return nil, 2022
	}
	//不确定长度所以要
	mapCitys := [](*king.MapCity){}
	for _, mapCity := range mapCityVOs {
		seelog.Trace("ok coming", mapCity.MapId)
		if mapCity.MapId == map_id {
			playerInfoVO, err := playerInfo_db.GetPlayerInfoVOBy(mapCity.Id)
			if err != nil {
				seelog.Trace("GetPlayerInfoVOBy error, errCode:", 2023)
				return nil, 2023
			}
			robableRes := getRobableRes(playerInfoVO)
			info := getLoadPlayerInfo(playerInfoVO)
			//状态和皮肤即时获取
			cityVO, err := city_db.GetCityVO(mapCity.Id, mapCity.CityId)
			if err != nil {
				seelog.Trace("GetCityVO error, errCode:", 2024)
				return nil, 2024
			}
			tmp := &king.MapCity{
				MapCityIdx: proto.Int32(mapCity.MapCityIdx),
				CityId:     proto.Int32(mapCity.CityId),
				X:          proto.Int32(mapCity.X),
				Y:          proto.Int32(mapCity.Y),
				Id:         proto.Int32(mapCity.Id),
				IdType:     proto.Int32(mapCity.IdType),
				Res:        robableRes,
				Info:       info,
				Status:     proto.Int32(mapCity.Status),
				SkinId:     proto.Int32(cityVO.SkinId),
				Durability: proto.Int32(mapCity.Durability),
			}
			//tmp.status
			mapCitys = append(mapCitys, tmp)
		}
	}
	return mapCitys, 0
}

type GetMapCitys struct {
	proto_t
}

func (g *GetMapCitys) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if g.bClose {
		return common.NewErrorCodeRet(107)
	}
	var c2s_GetMapCitys king.C2S_GetMapCitys
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s_GetMapCitys); err != nil {
		seelog.Trace("proto Unmarshal error!")
		return common.NewErrorCodeRet(2021)
	}
	seelog.Info(rpcC2SProto.PlayerId, ",", c2s_GetMapCitys.String())
	playerId := rpcC2SProto.PlayerId
	map_id := *c2s_GetMapCitys.MapId

	var s2c_GetMapCitys king.S2C_GetMapCitys
	s2c_GetMapCitys.MapId = proto.Int32(map_id)
	mapCitys, errCode := getMapCityInfo(playerId, map_id)
	if errCode != 0 {
		return common.NewErrorCodeRet(errCode)
	}
	s2c_GetMapCitys.MapCitys = mapCitys
	seelog.Trace("ok", mapCitys)
	return getProtoS2CDataWithoutPush(playerId, "S2C_GetMapCitys", &s2c_GetMapCitys)
}*/
