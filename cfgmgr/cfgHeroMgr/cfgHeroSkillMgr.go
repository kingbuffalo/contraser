package cfgHeroMgr

import (
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/startcfg"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CfgHeroSkill struct {
	Id                  int32                 `json:"id"`
	TriggerTime         int32                 `json:"trigger_time"`
	TriggerTarget       int32                 `json:"trigger_target"`
	SkillPos            int32                 `json:"skill_pos"`
	SkillPoint          int32                 `json:"skill_point"`
	EffectIds           []int32               `json:"effect_ids"`
	Continue            int32                 `json:"continue"`
	CfgHeroSkillEffects []*CfgHeroSkillEffect `json:"-"`
}

var skillId_map_CfgHeroSkill map[int32]*CfgHeroSkill

func initHeroSkill() {
	initHeroSkillEffects()
	fmt.Println("initHeroSkill")
	path := startcfg.GetCfgPath()
	fn := path + "cfg_hero_skill.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgHeroSkill = make([]*CfgHeroSkill, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	skillId_map_CfgHeroSkill = make(map[int32]*CfgHeroSkill, len(cfgArr))
	for _, v := range cfgArr {
		v.CfgHeroSkillEffects = make([]*CfgHeroSkillEffect, len(v.EffectIds))
		for ii, eid := range v.EffectIds {
			cfg := getCfgHeroSkillEffect(eid)
			if cfg == nil {
				errMsg := fmt.Sprintf("cfgHeroSkillEffect not found,skillId=%d,effectId=%d", v.Id, eid)
				panic(errMsg)
			}
			v.CfgHeroSkillEffects[ii] = cfg
		}
		skillId_map_CfgHeroSkill[v.Id] = v
	}
}

func getCfgHeroSkill(id int32) *CfgHeroSkill {
	return skillId_map_CfgHeroSkill[id]
}

func Gotoplatform_CfgHeroSkill() *CfgHeroSkill {
	return getCfgHeroSkill(enumErrCode.SKILL_ID_GOTO_PLATFORM)
}

func Leftplatform_CfgHeroSkill() *CfgHeroSkill {
	return getCfgHeroSkill(enumErrCode.SKILL_ID_LEFT_PLATFORM)
}

func NormalATK_CfgHeroSkill() *CfgHeroSkill {
	return getCfgHeroSkill(enumErrCode.SKILL_ID_ATTACK)
}
