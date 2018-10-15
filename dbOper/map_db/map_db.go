package map_db

/*
一个玩定有N个地图 (mapId)
一个地图上有>1个城池 ( cityId ) 城池id不会相等，不同地图也不可以相等
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	//"github.com/kingbuffalo/seelog"
	"strconv"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"buffalo/king/dbOper/army_db"
)

type MapVO struct {
	PlayerId int32 `json:"player_id" gorm:"primary_key"`
	MapId    int32 `json:"map_id" gorm:"primary_key"`

	CityId1 int32 `json:city_id1`
	CityId2 int32 `json:city_id2`
	CityId3 int32 `json:city_id3`
	CityId4 int32 `json:city_id4`
	CityId5 int32 `json:city_id5`
}

func (MapVO) TableName() string {
	return "map"
}

func getMapVOKey(playerId int32) string {
	return fmt.Sprintf("h:sg:MapVOKey:%d", playerId)
}

func (vo *MapVO) bInThisMap(cityId int32) bool {
	return cityId == vo.CityId1 || cityId == vo.CityId2 || cityId == vo.CityId3 || cityId == vo.CityId4 || cityId == vo.CityId5
}

func GetArmysInOneMap(playerId, mapId int32) ([](*army_db.ArmyVO), error) {
	mapVO, err := GetMapVO(playerId, mapId)
	if err != nil {
		return nil, err
	}
	if mapVO == nil {
		return nil, errors.New("mapVO is not exist")
	}

	armyVOs, err := army_db.GetArmyVOs(playerId)
	ret := make([](*army_db.ArmyVO), 0)
	for _, v := range armyVOs {
		if mapVO.bInThisMap(v.CityId) {
			ret = append(ret, v)
		}
	}
	return ret, nil
}

func GetMapVOs(playerId int32) ([](*MapVO), error) {
	key := getMapVOKey(playerId)
	keyMapJson := gameutil.HGetAll(key)
	if keyMapJson == nil {
		//TODO add mysql oper
		var mapVOs []*MapVO = [](*MapVO){}
		db := gameutil.GetDB()
		db.Where("player_id=?", playerId).Find(&mapVOs)
		if len(mapVOs) == 0 {
			dm := createDefaultMap(playerId, enumErrCode.DEFAULT_CITY_MAP_ID)
			if err := dm.save_helper(true, true); err != nil {
				return nil, err
			}
			return [](*MapVO){dm}, nil
		} else {
			for _, v := range mapVOs {
				if err := v.SaveVO(false); err != nil {
					return nil, err
				}
			}
			return mapVOs, nil
		}
	}
	ret := make([](*MapVO), len(keyMapJson))
	idx := 0
	for _, v := range keyMapJson {
		var vo MapVO
		if err := json.Unmarshal([]byte(v), &vo); err != nil {
			return nil, err
		}
		ret[idx] = &vo
		idx++
	}
	return ret, nil
}

func createDefaultMap(playerId, mapId int32) *MapVO {
	return &MapVO{
		PlayerId: playerId,
		MapId:    mapId,

		CityId1: enumErrCode.DEFAULT_CITY_ID,
		CityId2: 0,
		CityId3: 0,
		CityId4: 0,
		CityId5: 0,
	}

}

func GetMapVO(playerId, mapId int32) (*MapVO, error) {
	key := getMapVOKey(playerId)
	fieldStr := strconv.Itoa(int(mapId))
	b := gameutil.HGet(key, fieldStr)
	if b == nil {
		ret := createDefaultMap(playerId, mapId)
		err := ret.SaveVO(false)
		return ret, err
	}
	var vo MapVO
	if err := json.Unmarshal(b, &vo); err != nil {
		return nil, err
	}
	return &vo, nil
}

func (vo *MapVO) save_helper(bSaveToMysql bool, bNew bool) error {
	key := getMapVOKey(vo.PlayerId)
	fieldStr := strconv.Itoa(int(vo.MapId))
	b, err := json.Marshal(vo)
	if err != nil {
		return err
	}
	gameutil.HSet(key, fieldStr, b)
	if bSaveToMysql {
		db := gameutil.GetDB()
		if bNew {
			db.Create(vo)
		} else {
			db.Save(vo)
		}
	}
	return nil
}

func (vo *MapVO) SaveVO(bSaveToMysql bool) error {
	return vo.save_helper(bSaveToMysql, false)
}
