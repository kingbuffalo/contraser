package cfgHeroRecruitMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
)

type CfgHeroRecruit struct {
	//TODO add member here
	Type                   int32 `json:"type"`
	MoneyType              int32 `json:"money_type"`
	MoneyNum               int32 `json:"money_num"`
	Times                  int32 `json:"times"`
	HeroRecruitPoolId      int32 `json:"hero_recruit_pool_id"`
	TenthHeroRecruitPoolId int32 `json:"tenth_hero_recruit_pool_id"`
}

var type_map_cfg map[int32]*CfgHeroRecruit

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_hero_recruit.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgHeroRecruit = make([]*CfgHeroRecruit, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	type_map_cfg = make(map[int32]*CfgHeroRecruit, len(cfgArr))
	for _, v := range cfgArr {
		type_map_cfg[v.Type] = v
	}
}

func GetCfgHeroRecruit(rType int32) *CfgHeroRecruit {
	return type_map_cfg[rType]
}

func RandHeroId(rType, accTimes, tenTimes int32, isAccu bool) ([][]int32, int32, int32) {
	cfg := GetCfgHeroRecruit(rType)
	return randHeroId(cfg.HeroRecruitPoolId, cfg.Times, accTimes, tenTimes, isAccu)
}
