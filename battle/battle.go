package battle

import (
	"buffalo/king/cfgmgr/cfgHeroMgr"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/battleResult_db"
	"buffalo/king/king"
	"github.com/kingbuffalo/seelog"
	//"sort"
	//"fmt"
	//"strings"
)

const (
	sKILL_MAX int = 20 //一个英雄最多有3个trigger技能 一共30个，又因第一个英雄登场trigger 与进入战场 trigger是同时进行的，故60个
)

var time_logForBetterRead = map[int32]string{
	enumErrCode.TRIGGER_TIME_ENTER_BATTLE:    "TRIGGER_TIME_ENTER_BATTLE",
	enumErrCode.TRIGGER_TIME_GOTO_PLATFORM:   "TRIGGER_TIME_GOTO_PLATFORM",
	enumErrCode.TRIGGER_TIME_BEFORE_ATK:      "TRIGGER_TIME_BEFORE_ATK",
	enumErrCode.TRIGGER_TIME_ATK:             "TRIGGER_TIME_ATK",
	enumErrCode.TRIGGER_TIME_AFTER_ATK:       "TRIGGER_TIME_AFTER_ATK",
	enumErrCode.TRIGGER_TIME_BEFOR_ALLACTION: "TRIGGER_TIME_BEFOR_ALLACTION",
	enumErrCode.TRIGGER_TIME_LEFT_PLATFORM:   "TRIGGER_TIME_LEFT_PLATFORM",
	enumErrCode.TRIGGER_TIME_SORT:            "TRIGGER_TIME_SORT",
}

var target_logForBetterRead = map[int32]string{
	enumErrCode.TRIGGER_TARGET_SELF:    "TRIGGER_TARGET_SELF",
	enumErrCode.TRIGGER_TARGET_TEAMER:  "TRIGGER_TARGET_TEAMER",
	enumErrCode.TRIGGER_TARGET_ENEMY:   "TRIGGER_TARGET_ENEMY",
	enumErrCode.TRIGGER_TARGET_SOMEONE: "TRIGGER_TARGET_SOMEONE",
}

var targetIdx_logForBetterRead = map[int32]string{
	enumErrCode.TAR_IDX_SELF:               "TAR_IDX_SELF",
	enumErrCode.TAR_IDX_TEAMER_PLATFORM:    "TAR_IDX_TEAMER_PLATFORM",
	enumErrCode.TAR_IDX_TEAMER_BEFORE_SELF: "TAR_IDX_TEAMER_BEFORE_SELF",
	enumErrCode.TAR_IDX_TEAMER_AFTER_SELF:  "TAR_IDX_TEAMER_AFTER_SELF",
	enumErrCode.TAR_IDX_TEAMER_ALL:         "TAR_IDX_TEAMER_ALL",
	enumErrCode.TAR_IDX_OPPO_THE_SAME_POS:  "TAR_IDX_OPPO_THE_SAME_POS",
	enumErrCode.TAR_IDX_OPPO_PLATFORM:      "TAR_IDX_OPPO_PLATFORM",
	enumErrCode.TAR_IDX_OPPO_BEFORE_SELF:   "TAR_IDX_OPPO_BEFORE_SELF",
	enumErrCode.TAR_IDX_OPPO_AFTER_SELF:    "TAR_IDX_OPPO_AFTER_SELF",
	enumErrCode.TAR_IDX_OPPO_ALL:           "TAR_IDX_OPPO_ALL",
}

type manSkill_t struct {
	round    int32
	subRound int32
	pos      int32
}

type sortedFighterSkill_t struct {
	f            *fighter_t
	cfgHeroSkill *cfgHeroMgr.CfgHeroSkill
}

type Battle struct {
	playerId                        int32
	aq                              *queue_t
	dq                              *queue_t
	round                           int32
	subRound                        int32
	sortedFighterSkillArr           [sKILL_MAX]sortedFighterSkill_t
	sortedFighterSkillArrContentIdx int
	sortedFighterSkillArrRunIdx     int
	quicker                         *fighter_t
	slower                          *fighter_t
	bEnd                            bool
	conStep                         int32
	conStepNext                     int32
	eventId                         int32
	rounds                          [](*battleResult_db.RoundVO)
	currentRounds                   *[](*battleResult_db.RoundVO)
	alreadyDoOneAction              bool
}

