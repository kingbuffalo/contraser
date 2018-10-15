package cfgHeroMgr

import (
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/startcfg"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CfgHeroSkillEffect struct {
	Id                    int32             `json:"id"`
	SkillType             int32             `json:"skill_type"`
	SkillValue            int32             `json:"skill_value"`
	LastNum               int32             `json:"last_num"`
	TargetIdx             int32             `json:"target_idx"`
	CalcIdTarget          int32             `json:"calc_id_target"`
	CalcIdSource          int32             `json:"calc_id_source"`
	Des                   string            `json:"des"`
	CfgHeroSkillBuffValue *CfgHeroSkillBuff `json:"-"`
}

var id_map_CfgHeroSkillEffect map[int32]*CfgHeroSkillEffect

func initHeroSkillEffects() {
	initHeroSkillBuff()
	path := startcfg.GetCfgPath()
	fn := path + "cfg_hero_skill_eff.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgHeroSkillEffect = make([]*CfgHeroSkillEffect, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	id_map_CfgHeroSkillEffect = make(map[int32]*CfgHeroSkillEffect, len(cfgArr))
	for _, v := range cfgArr {
		if v.SkillType == enumErrCode.SKILL_TYPE_BUFFER {
			cfg := getCfgHeroSkillBuff(v.SkillValue)
			if cfg == nil {
				errMsg := fmt.Sprintf("buff not found,effectId=%d,buffId=%d", v.Id, v.SkillValue)
				panic(errMsg)
			}
			v.CfgHeroSkillBuffValue = cfg
		}
		id_map_CfgHeroSkillEffect[v.Id] = v
	}
}

func getCfgHeroSkillEffect(id int32) *CfgHeroSkillEffect {
	return id_map_CfgHeroSkillEffect[id]
}
