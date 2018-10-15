package city_db

/*
import (
	"buffalo/king/common/gameutil"
	"encoding/json"
	"github.com/kingbuffalo/seelog"
	"strconv"
)

func (v *CityVO) Clone() *CityVO {
	return &CityVO{
		PlayerId:         v.PlayerId,
		CityId:           v.CityId,
		ArmyId1:          v.ArmyId1,
		ArmyId2:          v.ArmyId2,
		ArmyId3:          v.ArmyId3,
		ArmyId4:          v.ArmyId4,
		ArmyId5:          v.ArmyId5,
		ArmyId6:          v.ArmyId6,
		BuildingIdx1:     v.BuildingIdx1,
		BuildingIdx2:     v.BuildingIdx2,
		BuildingIdx3:     v.BuildingIdx3,
		BuildingIdx4:     v.BuildingIdx4,
		SkinId:           v.SkinId,
		MapId:            v.MapId,
		DestroyTimestamp: v.DestroyTimestamp,
		DefenceArmyId:    v.DefenceArmyId,
	}
}

func (v *CityVO) getRedisKey() (string, string) {
	key := getCityVOKey(v.PlayerId)
	fieldStr := strconv.Itoa(int(v.CityId))
	return key, fieldStr
}

func (v *CityVO) DiffSave(oldv *CityVO) error {
	key, field := v.getRedisKey()
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	gameutil.HSet(key, field, b)

	diff := v.getDiff(oldv)
	db := gameutil.GetDB()
	seelog.Trace("Update")
	db.Model(&v).Updates(diff)
	return nil
}

func (v *CityVO) getDiff(oldv *CityVO) *CityVO {
	var ret CityVO

	if v.ArmyId1 != oldv.ArmyId1 {
		ret.ArmyId1 = v.ArmyId1
	}
	if v.ArmyId2 != oldv.ArmyId2 {
		ret.ArmyId2 = v.ArmyId2
	}
	if v.ArmyId3 != oldv.ArmyId3 {
		ret.ArmyId3 = v.ArmyId3
	}
	if v.ArmyId4 != oldv.ArmyId4 {
		ret.ArmyId4 = v.ArmyId4
	}
	if v.ArmyId5 != oldv.ArmyId5 {
		ret.ArmyId5 = v.ArmyId5
	}
	if v.ArmyId6 != oldv.ArmyId6 {
		ret.ArmyId6 = v.ArmyId6
	}
	if v.BuildingIdx1 != oldv.BuildingIdx1 {
		ret.BuildingIdx1 = v.BuildingIdx1
	}
	if v.BuildingIdx2 != oldv.BuildingIdx2 {
		ret.BuildingIdx2 = v.BuildingIdx2
	}
	if v.BuildingIdx3 != oldv.BuildingIdx3 {
		ret.BuildingIdx3 = v.BuildingIdx3
	}
	if v.BuildingIdx4 != oldv.BuildingIdx4 {
		ret.BuildingIdx4 = v.BuildingIdx4
	}
	if v.SkinId != oldv.SkinId {
		ret.SkinId = v.SkinId
	}
	if v.MapId != oldv.MapId {
		ret.MapId = v.MapId
	}
	if v.DestroyTimestamp != oldv.DestroyTimestamp {
		ret.DestroyTimestamp = v.DestroyTimestamp
	}
	if v.DefenceArmyId != oldv.DefenceArmyId {
		ret.DefenceArmyId = v.DefenceArmyId
	}

	return &ret
}*/
