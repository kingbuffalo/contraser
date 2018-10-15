package battle

import (
	"buffalo/king/common/enumErrCode"
	//"buffalo/king/dbOper/hero_db"
	"github.com/kingbuffalo/seelog"
)

type queue_t struct {
	b          *Battle
	attacker   *fighter_t
	teamer     [enumErrCode.QUEUE_MAX](*fighter_t)
	skillPoint int32
	oppo_q     *queue_t
}

func (q *queue_t) gotoPlatform() {
	seelog.Trace("gotoPlatform")
	q.b.setEmptySortFighterSkill()
	q.attacker.gotoPlatform()
}

func (q *queue_t) round0_setPlatformAttacker() {
	seelog.Trace("round0_setPlatformAttacker")
	q.attacker = q.teamer[0]
	seelog.Trace("round0_setPlatformAttacker", q.attacker)
	q.attacker.round0GotoPlatform()
}

func (q *queue_t) endOfThisStep() {
	seelog.Trace("endOfThisStep")
	for i, _ := range q.teamer {
		v := q.teamer[i]
		if v != nil {
			v.endOfThisStep()
		}
	}
	q.skillPoint += enumErrCode.SKILL_POINT_STEP_ADD
	if q.skillPoint > enumErrCode.SKILL_POINT_STEP_MAX {
		q.skillPoint = enumErrCode.SKILL_POINT_STEP_MAX
	}
}

func (q *queue_t) round0_allFighterEnterBattle() {
	seelog.Trace("round0_allFighterEnterBattle")
	for _, v := range q.teamer {
		if v != nil {
			v.enterBattle()
		}
	}
}

func (q *queue_t) beforeAction() {
	seelog.Trace("beforeAction")
	for _, v := range q.teamer {
		if v != nil {
			v.beforeAction()
		}
	}
}

func (q *queue_t) round0_dispatchGotoPlatform() {
	seelog.Trace("round0_dispatchGotoPlatform")
	q.attacker.round0GotoPlatformTrigger()
}

func (q *queue_t) checkAttackerLeftPlatform() int32 {
	seelog.Trace("checkAttackerLeftPlatform")
	var beTheNewHeroGotoPlatform bool = false
	var oldAttack *fighter_t
	if q.attacker != nil {
		if !q.attacker.beALive() {
			oldAttack = q.attacker
			q.attacker = q.searchTheFirstHeroIsLive()
			beTheNewHeroGotoPlatform = true
		}
	} else {
		q.attacker = q.searchTheFirstHeroIsLive()
		beTheNewHeroGotoPlatform = true
	}
	if q.attacker == nil {
		if q.b.aq == q {
			return enumErrCode.BATTLE_END_RESULT_FAIL
		} else {
			return enumErrCode.BATTLE_END_RESULT_SUC
		}
	}
	if !q.attacker.beALive() {
		if q.b.aq == q {
			return enumErrCode.BATTLE_END_RESULT_FAIL
		} else {
			return enumErrCode.BATTLE_END_RESULT_SUC
		}
	}

	if beTheNewHeroGotoPlatform {
		q.b.setEmptySortFighterSkill()
		if oldAttack != nil {
			oldAttack.leftThePlatForm()
			return -1
		}
	}
	return enumErrCode.BATTLE_NOT_END
}

func (q *queue_t) beSomeLive() bool {
	seelog.Trace("beSomeLive")
	for i, _ := range q.teamer {
		v := q.teamer[i]
		if v != nil {
			if v.beALive() {
				return true
			}
		}
	}
	return false
}

func (q *queue_t) sortAction() {
	seelog.Trace("sortAction")
	for i, _ := range q.teamer {
		f := q.teamer[i]
		if f != nil {
			f.sortAction()
		}
	}
}

func (q *queue_t) searchTheFirstHeroIsLive() *fighter_t {
	seelog.Trace("searchTheFirstHeroIsLive")
	var i int32
	for i = 0; i < enumErrCode.QUEUE_MAX; i++ {
		f := q.teamer[i]
		if f != nil {
			if f.beALive() {
				return f
			}
		}
	}
	return nil
}
