package hero_db

import (
	"buffalo/king/cfgmgr/cfgHeroMgr"
	//"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"encoding/json"
	"fmt"
	"github.com/kingbuffalo/seelog"
	"strconv"
)

/*
一个玩家不能同时拥有2个相同的武将
*/
type HeroVO struct {
	PlayerId int32 `json:"player_id" gorm:"primary_key"`
	Id       int32 `json:"id" gorm:"primary_key"` //redis Hash fieldKey

	Exp          int32 `json:"exp"`
	Level        int32 `json:"level"`
	Star         int32 `json:"star"`
	ArmyId       int32 `json:"army_id"`
	ArmyLv       int32 `json:"army_lv"`
	ArmySkill1Lv int32 `json:"army_skill1_lv"`
	ArmySkill2Lv int32 `json:"army_skill2_lv"`
	ArmySkill3Lv int32 `json:"army_skill3_lv"`
	Chip         int32 `json:"chip"`
}

func (HeroVO) TableName() string {
	return "hero"
}

func getHeroVOKey(playerId int32) string {
	return fmt.Sprintf("h:king:HeroVOKey:%d", playerId)
}

func GetHeroVOs(playerId int32) ([](*HeroVO), error) {
	key := getHeroVOKey(playerId)
	keyMapJson := gameutil.HGetAll(key)
	if keyMapJson == nil {
		var vos []*HeroVO = [](*HeroVO){}
		db := gameutil.GetDB()
		db.Where("player_id=?", playerId).Find(&vos)
		for _, v := range vos {
			if err := v.Save(false); err != nil {
				return nil, err
			}
		}
		return vos, nil
	}
	ret := make([](*HeroVO), len(keyMapJson))
	idx := 0
	for _, v := range keyMapJson {
		var vo HeroVO
		if err := json.Unmarshal([]byte(v), &vo); err != nil {
			return nil, err
		}
		ret[idx] = &vo
		idx++
	}
	return ret, nil
}

func GetHeroVO(playerId, Id int32) (*HeroVO, error) {
	key := getHeroVOKey(playerId)
	fieldStr := strconv.Itoa(int(Id))
	b := gameutil.HGet(key, fieldStr)
	if b == nil {
		if !gameutil.GExists(key) {
			if vos, err := GetHeroVOs(playerId); err != nil {
				return nil, err
			} else {
				for _, v := range vos {
					if v.Id == Id {
						return v, nil
					}
				}
			}
		}
		return nil, nil
	}
	var vo HeroVO
	if err := json.Unmarshal(b, &vo); err != nil {
		return nil, err
	}
	return &vo, nil
}

func (vo *HeroVO) save_helper(bSaveToMysql bool, bNew bool) error {
	seelog.Trace("save_helper,bSaveToMysql=", bSaveToMysql, ",bNew=", bNew)
	key, field := vo.getRedisKey()
	b, err := json.Marshal(vo)
	if err != nil {
		return err
	}
	gameutil.HSet(key, field, b)
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

func (vo *HeroVO) Save(bSaveToMysql bool) error {
	return vo.save_helper(bSaveToMysql, false)
}

func (vo *HeroVO) NewSave() error {
	return vo.save_helper(true, true)
}

func NewHero(cfgHero *cfgHeroMgr.CfgHero, playerId int32) *HeroVO {
	return &HeroVO{
		PlayerId: playerId,
		Id:       cfgHero.Id,
		Star:     cfgHero.StarLv,
	}
}

func SaveDefaultHeroVOs(heroVOs []*HeroVO) {
	for _, v := range heroVOs {
		if err := v.save_helper(true, true); err != nil {
			seelog.WarnWE("SaveHeroVO,err=", err)
		}
	}
}
