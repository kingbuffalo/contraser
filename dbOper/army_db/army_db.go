package army_db

import (
	"encoding/json"
	//"errors"
	"fmt"
	//"github.com/kingbuffalo/seelog"
	"strconv"
	//"time"
	//"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	//"buffalo/king/dbOper/hero_db"
	//"buffalo/king/dbOper/playerIncr_db"
)

/*
一个玩定有N个地图 (mapId)
一个地图上有>1个城池 ( cityId ) 城池id不会相等，不同地图也不可以相等
一个城池有N个军队
 armyId//以armyId作为唯一标记(这样才方便军队在不同的城市中走动)
*/
type ArmyVO struct {
	PlayerId int32 `json:"player_id" gorm:"primary_key"`
	ArmyId   int32 `json:"army_id" gorm:"primary_key"`

	HeroId1 int32 `json:"hero_id1"`
	HeroId2 int32 `json:"hero_id2"`
	HeroId3 int32 `json:"hero_id3"`
	HeroId4 int32 `json:"hero_id4"`
	HeroId5 int32 `json:"hero_id5"`
}

//func (a *ArmyVO) BFreeInCity() bool {
//if a.MapCityIdx == 0 {
//return true
//}
//ct := int32(time.Now().Unix() + 1)
//return a.SourceTimestamp > a.TargetTimestamp && ct > a.SourceTimestamp
//}

//func (a *ArmyVO) UpdateMapCityIdx() {
//if a.MapCityIdx != 0 {
//if a.BFreeInCity() {
//a.MapCityIdx = 0
//}
//}
//}

func (ArmyVO) TableName() string {
	return "army"
}

func getArmyVOKey(playerId int32) string {
	return fmt.Sprintf("h:sg:ArmyVOKey:%d", playerId)
}

func GetArmyVOs(playerId int32) ([](*ArmyVO), error) {
	key := getArmyVOKey(playerId)
	keyMapJson := gameutil.HGetAll(key)
	if keyMapJson == nil {
		var armyVOs []*ArmyVO = [](*ArmyVO){}
		db := gameutil.GetDB()
		db.Where("player_id=?", playerId).Find(&armyVOs)
		for _, v := range armyVOs {
			if err := v.Save(false); err != nil {
				return nil, err
			}
		}
		return armyVOs, nil
	}
	ret := make([](*ArmyVO), len(keyMapJson))
	idx := 0
	for _, v := range keyMapJson {
		var vo ArmyVO
		if err := json.Unmarshal([]byte(v), &vo); err != nil {
			return nil, err
		}
		ret[idx] = &vo
		idx++
	}
	return ret, nil
}

func GetArmyVO(playerId, armyId int32) (*ArmyVO, error) {
	key := getArmyVOKey(playerId)
	fieldStr := strconv.Itoa(int(armyId))
	b := gameutil.HGet(key, fieldStr)
	if b == nil {
		db := gameutil.GetDB()
		var armyVO ArmyVO
		db.Where("player_id=? AND army_id=?", playerId, armyId).Find(&armyVO)
		if err := armyVO.Save(false); err != nil {
			return nil, err
		}
		return &armyVO, nil
	}
	var vo ArmyVO
	if err := json.Unmarshal(b, &vo); err != nil {
		return nil, err
	}
	return &vo, nil
}

func (vo *ArmyVO) save_helper(bSaveToMysql bool, bNew bool) error {
	key := getArmyVOKey(vo.PlayerId)
	fieldStr := strconv.Itoa(int(vo.ArmyId))
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

func (vo *ArmyVO) Save(bSaveToMysql bool) error {
	return vo.save_helper(bSaveToMysql, false)
}

func (vo *ArmyVO) SaveNew() error {
	return vo.save_helper(true, true)
}
