package battle

import (
	//"buffalo/king/cfgmgr/cfgArmyMgr"
	"buffalo/king/cfgmgr/cfgHeroMgr"
	"buffalo/king/common/enumErrCode"
	//"buffalo/king/dbOper/battleResult_db"
	"buffalo/king/dbOper/hero_db"
	"github.com/kingbuffalo/seelog"
	//"container/list"
)

func newFighter(q *queue_t, heroVO *hero_db.HeroVO, pos int32) *fighter_t {
	seelog.Trace("newFighter", heroVO, pos)
	cfg := cfgHeroMgr.GetCfgHero(heroVO.Id)
	f := &fighter_t{
		q:         q,
		pos:       pos,
		heroVO:    heroVO,
		cfgHero:   cfg,
		star:      heroVO.Star,
		level:     heroVO.Level,
		posMapbuf: make(map[int32]*buf_t),
	}
	s, w, b, z := cfg.GetInitAtr(f.star, f.level)
	f.heroAtrOrig[enumErrCode.ATR_SPEED] = s
	f.heroAtrOrig[enumErrCode.ATR_WULI] = w
	f.heroAtrOrig[enumErrCode.ATR_ZHILI] = z
	f.heroAtrOrig[enumErrCode.ATR_TONGSHAI] = b
	f.heroAtr[enumErrCode.CALC_ATR_BLEED] = 100
	//TODO add weapon and so on
	return f
}

func newFighterPVENPC(q *queue_t, heroId, troops, level, pos int32) *fighter_t {
	seelog.Trace("newFighter", heroId, troops, level, pos)
	cfg := cfgHeroMgr.GetCfgHero(heroId)
	if cfg == nil {
		seelog.Trace("cfg is nil,heroId=%d", heroId)
		return nil
	}

	f := &fighter_t{
		q:         q,
		pos:       pos,
		heroVO:    nil,
		cfgHero:   cfg,
		star:      cfg.StarLv,
		level:     level,
		posMapbuf: make(map[int32]*buf_t),
	}
	s, w, b, z := cfg.GetInitAtr(f.star, f.level)
	f.heroAtrOrig[enumErrCode.ATR_SPEED] = s
	f.heroAtrOrig[enumErrCode.ATR_WULI] = w
	f.heroAtrOrig[enumErrCode.ATR_ZHILI] = z
	f.heroAtrOrig[enumErrCode.ATR_TONGSHAI] = b
	f.heroAtr[enumErrCode.CALC_ATR_BLEED] = 100
	return f
}
