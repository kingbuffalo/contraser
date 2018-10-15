package army_db

import (
	"buffalo/king/common/gameutil"
	"encoding/json"
	"strconv"
)

func (v *ArmyVO) Clone() *ArmyVO {
	return &ArmyVO{
		PlayerId: v.PlayerId,
		ArmyId:   v.ArmyId,
		//		CityId:          v.CityId,
		//		TargetTimestamp: v.TargetTimestamp,
		//		SourceTimestamp: v.SourceTimestamp,
		//		MapCityIdx:      v.MapCityIdx,
		HeroId1: v.HeroId1,
		HeroId2: v.HeroId2,
		HeroId3: v.HeroId3,
	}
}

func (v *ArmyVO) DiffSave(oldv *ArmyVO) error {
	key, field := v.getRedisKey()
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	gameutil.HSet(key, field, b)

	diff := v.getDiff(oldv)
	db := gameutil.GetDB()
	db.Model(v).Updates(diff)
	return nil
}

func (v *ArmyVO) getRedisKey() (string, string) {
	key := getArmyVOKey(v.PlayerId)
	field := strconv.Itoa(int(v.ArmyId))
	return key, field
}

func (v *ArmyVO) getDiff(oldv *ArmyVO) *ArmyVO {
	var ret ArmyVO

	/*
		if v.CityId != oldv.CityId {
			ret.CityId = v.CityId
		}
		if v.TargetTimestamp != oldv.TargetTimestamp {
			ret.TargetTimestamp = v.TargetTimestamp
		}
		if v.SourceTimestamp != oldv.SourceTimestamp {
			ret.SourceTimestamp = v.SourceTimestamp
		}
		if v.MapCityIdx != oldv.MapCityIdx {
			ret.MapCityIdx = v.MapCityIdx
		}
	*/
	if v.HeroId1 != oldv.HeroId1 {
		ret.HeroId1 = v.HeroId1
	}
	if v.HeroId2 != oldv.HeroId2 {
		ret.HeroId2 = v.HeroId2
	}
	if v.HeroId3 != oldv.HeroId3 {
		ret.HeroId3 = v.HeroId3
	}

	return &ret
}
