package playerIncr_db

import (
	"encoding/json"
	"fmt"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
)

type PlayerIncr struct {
	PlayerId   int32 `json:"player_id"`
	ArmyId     int32 `json:"army_id"`
	BuildIdx   int32 `json:"build_idx"`
	CityId     int32 `json:"city_id"`
	MapId      int32 `json:"map_id"`
	MapCityIdx int32 `json:"map_city_idx"`
}

//func (PlayerIncr) TableName() string {
//	return ""cityyer_incr
//}
func getKey(playerId int32) string {
	return fmt.Sprintf("h:sg:playerIncr:%d", playerId)
}

func newPlayerIncr(playerId int32) *PlayerIncr {
	return &PlayerIncr{}
}

func IncArmyId(playerId int32) (int32, error) {
	key := getKey(playerId)
	field := "ArmyId"
	ret := gameutil.HIncrBy(key, field, 1)
	return int32(ret), nil
}

func IncBattleResultId(playerId int32) int32 {
	key := getKey(playerId)
	field := "BattleResultId"
	var battleResultIdStart int64 = 1
	ret := gameutil.HIncrBy(key, field, battleResultIdStart)
	if int(ret) > enumErrCode.BATTLE_RESULT_IDX_MAX {
		gameutil.HSetAllType(key, field, battleResultIdStart)
		ret = 1
	}
	return int32(ret)
}

//func Get(playerId int32) (*PlayerIncr, error) {
//key := getKey(playerId)
//if rep == nil {
////TODO query from mysql
//playerIncr := newPlayerIncr(playerId)
//if err := playerIncr.Save(false); err != nil {
//return nil, err
//}
//}

//var playerIncr PlayerIncr
//playerIncr.PlayerId  = playerId
//playerIncr.ArmyId = playerId
//playerIncr.PlayerId  = playerId
//playerIncr.PlayerId  = playerId
//playerIncr.PlayerId  = playerId

//PlayerId   int32 `json:"player_id"`
//ArmyId     int32 `json:"army_id"`
//BuildIdx   int32 `json:"build_idx"`
//CityId     int32 `json:"city_id"`
//MapId      int32 `json:"map_id"`
//MapCityIdx int32 `json:"map_city_idx"`
//err := json.Unmarshal(rep, &playerIncr)
//if err != nil {
//return nil, err
//}
//return &playerIncr, nil
//}

func (playerIncr *PlayerIncr) Save(bSaveToMysql bool) error {
	key := getKey(playerIncr.PlayerId)
	b, err := json.Marshal(playerIncr)
	if err != nil {
		return err
	}
	gameutil.GSet(key, b)

	if bSaveToMysql {
		//TODO
	}

	return nil
}
