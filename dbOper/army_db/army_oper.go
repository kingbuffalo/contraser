package army_db

import (
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/hero_db"
	"errors"
	"github.com/kingbuffalo/seelog"
)

func (a *ArmyVO) GetHeroVOs() ([]*hero_db.HeroVO, error) {
	ret := make([](*hero_db.HeroVO), 0)
	playerId := a.PlayerId
	heroIds := a.GetHeroIds()
	for _, v := range heroIds {
		if h, err := hero_db.GetHeroVO(playerId, v); err != nil {
			return nil, err
		} else {
			if h != nil {
				ret = append(ret, h)
			}
		}
	}
	if len(ret) == 0 {
		return nil, errors.New("all heroId = 0")
	}
	return ret, nil
}

func (a *ArmyVO) BHasDeadHero() bool {
	return false
}

func (vo *ArmyVO) RemoveHero(heroId int32) {
	if vo.HeroId1 == heroId {
		vo.HeroId1 = 0
		return
	}
	if vo.HeroId2 == heroId {
		vo.HeroId2 = 0
		return
	}
	if vo.HeroId3 == heroId {
		vo.HeroId3 = 0
		return
	}
	if vo.HeroId4 == heroId {
		vo.HeroId4 = 0
		return
	}
	if vo.HeroId5 == heroId {
		vo.HeroId5 = 0
		return
	}
}

func (a *ArmyVO) GetHeroIds() []int32 {
	return []int32{a.HeroId1, a.HeroId2, a.HeroId3, a.HeroId4, a.HeroId5}
}

func (a *ArmyVO) GetHeroId(pos int) int32 {
	if pos == enumErrCode.ARMY_POS_1 {
		return a.HeroId1
	}
	if pos == enumErrCode.ARMY_POS_2 {
		return a.HeroId2
	}
	if pos == enumErrCode.ARMY_POS_3 {
		return a.HeroId3
	}
	if pos == enumErrCode.ARMY_POS_4 {
		return a.HeroId4
	}
	return a.HeroId5
}

func (a *ArmyVO) GetHeroPos(heroId int32) int {
	if heroId == a.HeroId1 {
		return enumErrCode.ARMY_POS_1
	}
	if heroId == a.HeroId2 {
		return enumErrCode.ARMY_POS_2
	}
	if heroId == a.HeroId3 {
		return enumErrCode.ARMY_POS_3
	}
	if heroId == a.HeroId4 {
		return enumErrCode.ARMY_POS_4
	}
	if heroId == a.HeroId5 {
		return enumErrCode.ARMY_POS_5
	}
	return 0
}
func (a *ArmyVO) BEmpty() bool {
	return a.HeroId1 == 0 && a.HeroId2 == 0 && a.HeroId3 == 0 && a.HeroId4 == 0 && a.HeroId5 == 0
}

func (a *ArmyVO) SetHeroId(pos int, heroId int32) {
	if pos == enumErrCode.ARMY_POS_1 {
		a.HeroId1 = heroId
		return
	}
	if pos == enumErrCode.ARMY_POS_2 {
		a.HeroId2 = heroId
		return
	}
	if pos == enumErrCode.ARMY_POS_3 {
		a.HeroId3 = heroId
		return
	}
	if pos == enumErrCode.ARMY_POS_4 {
		a.HeroId4 = heroId
		return
	}
	a.HeroId3 = heroId
}

func NewPlayerArmy(playerId int32) {
	max := enumErrCode.ARMY_COUNT_MAX
	for i := 0; i < max; i++ {
		a := &ArmyVO{
			PlayerId: playerId,
			ArmyId:   int32(i + 1),
		}
		if err := a.SaveNew(); err != nil {
			seelog.WarnWE("army saveNew err=", err)
		}
	}
}
