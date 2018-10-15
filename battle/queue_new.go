package battle

import (
	"buffalo/king/cfgmgr/cfgPVELevelMgr"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/hero_db"
	"github.com/kingbuffalo/seelog"
)

func newQueue(b *Battle, heroVOs []*hero_db.HeroVO, bAttacker bool) *queue_t {
	seelog.Trace("newQueue", heroVOs, bAttacker)
	q := &queue_t{
		b:          b,
		skillPoint: enumErrCode.SKILL_POINT_INIT,
	}
	for i, v := range heroVOs {
		pos := i + 1
		if !bAttacker {
			pos += 10
		}
		f := newFighter(q, v, int32(pos))
		q.teamer[i] = f
	}
	return q
}

func newQueuePVENpc(b *Battle, cfgPVELevel *cfgPVELevelMgr.CfgPVELevel) *queue_t {
	seelog.Trace("newQueuePVENpc", cfgPVELevel)
	q := &queue_t{
		b:          b,
		skillPoint: cfgPVELevel.InitSkillPoint,
	}
	for i, heroId := range cfgPVELevel.DefenceHeroIds {
		pos := int32(i + 11)
		level := cfgPVELevel.DefenceHeroLvs[i]
		troops := cfgPVELevel.DefenceHeroTroops[i]
		f := newFighterPVENPC(q, heroId, troops, level, pos)
		q.teamer[i] = f
		seelog.Trace("newQueuePVENpc", i, f)
	}
	return q

}
