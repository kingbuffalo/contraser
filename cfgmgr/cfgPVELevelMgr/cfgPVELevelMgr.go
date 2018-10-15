package cfgPVELevelMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
)

type CfgPVELevel struct {
	ChapterId         int32        `json:"chapter_id"`
	LevelId           int32        `json:"level_id"`
	PrevLevel         int32        `json:"prev_level"`
	DefenceHeroIds    []int32      `json:"defence_heroIds"`
	DefenceHeroTroops []int32      `json:"defence_heroTroops"`
	DefenceHeroLvs    []int32      `json:"defence_heroLvs"`
	Reward            [][]int32    `json:"reward"`
	ChapterReward     [][]int32    `json:"chapter_reward"`
	Power             int32        `json:"power"`
	InitSkillPoint    int32        `json:"init_skill_point"`
	NextCfg           *CfgPVELevel `json:"-"`
	PrevCfg           *CfgPVELevel `json:"-"`
}

var levelIdMapCfg map[int32]*CfgPVELevel
var firstLevelId int32

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_pve_level.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgPVELevel = make([]*CfgPVELevel, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}

	levelIdMapCfg = make(map[int32]*CfgPVELevel, len(cfgArr))
	for _, v := range cfgArr {
		levelIdMapCfg[v.LevelId] = v
	}
	for k, v := range levelIdMapCfg {
		if v.PrevLevel != 0 {
			levelIdMapCfg[v.PrevLevel].NextCfg = v
			levelIdMapCfg[k].PrevCfg = levelIdMapCfg[v.PrevLevel]
		}
	}
	for _, v := range levelIdMapCfg {
		if v.PrevCfg == nil {
			firstLevelId = v.LevelId
			break
		}
	}
}

func GetFirstLevelId() int32 {
	return firstLevelId
}

func GetCfgPVELevel(levelId int32) *CfgPVELevel {
	return levelIdMapCfg[levelId]
}

func (cfg *CfgPVELevel) IsValid(levelId int32) bool {
	if cfg.NextCfg == nil {
		return true
	}
	if cfg.NextCfg.LevelId == levelId {
		return true
	}
	checkCfg := cfg
	for checkCfg != nil {
		if checkCfg.LevelId == levelId {
			return true
		}
		checkCfg = checkCfg.PrevCfg
	}
	return false
}
