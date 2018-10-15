package cfgHeroMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
)

type CfgHeroSkillBuff struct {
	//TODO add member here
	Id        int32  `json:"id"`
	BuffType  int32  `json:"buff_type"`
	BuffValue int32  `json:"buff_value"`
	LastNum   int32  `json:"last_num"`
	EffectAtr int32  `json:"effect_atr"`
	CalcId    int32  `json:"calc_id"`
	Name      string `json:"name"`
	Des       string `json:"des"`
}

var id_map_CfgHeroSkillBuff map[int32]*CfgHeroSkillBuff

func initHeroSkillBuff() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_hero_skill_buff.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgHeroSkillBuff = make([]*CfgHeroSkillBuff, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}

	id_map_CfgHeroSkillBuff = make(map[int32]*CfgHeroSkillBuff, len(cfgArr))
	for _, v := range cfgArr {
		id_map_CfgHeroSkillBuff[v.Id] = v
	}
}

func getCfgHeroSkillBuff(id int32) *CfgHeroSkillBuff {
	return id_map_CfgHeroSkillBuff[id]
}
