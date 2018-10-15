package comFunc

/*
import (
	"buffalo/king/cfgmgr/cfgHeroMgr"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/hero_db"
	"buffalo/king/dbOper/playerInfo_db"
	"github.com/kingbuffalo/seelog"
)

type AddLocIdCount struct {
	PlayerId          int32
	PlayerInfoVOV     *playerInfo_db.PlayerInfoVO
	Old_PlayerInfoVOV *playerInfo_db.PlayerInfoVO
	HeroVOs           []*hero_db.HeroVO
	Old_HeroVOs       []*hero_db.HeroVO
}

func NewAddLocIdCount(playerId int32, playerInfoVO *playerInfo_db.PlayerInfoVO, heroVOs []*hero_db.HeroVO) *AddLocIdCount {
	return &AddLocIdCount{
		PlayerId:      playerId,
		PlayerInfoVOV: playerInfoVO,
		HeroVOs:       heroVOs,
	}
}

func (a *AddLocIdCount) getHeroVO(heroId int32) (*hero_db.HeroVO, error) {
	for _, v := range a.HeroVOs {
		if v.Id == heroId {
			return v, nil
		}
	}
	heroVO, err := hero_db.GetHeroVO(a.PlayerId, heroId)
	if err != nil {
		return nil, err
	}
	if heroVO != nil {
		a.HeroVOs = append(a.HeroVOs, heroVO)
	}
	return heroVO, nil
}

func (a *AddLocIdCount) addSingleHelp(loc, id, co int32) error {
	switch loc {
	case enumErrCode.LOC_HERO:
		h, err := a.getHeroVO(id)
		if err != nil {
			return err
		}
		cfg := cfgHeroMgr.GetCfgHero(id)
		if h == nil {
			h = hero_db.NewHero(cfg, a.PlayerId)
			a.HeroVOs = append(a.HeroVOs, h)
		} else {
			h.Chip += cfg.Chip
		}
	}
	return nil
}

func (a *AddLocIdCount) AddCfg(locidcoInCfg [][]int32) error {
	if a.HeroVOs == nil {
		heroVOs := []*hero_db.HeroVO{}
		a.HeroVOs = heroVOs
	}
	a.Old_HeroVOs = make([]*hero_db.HeroVO, len(a.HeroVOs))
	for i, v := range a.HeroVOs {
		a.Old_HeroVOs[i] = v.Clone()
	}

	if a.PlayerInfoVOV == nil {
		var err error
		a.PlayerInfoVOV, err = playerInfo_db.GetPlayerInfoVOBy(a.PlayerId)
		if err != nil {
			return err
		}
	}
	a.Old_PlayerInfoVOV = a.PlayerInfoVOV.Clone()

	for _, int32Arr := range locidcoInCfg {
		loc := int32Arr[0]
		id := int32Arr[1]
		co := int32Arr[2]
		if err := a.addSingleHelp(loc, id, co); err != nil {
			return err
		}
	}
	return nil
}

func (a *AddLocIdCount) getOldHeroVO(heroId int32) *hero_db.HeroVO {
	for _, v := range a.Old_HeroVOs {
		if v.Id == heroId {
			return v
		}
	}
	return nil
}

func (a *AddLocIdCount) DiffSave() error {
	for _, v := range a.HeroVOs {
		seelog.Trace("DiffSave", v)
		old_v := a.getOldHeroVO(v.Id)
		if old_v == nil {
			if err := v.NewSave(); err != nil {
				return err
			}
		} else {
			if err := v.DiffSave(old_v); err != nil {
				return err
			}
		}
	}
	if err := a.PlayerInfoVOV.DiffSave(a.Old_PlayerInfoVOV); err != nil {
		return err
	}
	return nil
}*/
