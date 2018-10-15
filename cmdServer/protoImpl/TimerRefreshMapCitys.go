package protoImpl

/*

import (
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common"
	"buffalo/king/dbOper/mapCity_db"
	"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/king"
)

type TimerRefreshMapCitys struct {
	proto_t
}

func getPortMapCityInfo(player_id int32, map_id int32) ([](*king.TimerMapCity), int) {
	mapCityVOs, err := mapCity_db.GetMapCityVOs(player_id, map_id)
	if err != nil {
		seelog.Trace("GetMapCityVOs error, err_code, 2121")
		return nil, 2121
	}
	mapCitys := [](*king.TimerMapCity){}
	for _, mapCity := range mapCityVOs {
		seelog.Trace("ok coming", mapCity.MapId)
		if mapCity.MapId == map_id {

			playerInfoVO, err := playerInfo_db.GetPlayerInfoVOBy(player_id)
			if err != nil {
				seelog.Trace("GetPlayerInfoVOBy error, errCode:", 2023)
				return nil, 2122
			}

			robableRes := getRobableRes(playerInfoVO) //也要获取
			tmp := &king.TimerMapCity{
				MapCityIdx: proto.Int32(mapCity.MapCityIdx),
				Res:        robableRes,
				Status:     proto.Int32(mapCity.Status),
			}
			//tmp.status
			mapCitys = append(mapCitys, tmp)
		}
	}
	return mapCitys, 0
}

func (p *TimerRefreshMapCitys) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	var c2s king.C2S_TimerRefreshMapCitys
	if err := proto.Unmarshal(rpcC2SProto.Data, &c2s); err != nil {
		return common.NewErrorCodeRet(2120)
	}
	seelog.Info(c2s.String())
	playerId := rpcC2SProto.PlayerId
	map_id := *c2s.MapId

	var s2c king.S2C_TimerRefreshMapCitys
	s2c.MapId = proto.Int32(map_id)
	mapCitys, err_code := getPortMapCityInfo(playerId, map_id)
	if err_code != 0 {
		return common.NewErrorCodeRet(err_code)
	}
	s2c.MapCitys = mapCitys
	seelog.Trace("ok", mapCitys)
	return getProtoS2CDataWithoutPush(playerId, "S2C_TimerRefreshMapCitys", &s2c)
}*/
