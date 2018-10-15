package city_db

/*
import (
	"encoding/json"
	"fmt"
	"strconv"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"buffalo/king/dbOper/army_db"
)

type CityVO struct {
	PlayerId         int32 `json:"player_id" gorm:"primary_key"`
	CityId           int32 `json:"city_id" gorm:"primary_key"`
	ArmyId1          int32 `json:"army_id1"`
	ArmyId2          int32 `json:"army_id2"`
	ArmyId3          int32 `json:"army_id3"`
	ArmyId4          int32 `json:"army_id4"`
	ArmyId5          int32 `json:"army_id5"`
	ArmyId6          int32 `json:"army_id6"`
	BuildingIdx1     int32 `json:"building_idx1"`
	BuildingIdx2     int32 `json:"building_idx2"`
	BuildingIdx3     int32 `json:"building_idx3"`
	BuildingIdx4     int32 `json:"building_idx4"`
	SkinId           int32 `json:"skin_id"`
	MapId            int32 `json:"map_id"`
	DestroyTimestamp int32 `json:"destroy_timestamp"`
	DefenceArmyId    int32 `json:"defence_army_id"`
	Durability       int32 `json:"durability"`
}

func (CityVO) TableName() string {
	return "city"
}

func createDefaultCity(playerId int32) *CityVO {
	cityId := enumErrCode.DEFAULT_CITY_ID
	army_db.CreateDefalutArmyInCity(playerId, 6, cityId)
	mapId := enumErrCode.DEFAULT_CITY_MAP_ID
	return &CityVO{
		PlayerId:      playerId,
		CityId:        cityId,
		ArmyId1:       1,
		ArmyId2:       2,
		ArmyId3:       3,
		ArmyId4:       4,
		ArmyId5:       5,
		ArmyId6:       6,
		BuildingIdx1:  1,
		BuildingIdx2:  2,
		BuildingIdx3:  3,
		BuildingIdx4:  4,
		SkinId:        enumErrCode.DEFAULT_CITY_SKIN_ID,
		MapId:         mapId,
		DefenceArmyId: 1,
	}
}

func getCityVOKey(playerId int32) string {
	return fmt.Sprintf("h:sg:CityVOKey:%d", playerId)
}

func GetCityVOsInOneMap(playerId, mapId int32) ([](*CityVO), error) {
	vos, err := GetCityVOs(playerId)
	if err != nil {
		return nil, err
	}
	ret := make([]*CityVO, 0)
	for _, v := range vos {
		if v.MapId == mapId {
			ret = append(ret, v)
		}
	}
	return ret, nil
}
func GetCityVOs(playerId int32) ([](*CityVO), error) {
	key := getCityVOKey(playerId)
	keyMapJson := gameutil.HGetAll(key)
	if keyMapJson == nil {
		var cityVOs []*CityVO = [](*CityVO){}
		db := gameutil.GetDB()
		db.Where("player_id=?", playerId).Find(&cityVOs)
		if len(cityVOs) == 0 {
			dc := createDefaultCity(playerId)
			if err := dc.save_helper(true, true); err != nil {
				return nil, err
			}
			return [](*CityVO){dc}, nil
		} else {
			for _, v := range cityVOs {
				if err := v.SaveVO(false); err != nil {
					return nil, err
				}
			}
			return cityVOs, nil
		}
	}
	ret := make([](*CityVO), len(keyMapJson))
	idx := 0
	for _, v := range keyMapJson {
		var vo CityVO
		if err := json.Unmarshal([]byte(v), &vo); err != nil {
			return nil, err
		}
		ret[idx] = &vo
		idx++
	}
	return ret, nil
}

func GetCityVO(playerId, cityId int32) (*CityVO, error) {
	key := getCityVOKey(playerId)
	fieldStr := strconv.Itoa(int(cityId))
	b := gameutil.HGet(key, fieldStr)
	if b == nil {
		if !gameutil.GExists(key) {
			if vos, err := GetCityVOs(playerId); err != nil {
				return nil, err
			} else {
				for _, v := range vos {
					if v.CityId == cityId {
						return v, nil
					}
				}
			}
		} else {
			db := gameutil.GetDB()
			var vo CityVO
			db.Where("player_id=? AND city_id =?", playerId, cityId).Find(&vo)
			if vo.PlayerId != 0 {
				err := vo.SaveVO(false)
				return &vo, err
			}
		}
		return nil, nil
	}
	var vo CityVO
	if err := json.Unmarshal(b, &vo); err != nil {
		return nil, err
	}
	return &vo, nil
}

func (vo *CityVO) save_helper(bSaveToMysql bool, bNew bool) error {
	key := getCityVOKey(vo.PlayerId)
	fieldStr := strconv.Itoa(int(vo.CityId))
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

func (vo *CityVO) SaveVO(bSaveToMysql bool) error {
	return vo.save_helper(bSaveToMysql, false)
}*/