func (b *Battle) SetPlayerId(pid int32) {
	b.playerId = pid
}
func (b *Battle) round0() {
	seelog.Trace("round0")
	b.aq.round0_setPlatformAttacker()
	b.dq.round0_setPlatformAttacker()
	b.aq.round0_allFighterEnterBattle()
	b.dq.round0_allFighterEnterBattle()
	b.aq.round0_dispatchGotoPlatform()
	b.dq.round0_dispatchGotoPlatform()
	b.round0DoAllSkill()
}

func (b *Battle) getQuickerSlower() {
	seelog.Trace("getQuickerSlower")
	a1 := b.aq.attacker.getSpeed()
	a2 := b.dq.attacker.getSpeed()
	if a1 >= a2 {
		b.quicker = b.aq.attacker
		b.slower = b.dq.attacker
	} else {
		b.quicker = b.dq.attacker
		b.slower = b.aq.attacker
	}
}

func (b *Battle) round0DoAllSkill() {
	seelog.Trace("round0DoAllSkill")
	for b.sortedFighterSkillArrRunIdx < b.sortedFighterSkillArrContentIdx {
		b.theMaxSpeedDoSkill()
	}
}

func (b *Battle) theMaxSpeedDoSkill() bool {
	seelog.Trace("theMaxSpeedDoSkill")
	if b.sortedFighterSkillArrRunIdx < b.sortedFighterSkillArrContentIdx {
		var maxSpeed int32 = -1
		var maxSpeed_sortedFighterSkill *sortedFighterSkill_t
		swapIdx := b.sortedFighterSkillArrRunIdx
		for i := b.sortedFighterSkillArrRunIdx; i < b.sortedFighterSkillArrContentIdx; i++ {
			f := b.sortedFighterSkillArr[i].f
			f.updateHeroAtr()
			fspeed := f.heroAtr[enumErrCode.CALC_ATR_SPEED]
			if maxSpeed < fspeed {
				maxSpeed = fspeed
				maxSpeed_sortedFighterSkill = &b.sortedFighterSkillArr[i]
				swapIdx = i
			}
		}
		bReleaseSkill := false
		if maxSpeed_sortedFighterSkill != nil {
			f := maxSpeed_sortedFighterSkill.f
			cfg := maxSpeed_sortedFighterSkill.cfgHeroSkill
			f.releaseTriggerSkill(cfg)
			bReleaseSkill = true
		}
		idx := b.sortedFighterSkillArrRunIdx
		if idx != swapIdx {
			b.sortedFighterSkillArr[swapIdx].f = b.sortedFighterSkillArr[idx].f
			b.sortedFighterSkillArr[swapIdx].cfgHeroSkill = b.sortedFighterSkillArr[idx].cfgHeroSkill
		}
		b.sortedFighterSkillArrRunIdx++
		seelog.Trace("theMaxSpeedDoSkill return:", bReleaseSkill)
		return bReleaseSkill
	}
	seelog.Trace("theMaxSpeedDoSkill return:", false)
	return false
}

func (b *Battle) getOrCreateRoundResult() *battleResult_db.RoundVO {
	seelog.Trace("getOrCreateRoundResult")
	for _, r := range b.rounds {
		if r.Round == b.round {
			return r
		}
	}
	r := battleResult_db.NewRoundVO(b.round)
	b.rounds = append(b.rounds, r)
	return r
}

func (b *Battle) addAction(attackPos, skillId int32, events []*battleResult_db.BattleEventVO, targetBeActions []*battleResult_db.TargetBeActionVO) {
	seelog.Tracef("b addAction,apos=%d,skillId=%d", attackPos, skillId)
	if len(events) > 0 {
		for i, _ := range events {
			events[i].Id = b.eventId + 1
			b.eventId++
		}
	}
	r := b.getOrCreateRoundResult()
	r.AddAction(attackPos, skillId, events, targetBeActions)

	for _, v := range *b.currentRounds {
		if v.Round == r.Round {
			return
		}
	}
	*b.currentRounds = append(*b.currentRounds, r)
}

func (b *Battle) addFighterSkill(f *fighter_t, cfgHeroSkill *cfgHeroMgr.CfgHeroSkill) {
	seelog.Trace("addFighterSkill", f.pos, cfgHeroSkill)
	if b.sortedFighterSkillArrContentIdx < sKILL_MAX-1 {
		b.sortedFighterSkillArr[b.sortedFighterSkillArrContentIdx].f = f
		b.sortedFighterSkillArr[b.sortedFighterSkillArrContentIdx].cfgHeroSkill = cfgHeroSkill
		b.sortedFighterSkillArrContentIdx++
	}
}

