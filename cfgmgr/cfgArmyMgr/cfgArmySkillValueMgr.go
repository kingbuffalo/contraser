package cfgArmyMgr

import (
	"buffalo/king/common/startcfg"
	"encoding/json"
	"io/ioutil"
)

type CfgArmySkillValue struct {
	Id            int32         `json:"id"`
	Level         int32         `json:"level"`
	Value         int32         `json:"value"`
	Des           string        `json:"des"`
	CfgArmySkillV *CfgArmySkill `json:"-"`
}

var armySkillId_map_lv_map_cfg map[int32](map[int32]*CfgArmySkillValue)

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_army_skill_value.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgArmySkillValue = make([]*CfgArmySkillValue, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	armySkillIdMax := 12
	armySkillId_map_lv_map_cfg = make(map[int32](map[int32]*CfgArmySkillValue), armySkillIdMax)

	for _, v := range cfgArr {
		lv_map_cfg := armySkillId_map_lv_map_cfg[v.Id]
		if lv_map_cfg == nil {
			lv_map_cfg = make(map[int32]*CfgArmySkillValue, 0)
		}
		lv_map_cfg[v.Level] = v
		v.CfgArmySkillV = getCfgArmySkill(v.Id)
		armySkillId_map_lv_map_cfg[v.Id] = lv_map_cfg
	}
}

func (cfg *CfgArmySkillValue) GetValue(targetArmyType, atr int32) int32 {
	if cfg.CfgArmySkillV.ValueAtr == atr {
		bRightType := cfg.CfgArmySkillV.AllTarget == 1
		if !bRightType {
			bRightType = cfg.CfgArmySkillV.bThisArmyType(targetArmyType)
		}
		if bRightType {
			return cfg.Value
		}
	}
	return 0
}

func getCfgArmySkillValue(armySkillId int32, level int32) *CfgArmySkillValue {
	lv_map_cfg := armySkillId_map_lv_map_cfg[armySkillId]
	if lv_map_cfg != nil {
		return lv_map_cfg[level]
	}
	return nil
}
