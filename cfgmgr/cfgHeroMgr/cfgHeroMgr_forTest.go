package cfgHeroMgr

/*
import (
	"buffalo/king/cfgmgr/cfgHeroSkillMgr"
	"buffalo/king/common/enumErrCode"
	"fmt"
)

func CreateForTestCfg() {
	cfgHeroArr := [](*CfgHero){
		&CfgHero{
			Id:                  1,
			StarLv:              1,
			Chip:                1,
			Quality:             1,
			Type:                1,
			AttackHeroSkillId:   1,
			AttackCitySkillId:   1,
			AttackHeroSkillList: []int32{},
			AttackCitySkillList: []int32{},
			MaxTroops:           10000,
			AttInit:             100,
			DefInit:             100,
			StrInit:             100,
			SieInit:             100,
			SpeInit:             100,
			AttLvCorr:           100,
			DefLvCorr:           100,
			StrLvCorr:           100,
			SieLvCorr:           100,
			SpeLvCorr:           100,
			MaxTroopsCorr:       100,
		},
	}
	cfgHeroArrLen := len(cfgHeroArr)
	for i := 0; i < cfgHeroArrLen; i++ {
		v := cfgHeroArr[i]
		v.attrs[enumErrCode.HERO_ATTR_IDX_ATT] = v.AttInit
		v.attrs[enumErrCode.HERO_ATTR_IDX_DEF] = v.DefInit
		v.attrs[enumErrCode.HERO_ATTR_IDX_STR] = v.StrInit
		v.attrs[enumErrCode.HERO_ATTR_IDX_SIE] = v.SieInit
		v.attrs[enumErrCode.HERO_ATTR_IDX_SPE] = v.SpeInit
		v.attrs[enumErrCode.HERO_ATTR_IDX_MAX_TROOPS] = v.MaxTroops

		v.attrCorrs[enumErrCode.HERO_ATTR_IDX_ATT] = v.AttLvCorr
		v.attrCorrs[enumErrCode.HERO_ATTR_IDX_DEF] = v.DefLvCorr
		v.attrCorrs[enumErrCode.HERO_ATTR_IDX_STR] = v.StrLvCorr
		v.attrCorrs[enumErrCode.HERO_ATTR_IDX_SIE] = v.SieLvCorr
		v.attrCorrs[enumErrCode.HERO_ATTR_IDX_SPE] = v.SpeLvCorr
		v.attrCorrs[enumErrCode.HERO_ATTR_IDX_MAX_TROOPS] = v.MaxTroopsCorr

		v.ah_cfgHeroSkillArr = make([](*cfgHeroSkillMgr.CfgHeroSkill), len(v.AttackHeroSkillList))
		for i, skillId := range v.AttackHeroSkillList {
			cfgSkill := cfgHeroSkillMgr.GetCfgHeroSkill(skillId)
			if cfgSkill == nil {
				errMsg := fmt.Sprintf("heroId=%d,skillId=%d not found", v.Id, skillId)
				panic(errMsg)
			}
			v.ah_cfgHeroSkillArr[i] = cfgSkill
		}

		v.ac_cfgHeroSkillArr = make([](*cfgHeroSkillMgr.CfgHeroSkill), len(v.AttackCitySkillList))
		for i, skillId := range v.AttackCitySkillList {
			cfgSkill := cfgHeroSkillMgr.GetCfgHeroSkill(skillId)
			if cfgSkill == nil {
				errMsg := fmt.Sprintf("heroId=%d,skillId=%d not found", v.Id, skillId)
				panic(errMsg)
			}
			v.ac_cfgHeroSkillArr[i] = cfgSkill
		}

		v.nh_cfgHeroSkill = cfgHeroSkillMgr.GetCfgHeroSkill(v.AttackHeroSkillId)
		if v.nh_cfgHeroSkill == nil {
			errMsg := fmt.Sprintf("heroId=%d,skillId=%d not found", v.Id, v.AttackHeroSkillId)
			panic(errMsg)
		}
		v.nc_cfgHeroSkill = cfgHeroSkillMgr.GetCfgHeroSkill(v.AttackCitySkillId)
		if v.nc_cfgHeroSkill == nil {
			errMsg := fmt.Sprintf("heroId=%d,skillId=%d not found", v.Id, v.AttackCitySkillId)
			panic(errMsg)
		}

		idMapCfgHero[v.Id] = v
	}
}*/