func (b *Battle) StartFight() []*king.Round {
	seelog.Trace("-------------------Fight-------------------")
	var rs []*battleResult_db.RoundVO = []*battleResult_db.RoundVO{}
	b.currentRounds = &rs
	b.round0()
	b.round = 1
	b.subRound = 0
	b.conStep = enumErrCode.END_OF_THIS_STEP

	var r []*king.Round = make([]*king.Round, len(rs))
	for i, v := range rs {
		r[i] = v.GenProtoData()
	}
	return r
}

func (b *Battle) ConRunOneStep(skillId1, skillid2, skillid3 int32) []*king.Round {
	seelog.Trace("ConRunOneStep")
	var rs []*battleResult_db.RoundVO = []*battleResult_db.RoundVO{}
	b.currentRounds = &rs
	b.alreadyDoOneAction = false
	b.atLeastDoOneAction()
	var r []*king.Round = make([]*king.Round, len(rs))
	for i, v := range rs {
		r[i] = v.GenProtoData()
	}
	return r

}

func (b *Battle) atLeastDoOneAction() {
	seelog.Trace("atLeastDoOneAction")
	seelog.Trace(b.sortedFighterSkillArrRunIdx, b.sortedFighterSkillArrContentIdx)
	if b.alreadyDoOneAction {
		return
	}
	if b.sortedFighterSkillArrRunIdx < b.sortedFighterSkillArrContentIdx {
		bHasReleaseSkill := false
		for b.sortedFighterSkillArrRunIdx < b.sortedFighterSkillArrContentIdx {
			if b.theMaxSpeedDoSkill() {
				bHasReleaseSkill = true
				break
			}
		}
		if !bHasReleaseSkill {
			battleResult := b.conFight(true)
			if battleResult == enumErrCode.BATTLE_NOT_END {
				b.atLeastDoOneAction()
			}
		}
	} else {
		if battleResult := b.checkBattleEnd(); battleResult == enumErrCode.BATTLE_NOT_END {
			b.conFight(false)
			b.atLeastDoOneAction()
		}
	}
}

func (b *Battle) setEmptySortFighterSkill() {
	seelog.Trace("setEmptySortFighterSkill")
	b.sortedFighterSkillArrContentIdx = 0
	b.sortedFighterSkillArrRunIdx = 0
}

//func (b *Battle) sortedFighterSkillAction() bool {
//return false
//bHasReleaseSkill := false
//for i := 0; i < 10; i++ {
//b.sortedFighterSkillAction_Idx = b.sortedFighterSkillAction_Idx + 1
//if b.sortedFighterSkillAction_Idx < sKILL_MAX {
//idx := b.sortedFighterSkillAction_Idx
//sortedFighterSkillVO := b.sortedFighterSkillArr[idx]
//if sortedFighterSkillVO.f != nil {
//bHasReleaseSkill = sortedFighterSkillVO.f.releaseTriggerSkillInSort(sortedFighterSkillVO.tblHeroSkillVO)
//b.subRound = b.subRound + 1
//if bHasReleaseSkill {
//fmt.Printf("%d,%d\n", sortedFighterSkillVO.f.getBattleUnitPos(), sortedFighterSkillVO.tblHeroSkillVO.FldSkillId)
//}
////TODO  add manSkill here
//} else {
//break
//}
//if sortedFighterSkillVO.tblHeroSkillVO.FldContinue == 0 {
//break
//}
//} else {
//break
//}
//}
//if bHasReleaseSkill {
//fmt.Println("speed,wuli,zhili,tongshai,bleed,bleedmax,atk,def,zhiatk,zhidef,arrowatk,calcSpeed")
//for _, v := range b.aq.teamer {
//if v.bValid {
//v.traceMeAtr()
//}
//}
//for _, v := range b.dq.teamer {
//if v.bValid {
//v.traceMeAtr()
//}
//}
//}

//return bHasReleaseSkill
//}

func (b *Battle) beforeAction() {
	seelog.Trace("beforeAction")
	b.setEmptySortFighterSkill()
	b.aq.beforeAction()
	b.dq.beforeAction()
	//b.sortedFighterSkillAction()
}

