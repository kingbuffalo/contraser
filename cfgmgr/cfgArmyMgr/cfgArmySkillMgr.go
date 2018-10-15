package cfgArmyMgr

import (
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/startcfg"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CfgArmySkill struct {
	Id          int32  `json:"id"`
	ArmyType    int32  `json:"army_type"`
	ValueAtr    int32  `json:"value_atr"`
	AllTarget   int32  `json:"all_target"`
	OppoCavaLry int32  `json:"oppo_cavalry"`
	OppoPikeman int32  `json:"oppo_pikeman"`
	OppoShields int32  `json:"oppo_shields"`
	OppoArchers int32  `json:"oppo_archers"`
	Name        string `json:"name"`
	Des         string `json:"des"`
	OpenLevel   int32  `json:"open_level"`
}

var armyType_map_cfgArmySkillArr map[int32]([]*CfgArmySkill)
var armySkillId_map_cfgArmySkill map[int32](*CfgArmySkill)

func init() {
	path := startcfg.GetCfgPath()
	fn := path + "cfg_army_skill.json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	var cfgArr []*CfgArmySkill = make([]*CfgArmySkill, 0)
	if err = json.Unmarshal(b, &cfgArr); err != nil {
		panic(err)
	}
	armyType_map_cfgArmySkillArr = make(map[int32]([]*CfgArmySkill), 0)
	armySkillId_map_cfgArmySkill = make(map[int32]*CfgArmySkill, len(cfgArr))
	for _, v := range cfgArr {
		arr := armyType_map_cfgArmySkillArr[v.ArmyType]
		if arr == nil {
			arr = make([]*CfgArmySkill, 0)
		}
		arr = append(arr, v)
		armyType_map_cfgArmySkillArr[v.ArmyType] = arr
		armySkillId_map_cfgArmySkill[v.Id] = v
	}
	checkArmySkill(enumErrCode.ARMY_TY_CAVALRY)
	checkArmySkill(enumErrCode.ARMY_TY_PIKEMAN)
	checkArmySkill(enumErrCode.ARMY_TY_SHIELDS)
	checkArmySkill(enumErrCode.ARMY_TY_ARCHERS)
}

func (cfg *CfgArmySkill) bThisArmyType(targetArmyType int32) bool {
	switch targetArmyType {
	case enumErrCode.ARMY_TY_CAVALRY:
		return cfg.OppoCavaLry == 1
	case enumErrCode.ARMY_TY_PIKEMAN:
		return cfg.OppoPikeman == 1
	case enumErrCode.ARMY_TY_SHIELDS:
		return cfg.OppoShields == 1
	case enumErrCode.ARMY_TY_ARCHERS:
		return cfg.OppoArchers == 1
	}
	return false
}

func checkArmySkill(armyType int32) {
	arr := GetAllCfgArmySkill(armyType)
	if arr == nil {
		errMsg := fmt.Sprintf("army skill not found,armyType=%d", armyType)
		panic(errMsg)
	}
	if len(arr) != int(enumErrCode.ARMY_SKILL_NUM) {
		errMsg := fmt.Sprintf("army skill not found,armyType=%d", armyType)
		panic(errMsg)
	}
}

func getCfgArmySkill(id int32) *CfgArmySkill {
	return armySkillId_map_cfgArmySkill[id]
}

func GetAllCfgArmySkill(armyType int32) []*CfgArmySkill {
	return armyType_map_cfgArmySkillArr[armyType]
}

func GetArmySkillId(armyType int32, levelIdx int32) int32 {
	return (armyType-1)*enumErrCode.ARMY_SKILL_NUM + levelIdx
}

func GetAllCfgArmySkillValue(armyType, lv1, lv2, lv3 int32) []*CfgArmySkillValue {
	//TODO 待优化
	var ret []*CfgArmySkillValue = make([]*CfgArmySkillValue, 4)
	id := GetArmySkillId(armyType, 1)
	ret[0] = getCfgArmySkillValue(id, lv1)
	id = GetArmySkillId(armyType, 2)
	ret[1] = getCfgArmySkillValue(id, lv2)
	id = GetArmySkillId(armyType, 3)
	ret[2] = getCfgArmySkillValue(id, lv3)
	return ret
}
