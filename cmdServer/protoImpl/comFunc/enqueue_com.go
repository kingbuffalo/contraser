package comFunc

import (
	"github.com/kingbuffalo/seelog"
	"buffalo/king/dbOper/army_db"
	"buffalo/king/dbOper/hero_db"
)

func EnQueueSetPos(armyVOs *[](*army_db.ArmyVO), heroVOs *[]*hero_db.HeroVO, playerId, heroId, armyId int32, pos int, getArmyErr, getHeroErr int) int {
	armyVO, errInt := getArmyVOInCache(armyVOs, playerId, armyId, getArmyErr)
	if errInt != 0 {
		return errInt
	}
	ori_heroId := armyVO.GetHeroId(pos)
	seelog.Trace("GetHeroId", pos, ori_heroId)
	if ori_heroId != 0 {
		ori_heroVO, errInt := getHeroVOInCache(heroVOs, playerId, ori_heroId, getHeroErr)
		if errInt != 0 {
			return errInt
		}
		ori_heroVO.ArmyId = 0
	}
	armyVO.SetHeroId(pos, heroId)
	if heroId == 0 {
		return 0
	}
	heroVO, errInt := getHeroVOInCache(heroVOs, playerId, heroId, getHeroErr)
	if errInt != 0 {
		return errInt
	}
	if heroVO.ArmyId != 0 {
		ori_armyVO, errInt := getArmyVOInCache(armyVOs, playerId, heroVO.ArmyId, getArmyErr)
		if errInt != 0 {
			return errInt
		}
		ori_armyVO.RemoveHero(heroId)
	}
	heroVO.ArmyId = armyId
	return 0
}

func getArmyVOInCache(armyVOs *[]*army_db.ArmyVO, playerId, armyId int32, getArmyErr int) (*army_db.ArmyVO, int) {
	for _, v := range *armyVOs {
		if v.PlayerId == playerId && v.ArmyId == armyId {
			return v, 0
		}
	}
	armyVO, err := army_db.GetArmyVO(playerId, armyId)
	if err != nil {
		seelog.WarnWE("GetArmyVO,err=", err)
		//return nil, 4035
		return nil, getArmyErr
	}
	if armyVO == nil {
		//return nil, 4041
		return nil, getArmyErr
	}
	*armyVOs = append(*armyVOs, armyVO)
	return armyVO, 0

}
func getHeroVOInCache(heroVOs *[]*hero_db.HeroVO, playerId, heroId int32, getHeroErr int) (*hero_db.HeroVO, int) {
	for _, v := range *heroVOs {
		if v.PlayerId == playerId && v.Id == heroId {
			return v, 0
		}
	}

	heroVO, err := hero_db.GetHeroVO(playerId, heroId)
	if err != nil {
		//return nil, 4032
		return nil, getHeroErr
	}
	if heroVO == nil {
		//return nil, 4033
		return nil, getHeroErr
	}
	*heroVOs = append(*heroVOs, heroVO)
	return heroVO, 0
}

func EnQueueSaveHeroArmy(armyVOs *[](*army_db.ArmyVO), heroVOs *[]*hero_db.HeroVO, armySaveErr, heroSaveErr int) int {
	for _, v := range *heroVOs {
		if err := v.Save(true); err != nil {
			seelog.WarnWE("heroVO save err=", err)
			return heroSaveErr
		}
	}
	for _, v := range *armyVOs {
		if err := v.Save(true); err != nil {
			seelog.WarnWE("armyVO save err = ", err)
			return armySaveErr
		}
	}
	return 0
}