func (b *Battle) checkBattleEnd() int32 {
	seelog.Trace("checkBattleEnd")
	if b.round > enumErrCode.BATTLE_STEP_MAX {
		return enumErrCode.BATTLE_END_RESULT_FAIL
	}
	if !b.aq.beSomeLive() {
		return enumErrCode.BATTLE_END_RESULT_FAIL
	}
	if !b.dq.beSomeLive() {
		return enumErrCode.BATTLE_END_RESULT_SUC
	}
	ret := b.aq.checkAttackerLeftPlatform()
	if ret == -1 {
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.ATTACKER_GOTO_PLATFORM)
		return enumErrCode.BATTLE_NOT_END
	} else {
		if ret != enumErrCode.BATTLE_NOT_END {
			return ret
		} else {
			ret = b.dq.checkAttackerLeftPlatform()
			if ret == -1 {
				b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.DEFENDER_GOTO_PLATFORM)
				return enumErrCode.BATTLE_NOT_END
			} else {
				return ret
			}
		}
	}
}

func (b *Battle) bHasSortFighterSkill() bool {
	seelog.Trace("bHasSortFighterSkill")
	return b.sortedFighterSkillArrContentIdx != 0
}

var step_logforbetterread = map[int32]string{
	enumErrCode.SKILL_ACTION:                    "SKILL_ACTION",
	enumErrCode.GEN_SKILL_ACTION_START:          "GEN_SKILL_ACTION_START",
	enumErrCode.ATTACKER_GOTO_PLATFORM:          "ATTACKER_GOTO_PLATFORM",
	enumErrCode.DEFENDER_GOTO_PLATFORM:          "DEFENDER_GOTO_PLATFORM",
	enumErrCode.END_OF_THIS_STEP:                "END_OF_THIS_STEP",
	enumErrCode.ATTACK_BY_SPEED_Q_ATTACK_BEFORE: "ATTACK_BY_SPEED_Q_ATTACK_BEFORE",
	enumErrCode.ATTACK_BY_SPEED_Q_ATTACK:        "ATTACK_BY_SPEED_Q_ATTACK",
	enumErrCode.ATTACK_BY_SPEED_Q_ATTACK_AFTER:  "ATTACK_BY_SPEED_Q_ATTACK_AFTER",
	enumErrCode.ATTACK_BY_SPEED_S_ATTACK_BEFORE: "ATTACK_BY_SPEED_S_ATTACK_BEFORE",
	enumErrCode.ATTACK_BY_SPEED_S_ATTACK:        "ATTACK_BY_SPEED_S_ATTACK",
	enumErrCode.ATTACK_BY_SPEED_S_ATTACK_AFTER:  "ATTACK_BY_SPEED_S_ATTACK_AFTER",
	enumErrCode.SORT_ACTION:                     "SORT_ACTION",
}

func (b *Battle) doStatueChangeHelp(conStep, conStepNext int32) {
	c := step_logforbetterread[conStep]
	var cn string = "0"
	if conStepNext != 0 {
		cn = step_logforbetterread[conStepNext]
	}
	seelog.Trace("doStatueChangeHelp,", c, " ", cn)
	if b.bHasSortFighterSkill() {
		b.conStep = conStep
		b.conStepNext = conStepNext
		//b.sortedFighterSkillAction()
	} else {
		b.conStep = conStepNext
		b.conFight(false)
	}
}

func (b *Battle) endOfThisStep() {
	seelog.Trace("endOfThisStep")
	b.aq.endOfThisStep()
	b.dq.endOfThisStep()
}

func (b *Battle) GetRoundResults() []*battleResult_db.RoundVO {
	return b.rounds
}

func (b *Battle) slowerAttackSlower() {
	seelog.Trace("slowerAttackSlower")
	b.setEmptySortFighterSkill()
	b.slower.attack(b.quicker)
}

func (b *Battle) slowerAttackSlower_after() {
	seelog.Trace("slowerAttackSlower_after")
	b.setEmptySortFighterSkill()
	b.slower.afterAttack(b.quicker)
}

func (b *Battle) slowerAttackSlower_before() {
	seelog.Trace("slowerAttackSlower_before")
	b.setEmptySortFighterSkill()
	b.slower.beforeAttack(b.quicker)
}

