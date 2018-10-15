package cfgHeroMgr

import (
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/startcfg"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CfgHero struct {
	Id              int32           `json:"id"`
	SkillIds        []int32         `json:"skill_ids"`
	ManSkillId      int32           `json:"skill_ids"`
	ArmyType        int32           `json:"army_type"`
	StarLv          int32           `json:"star_lv"`
	Wuli            int32           `json:"wuli"`
	Zhili           int32           `json:"zhili"`
	Tongshai        int32           `json:"tongshai"`
	Speed           int32           `json:"speed"`
	AtrType         int32           `json:"atr_type"`
	Chip            int32           `json:"chip"`
	CfgHeroSkillArr []*CfgHeroSkill `json:"-"`
	ManCfgHeroSkill *CfgHeroSkill   `json:"-"`
}

func (cfg *CfgHero) GetInitAtr(star, level int32) (int32, int32, int32, int32) {
	w, b, z := cfg.Wuli, cfg.Tongshai, cfg.Zhili
	switch cfg.AtrType {
	case enumErrCode.HERO_ATR_TYPE_YONG:
		w += level * 3
		b += level
		z += level
	case enumErrCode.HERO_ATR_TYPE_JIANG:
		w += level * 2
		b += level * 2
		z += level
	case enumErrCode.HERO_ATR_TYPE_ZHI:
		w += level
		b += level * 2
		z += level * 2
	case enumErrCode.HERO_ATR_TYPE_CI:
		w += level
		b += level
		z += level * 3
	}
	return cfg.Speed + level, w, b, z
}

var heroId_map_CfgHero map[int32]*CfgHero

func init() {
	initHeroSkill()

	path := startcfg.GetCfgPath()
	fn := path + "cfg_hero.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgHero = make([]*CfgHero, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	heroId_map_CfgHero = make(map[int32]*CfgHero, len(cfgArr))
	for _, v := range cfgArr {
		v.CfgHeroSkillArr = make([]*CfgHeroSkill, len(v.SkillIds))
		for ii, skillId := range v.SkillIds {
			cfgSkill := getCfgHeroSkill(skillId)
			if cfgSkill == nil {
				errMsg := fmt.Sprintf("cfgSkill not found--->heroId=%d,skillId=%d", v.Id, skillId)
				panic(errMsg)
			}
			v.CfgHeroSkillArr[ii] = cfgSkill
		}

		if v.ManSkillId != 0 {
			cfgSkill := getCfgHeroSkill(v.ManSkillId)
			if cfgSkill == nil {
				errMsg := fmt.Sprintf("cfgSkill not found--->heroId=%d,skillId=%d", v.Id, v.ManSkillId)
				panic(errMsg)
			}
			v.ManCfgHeroSkill = cfgSkill
		}
		heroId_map_CfgHero[v.Id] = v
	}
}

func GetCfgHero(id int32) *CfgHero {
	return heroId_map_CfgHero[id]
}
