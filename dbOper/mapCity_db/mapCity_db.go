package mapCity_db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kingbuffalo/seelog"
	"strconv"
	"buffalo/king/cfgmgr/cfgMapCityMgr"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	//"buffalo/king/dbOper/army_db"
	"buffalo/king/dbOper/city_db"
	"buffalo/king/dbOper/playerInfo_db"
)

/*
此数据不用落地
放redis就可以
一个玩定有N个地图 (mapId)
一个地图上有>1个npc（还没做） 也有1>个玩家。(自己的城池不在内)
这个目前只做玩家的
*/
type MapCityVO struct {
	PlayerId   int32 `json:"player_id"`
	MapCityIdx int32 `json:"map_city_idx"` // 按照个数递增

	MapId      int32 `json:"map_id"`
	CityId     int32 `json:"city_id"` //cityId or npcIdx
	IdType     int32 `json:"id_type"`
	Id         int32 `json:"id"`
	X          int32 `json:"x"`
	Y          int32 `json:"y"`
	Status     int32 `json:"status"`
	Durability int32 `json:"durability"`
}

func (MapCityVO) TableName() string {
	return "map_city"
}

func newMapCity(cityPoint *cfgMapCityMgr.CityPoint, playerId, mapCityIndex, idType, id, cityId, mapId, durability int32) MapCityVO {
	return MapCityVO{
		PlayerId:   playerId,
		MapCityIdx: mapCityIndex,
		MapId:      mapId,
		X:          cityPoint.PointX,
		Y:          cityPoint.PointY,
		IdType:     idType,
		Id:         id,
		CityId:     cityId,
		Status:     enumErrCode.MAP_CITY_STATUS_NORMAL,
		Durability: durability,
	}
}

func GetCityIdByMapId(playerId int32, mapId int32) ([](*city_db.CityVO), error) {
	//TODO move to city_db
	var ret [](*city_db.CityVO)
	cityVO, err := city_db.GetCityVOs(playerId)
	if err != nil {
		seelog.Trace("GetCityVOs error, player, map_id ", playerId, mapId)
		return nil, err
	}
	for _, cityInfo := range cityVO {
		if cityInfo.MapId == mapId {
			ret = append(ret, cityInfo)
		}
	}
	if len(ret) <= 0 {
		panic("this map do not exist the city")
	}

	return ret, nil
}

func RefreshMapCityVO(playerId, mapId int32, staticIdMapMapCityVOs map[int32]*MapCityVO) ([](*MapCityVO), error) {
	return createMapCityInfo(playerId, enumErrCode.TOTAL_NEIGHBOR, mapId, staticIdMapMapCityVOs)
}

func createMapCityInfo(playerId, neighborNum, mapId int32, staticIdMapMapCityVOs map[int32]*MapCityVO) ([](*MapCityVO), error) {
	seelog.Trace("createMapCityInfo: playerId,neighborNum,mapId,refresh", playerId, neighborNum, mapId)
	neighborsPoint, err := cfgMapCityMgr.Generate(mapId, neighborNum)
	if err != nil {
		seelog.WarnWE("Generate nil:", playerId, neighborNum)
		return nil, err
	}
	pidArr, err := playerInfo_db.GetRandPlayerIds(neighborNum, playerId)
	if len(pidArr) != int(neighborNum) {
		panic("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	}
	if err != nil {
		seelog.WarnWE("GetRandPlayerInfo nil:", playerId, neighborNum)
		return nil, err
	}

	mapCityVOs := make([](*MapCityVO), neighborNum)

	for i := 0; i < int(neighborNum); i++ {
		mapCityIdx := int32(i) + mapId*100
		bSet := false
		if staticIdMapMapCityVOs != nil {
			oldMapCityVO, ok := staticIdMapMapCityVOs[mapCityIdx]
			if ok {
				mapCityVOs[i] = oldMapCityVO
				bSet = true
			}
		}
		if !bSet {
			pid := pidArr[i]
			cityVOs, err := city_db.GetCityVOsInOneMap(pid, mapId)
			if err != nil {
				return nil, err
			}
			if len(cityVOs) == 0 {
				errMsg := fmt.Sprintf("city len = 0,playerId=%d,mapId=%d", pid, mapId)
				seelog.WarnWE(errMsg)
				return nil, errors.New(errMsg)
			}
			cityVO := cityVOs[0]
			tmp := newMapCity(neighborsPoint[i], playerId, mapCityIdx, 0, pid, cityVO.CityId, mapId, cityVO.Durability)
			if err := tmp.Save(); err != nil {
				return nil, err
			}
			mapCityVOs[i] = &tmp
		}
	}
	return mapCityVOs, nil
}

func getMapCityVOKey(playerId int32) string {
	return fmt.Sprintf("h:sg:MapCityVOKey:%d", playerId)
}

func GetMapCityVOs(playerId int32, mapId int32) ([](*MapCityVO), error) {
	seelog.Trace("playerId:", playerId)
	key := getMapCityVOKey(playerId)
	keyMapJson := gameutil.HGetAll(key)
	if keyMapJson == nil {
		seelog.Trace("keyMapJson  is nil", playerId)
		return createMapCityInfo(playerId, enumErrCode.TOTAL_NEIGHBOR, mapId, nil)
	}
	seelog.Trace("keyMapJson  is not nil, len:", len(keyMapJson))
	ret := make([](*MapCityVO), len(keyMapJson))
	idx := 0
	for _, v := range keyMapJson {
		var vo MapCityVO
		if err := json.Unmarshal([]byte(v), &vo); err != nil {
			return nil, err
		}
		ret[idx] = &vo
		idx++
	}
	return ret, nil
}

func GetMapCityVO(playerId, mapCityIdx int32) (*MapCityVO, error) {
	key := getMapCityVOKey(playerId)
	fieldStr := strconv.Itoa(int(mapCityIdx))
	b := gameutil.HGet(key, fieldStr)
	if b == nil {
		return nil, nil
	}
	var vo MapCityVO
	if err := json.Unmarshal(b, &vo); err != nil {
		return nil, err
	}
	return &vo, nil
}

func (vo *MapCityVO) Save() error {
	key := getMapCityVOKey(vo.PlayerId)
	fieldStr := strconv.Itoa(int(vo.MapCityIdx))
	b, err := json.Marshal(vo)
	if err != nil {
		return err
	}
	gameutil.HSet(key, fieldStr, b)
	return nil
}