func (b *Battle) quickerAttackSlower() {
	seelog.Trace("quickerAttackSlower")
	b.setEmptySortFighterSkill()
	b.quicker.attack(b.slower)
}

func (b *Battle) quickerAttackSlower_after() {
	seelog.Trace("quickerAttackSlower_after")
	b.setEmptySortFighterSkill()
	b.quicker.afterAttack(b.slower)
}

func (b *Battle) quickerAttackSlower_before() {
	seelog.Trace("quickerAttackSlower_before")
	b.setEmptySortFighterSkill()
	b.quicker.beforeAttack(b.slower)
}

func (b *Battle) conFight(bCheckBattleEnd bool) int32 {
	seelog.Trace("conFight", bCheckBattleEnd)
	if bCheckBattleEnd {
		if ret := b.checkBattleEnd(); ret != enumErrCode.BATTLE_NOT_END {
			seelog.Trace("checkBattleEnd", ret)
			return ret
		}
	}
	switch b.conStep {
	case enumErrCode.SKILL_ACTION:
		b.conStep = b.conStepNext
		//b.atLeastDoOneAction()
		//panic("stop")
		//这里应该停掉循环
		//b.atLeastDoOneAction()
		//if !b.sortedFighterSkillAction() {
		//b.conStep = b.conStepNext
		//b.conFight(false)
		//}
	case enumErrCode.ATTACKER_GOTO_PLATFORM:
		b.aq.gotoPlatform()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.END_OF_THIS_STEP)
	case enumErrCode.DEFENDER_GOTO_PLATFORM:
		b.dq.gotoPlatform()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.END_OF_THIS_STEP)
	case enumErrCode.END_OF_THIS_STEP:
		b.endOfThisStep()
		b.conStep = enumErrCode.GEN_SKILL_ACTION_START
		b.conFight(true)
	case enumErrCode.ATTACK_BY_SPEED_Q_ATTACK_BEFORE:
		b.getQuickerSlower()
		b.quickerAttackSlower_before()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.ATTACK_BY_SPEED_Q_ATTACK)
	case enumErrCode.ATTACK_BY_SPEED_Q_ATTACK:
		b.quickerAttackSlower()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.ATTACK_BY_SPEED_Q_ATTACK_AFTER)
	case enumErrCode.ATTACK_BY_SPEED_Q_ATTACK_AFTER:
		b.quickerAttackSlower_after()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.ATTACK_BY_SPEED_S_ATTACK_BEFORE)
	case enumErrCode.ATTACK_BY_SPEED_S_ATTACK_BEFORE:
		b.slowerAttackSlower_before()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.ATTACK_BY_SPEED_S_ATTACK)
	case enumErrCode.ATTACK_BY_SPEED_S_ATTACK:
		b.slowerAttackSlower()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.ATTACK_BY_SPEED_S_ATTACK_AFTER)
	case enumErrCode.ATTACK_BY_SPEED_S_ATTACK_AFTER:
		b.slowerAttackSlower_after()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.SORT_ACTION)
	case enumErrCode.SORT_ACTION:
		b.sortAction()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.END_OF_THIS_STEP)
	case enumErrCode.GEN_SKILL_ACTION_START:
		b.genBeforeAttackAction()
		b.doStatueChangeHelp(enumErrCode.SKILL_ACTION, enumErrCode.ATTACK_BY_SPEED_Q_ATTACK_BEFORE)
	}
	return enumErrCode.BATTLE_NOT_END
}

func (b *Battle) sortAction() {
	seelog.Trace("sortAction")
	b.setEmptySortFighterSkill()
	b.aq.sortAction()
	b.dq.sortAction()
}

func (b *Battle) genBeforeAttackAction() {
	seelog.Trace("genBeforeAttackAction")
	b.round = b.round + 1
	b.subRound = 0
	b.beforeAction()
}

/*

func (b *Battle) bBattleEnd() bool {
	return !b.aq.beSomeLive() || !b.dq.beSomeLive()
}

func (b *Battle) getFighter(unitPos int) *fighter_t {
	if unitPos < 5 {
		return &b.aq.teamer[unitPos]
	}
	return &b.dq.teamer[unitPos-5]
}




func (b *Battle) manReleaseSkill(unitPos int) {
	f := b.getFighter(unitPos)
	f.releaseManSkill()
}

}*/
