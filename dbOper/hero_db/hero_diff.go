package hero_db

import (
	"buffalo/king/common/gameutil"
	"encoding/json"
	"strconv"
)

func (v *HeroVO) Clone() *HeroVO {
	return &HeroVO{
		//TODO
		PlayerId: v.PlayerId,
		Id:       v.Id,
		Exp:      v.Exp,
		Level:    v.Level,
		Star:     v.Star,
		ArmyId:   v.ArmyId,
		ArmyLv:   v.ArmyLv,
	}
}

func (v *HeroVO) DiffSave(oldv *HeroVO) error {
	key, field := v.getRedisKey()
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	gameutil.HSet(key, field, b)

	diff := v.getDiff(oldv)
	db := gameutil.GetDB()
	db.Model(&v).Updates(diff)
	return nil
}

func (v *HeroVO) getRedisKey() (string, string) {
	key := getHeroVOKey(v.PlayerId)
	field := strconv.Itoa(int(v.Id))
	return key, field
}

func (v *HeroVO) getDiff(oldv *HeroVO) *HeroVO {
	var ret HeroVO
	ret.PlayerId = v.PlayerId
	ret.Id = v.Id

	if v.Exp != oldv.Exp {
		ret.Exp = v.Exp
	}
	if v.Level != oldv.Level {
		ret.Level = v.Level
	}
	if v.Star != oldv.Star {
		ret.Star = v.Star
	}
	if v.ArmyId != oldv.ArmyId {
		ret.ArmyId = v.ArmyId
	}

	if v.ArmyLv != oldv.ArmyLv {
		ret.ArmyLv = v.ArmyLv
	}

	return &ret
}
