package battle

import (
	"buffalo/king/cfgmgr/cfgArmyMgr"
	"buffalo/king/cfgmgr/cfgHeroMgr"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/dbOper/battleResult_db"
	"buffalo/king/dbOper/hero_db"
	"github.com/kingbuffalo/seelog"
	//"container/list"
)

type fighter_t struct {
	heroVO  *hero_db.HeroVO
	q       *queue_t
	pos     int32
	star    int32
	level   int32
	heroAtr [enumErrCode.ALL_ATR_MAX]int32
	//depends on cfgHero,weapon,and so on;calc_id_atr不受武智体等影响的初始值
	heroAtrOrig   [enumErrCode.ALL_ATR_MAX]int32
	bHeroAtrDirty bool
	bInPlatform   bool

	statue          uint64
	statueLast1Step uint64
	statueLast2Step uint64

	statueEffect1Times uint64
	statueEffect2Times uint64

	cfgHero              *cfgHeroMgr.CfgHero
	cfgArmySkillValueArr []*cfgArmyMgr.CfgArmySkillValue

	posMapbuf map[int32]*buf_t
}

type buf_t struct {
	cfg      *cfgHeroMgr.CfgHeroSkillBuff
	endRound int32
}

func (f *fighter_t) round0GotoPlatform() {
	seelog.Trace(f.pos, " round0GotoPlatform")
	f.bInPlatform = true
	cfg := cfgHeroMgr.Gotoplatform_CfgHeroSkill()
	f.releaseTriggerSkill(cfg)
}

func (f *fighter_t) round0GotoPlatformTrigger() {
	seelog.Trace(f.pos, "round0GotoPlatformTrigger")
	f.platformHeroOneActionBeDead_others(enumErrCode.TRIGGER_TIME_GOTO_PLATFORM)
}

func (f *fighter_t) gotoPlatform() {
	seelog.Trace(f.pos, "gotoPlatform")
	f.bInPlatform = true
	f.thePlatformHeroOneActionBeDead(enumErrCode.TRIGGER_TIME_GOTO_PLATFORM)
	cfg := cfgHeroMgr.Gotoplatform_CfgHeroSkill()
	f.releaseTriggerSkill(cfg)
}

func (f *fighter_t) enterBattle() {
	seelog.Trace(f.pos, "enterBattle")
	f.addTriggerSkill(enumErrCode.TRIGGER_TIME_ENTER_BATTLE, enumErrCode.TRIGGER_TARGET_SELF)
}

func (f *fighter_t) beALive() bool {
	seelog.Trace(f.pos, "beALive")
	return f.heroAtr[enumErrCode.CALC_ATR_BLEED] > 0
}

func (f *fighter_t) addTriggerSkill(triggerTime, triggertarget int32) {
	tt := time_logForBetterRead[triggerTime]
	t := target_logForBetterRead[triggertarget]
	seelog.Trace(f.pos, "addTriggerSkill,", tt, " ", t)
	if !f.beALive() {
		return
	}
	cfgskillArr := f.cfgHero.CfgHeroSkillArr
	for _, v := range cfgskillArr {
		f.addTriggerSkill_help(triggerTime, triggertarget, v)
	}
}

func (f *fighter_t) addTriggerSkill_help(triggerTime, triggertarget int32, cfgHeroSkill *cfgHeroMgr.CfgHeroSkill) {
	seelog.Trace(f.pos, "addTriggerSkill_help", triggerTime, triggertarget, cfgHeroSkill)
	if cfgHeroSkill != nil {
		if cfgHeroSkill.TriggerTime == triggerTime {
			if (cfgHeroSkill.SkillPos == enumErrCode.SKILL_POS_BOTH) || (cfgHeroSkill.SkillPos == enumErrCode.SKILL_POS_PLATFORM && f.bInPlatform) || (cfgHeroSkill.SkillPos == enumErrCode.SKILL_POS_OFF_PLATFORM && !f.bInPlatform) {
				bRightTrigger := cfgHeroSkill.TriggerTarget == triggertarget
				if !bRightTrigger {
					if cfgHeroSkill.TriggerTarget == enumErrCode.TRIGGER_TARGET_SOMEONE {
						bRightTrigger = triggertarget == enumErrCode.TRIGGER_TARGET_TEAMER || triggertarget == enumErrCode.TRIGGER_TARGET_ENEMY
					}
				}
				if bRightTrigger {
					f.q.b.addFighterSkill(f, cfgHeroSkill)
				}
			}
		}

	}
}

func (f *fighter_t) releaseTriggerSkill_helper(cfgHeroSkill *cfgHeroMgr.CfgHeroSkill) ([]*battleResult_db.BattleEventVO, []*battleResult_db.TargetBeActionVO, bool) {
	seelog.Trace(f.pos, "releaseTriggerSkill_helper", cfgHeroSkill)
	if cfgHeroSkill.Id == enumErrCode.SKILL_ID_GOTO_PLATFORM || cfgHeroSkill.Id == enumErrCode.SKILL_ID_LEFT_PLATFORM {
		return nil, nil, true
	} else {
		events := []*battleResult_db.BattleEventVO{}
		targetBeActions := []*battleResult_db.TargetBeActionVO{}
		bReleaseSkill := false
		cfgHeroSkillEffArr := cfgHeroSkill.CfgHeroSkillEffects
		for _, cfgHeroSkillEffect := range cfgHeroSkillEffArr {
			targetFighterArr := f.getTargetArr(cfgHeroSkillEffect.TargetIdx)
			if targetFighterArr != nil {
				for _, tf := range targetFighterArr {
					if tf != nil {
						bR := tf.beReleaseSkill(f, cfgHeroSkillEffect, &events, &targetBeActions)
						bReleaseSkill = bReleaseSkill || bR
					}
				}
			}
		}
		return events, targetBeActions, bReleaseSkill
	}
}

func (f *fighter_t) getSpeed() int32 {
	seelog.Trace(f.pos, "getSpeed")
	f.updateHeroAtr()
	return f.heroAtr[enumErrCode.CALC_ATR_SPEED]
}

func (f *fighter_t) releaseTriggerSkill(cfgHeroSkill *cfgHeroMgr.CfgHeroSkill) bool {
	seelog.Trace(f.pos, "releaseTriggerSkill", cfgHeroSkill)
	events, targetBeActions, bRelease := f.releaseTriggerSkill_helper(cfgHeroSkill)
	if bRelease {
		f.q.b.addAction(f.pos, cfgHeroSkill.Id, events, targetBeActions)
	}
	return bRelease
}

//func (f *fighter_t) gotoThePlatform() {
//f.bInPlatform = true
//f.thePlatformHeroOneActionBeDead(enumErrCode.TRIGGER_TIME_GOTO_PLATFORM)
//gotoSkillVO, _ := tblHeroSkillMgr.Gotoplatform_tblHeroSkillVO()
//f.q.b.addFighterSkill(f, gotoSkillVO)
//}

func (f *fighter_t) getTargetArr(targetIdx int32) [](*fighter_t) {
	t := targetIdx_logForBetterRead[targetIdx]
	seelog.Trace(f.pos, ",releaseTriggerSkill=", t)
	var targetFigherArr [](*fighter_t) = nil
	switch targetIdx {
	case enumErrCode.TAR_IDX_SELF:
		targetFigherArr = []*fighter_t{f}
	case enumErrCode.TAR_IDX_TEAMER_PLATFORM:
		targetFigherArr = []*fighter_t{f.q.attacker}
	case enumErrCode.TAR_IDX_TEAMER_BEFORE_SELF:
		idx := f.pos - 1
		if idx != -1 {
			targetFigherArr = []*fighter_t{f.q.teamer[idx]}
		}
	case enumErrCode.TAR_IDX_TEAMER_AFTER_SELF:
		idx := f.pos + 1
		if idx < enumErrCode.QUEUE_MAX {
			targetFigherArr = []*fighter_t{f.q.teamer[idx]}
		}
	case enumErrCode.TAR_IDX_TEAMER_ALL:
		targetFigherArr = f.q.teamer[:]
	case enumErrCode.TAR_IDX_OPPO_THE_SAME_POS:
		targetFigherArr = [](*fighter_t){
			f.q.oppo_q.teamer[f.pos],
		}
	case enumErrCode.TAR_IDX_OPPO_PLATFORM:
		targetFigherArr = []*fighter_t{f.q.oppo_q.attacker}
	case enumErrCode.TAR_IDX_OPPO_BEFORE_SELF:
		targetFigherArr = nil
	case enumErrCode.TAR_IDX_OPPO_AFTER_SELF:
		targetFigherArr = nil
	case enumErrCode.TAR_IDX_OPPO_ALL:
		targetFigherArr = f.q.oppo_q.teamer[:]
	}
	return targetFigherArr
}

func (f *fighter_t) beReleaseSkill(rf *fighter_t, cfgHeroSkillEffect *cfgHeroMgr.CfgHeroSkillEffect, events *[]*battleResult_db.BattleEventVO, targetBeActions *[]*battleResult_db.TargetBeActionVO) bool {
	seelog.Trace(f.pos, "beReleaseSkill", rf.pos, cfgHeroSkillEffect)
	skillType, skillValue := cfgHeroSkillEffect.SkillType, cfgHeroSkillEffect.SkillValue
	lastNum, calcIdTarget, calcIdReleaser := cfgHeroSkillEffect.LastNum, cfgHeroSkillEffect.CalcIdTarget, cfgHeroSkillEffect.CalcIdSource

	if f.bHasStatueAndTE(rf, enumErrCode.STATUE_UNABLE_BE_SKILL, enumErrCode.STATUE_UNABLE_BE_TEAMER_SKILL, enumErrCode.STATUE_UNABLE_BE_ENEMY_SKILL) {
		return false
	}
	tba := &battleResult_db.TargetBeActionVO{
		Pos:  f.pos,
		Type: skillType,
	}
	var values []int32
	ret := true
	switch skillType {
	case enumErrCode.SKILL_TYPE_ADD_BLEED:
		f.setBleedHelp(f.getBleed() + skillValue)
		values = []int32{skillValue}
	case enumErrCode.SKILL_TYPE_ATTACK:
		atk := rf.getAtr(f, enumErrCode.CALC_ATR_WU_ATK)
		def := rf.getAtr(f, enumErrCode.CALC_ATR_WU_DEF)
		hurt := atkHurtDef(atk, def)
		seelog.Trace("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx->atk=", atk, "def=", def, "hurt=", hurt)
		f.setBleedHelp(f.getBleed() - hurt)
		values = []int32{hurt}
	case enumErrCode.SKILL_TYPE_HURT:
		f.setBleedHelp(f.getBleed() + calcSkillValue(calcIdTarget, rf, f))
		rf.setBleedHelp(rf.getBleed() + calcSkillValue(calcIdReleaser, rf, f))
		//TODO add 2
		values = []int32{10}
	case enumErrCode.SKILL_TYPE_BUFFER:
		if f.bHasStatueAndTE(rf, enumErrCode.STATUE_UNABLE_BE_BUFFER, enumErrCode.STATUE_UNABLE_BE_TEAMER_BUFFER, enumErrCode.STATUE_UNABLE_BE_ENEMY_BUFFER) {
			ret = false
		} else {
			f.addBuffer(rf, cfgHeroSkillEffect.CfgHeroSkillBuffValue)
			values = []int32{cfgHeroSkillEffect.CfgHeroSkillBuffValue.Id}
		}
	case enumErrCode.SKILL_TYPE_EFFTIMES_STATUE:
		if f.bHasStatueAndTE(rf, enumErrCode.STATUE_UNABLE_BE_STATUE, enumErrCode.STATUE_UNABLE_BE_TEAMER_STATUE, enumErrCode.STATUE_UNABLE_BE_ENEMY_STATUE) {
			ret = false
		} else {
			f.addEffTimesStatue(rf, skillValue, lastNum)
			values = []int32{skillValue}
		}
	case enumErrCode.SKILL_TYPE_STATUE:
		if f.bHasStatueAndTE(rf, enumErrCode.STATUE_UNABLE_BE_STATUE, enumErrCode.STATUE_UNABLE_BE_TEAMER_STATUE, enumErrCode.STATUE_UNABLE_BE_ENEMY_STATUE) {
			ret = false
		} else {
			f.addStatue(rf, skillValue, lastNum)
			values = []int32{skillValue}
		}
	case enumErrCode.SKILL_TYPE_CLEAR_GOOD_STATUE:
		f.clearStatue(true)
	case enumErrCode.SKILL_TYPE_CLEAR_BAD_STATUE:
		f.clearStatue(false)
	case enumErrCode.SKILL_TYPE_CLEAR_GOOD_BUFFER:
		f.clearBuffer(true)
	case enumErrCode.SKILL_TYPE_CLEAR_BAD_BUFFER:
		f.clearBuffer(false)
	}
	if ret {
		tba.Values = values
		*targetBeActions = append(*targetBeActions, tba)
	}
	return ret
}

func (f *fighter_t) bInOneTeam(anotherF *fighter_t) bool {
	seelog.Trace(f.pos, "bInOneTeam", anotherF.pos)
	return f.q == anotherF.q
}

func (f *fighter_t) bHasStatueAndTE(rf *fighter_t, statue, teamerStatue, enemyStatue uint64) bool {
	seelog.Trace(f.pos, "bHasStatueAndTE", rf.pos, statue, teamerStatue, enemyStatue)
	if f.bHasStatue(statue) {
		return true
	}
	if f.bInOneTeam(rf) {
		if f.bHasStatue(teamerStatue) {
			return true
		}
	} else {
		if f.bHasStatue(enemyStatue) {
			return true
		}
	}
	return false
}

func (f *fighter_t) setBleedHelp(bl int32) {
	seelog.Trace(f.pos, "setBleedHelp", bl)
	f.heroAtr[enumErrCode.CALC_ATR_BLEED] = bl
	if f.heroAtr[enumErrCode.CALC_ATR_BLEED] > f.heroAtr[enumErrCode.CALC_ATR_BLEED_MAX] {
		f.heroAtr[enumErrCode.CALC_ATR_BLEED] = f.heroAtr[enumErrCode.CALC_ATR_BLEED_MAX]
	}
	if f.heroAtr[enumErrCode.CALC_ATR_BLEED] <= 0 {
		f.heroAtr[enumErrCode.CALC_ATR_BLEED] = 0
		if f.bHasStatue(enumErrCode.STATUE_UNABLE_DEAD) {
			f.heroAtr[enumErrCode.CALC_ATR_BLEED] = f.heroAtr[enumErrCode.CALC_ATR_WU_ATK]
		}
	}
}

func (f *fighter_t) getBleed() int32 {
	seelog.Trace(f.pos, "getBleed")
	return f.heroAtr[enumErrCode.CALC_ATR_BLEED]
}

func (f *fighter_t) updateHeroAtr() {
	seelog.Trace(f.pos, "updateHeroAtr")
	if f.bHeroAtrDirty == true {
		//f.batBufList.clearTimeEndBuf(f.q.b.step)
		f.bHeroAtrDirty = false
		var i int32
		for i = 0; i < enumErrCode.ALL_ATR_MAX; i++ {
			f.heroAtr[i] = f.heroAtrOrig[i]
		}
		f.heroAtr[enumErrCode.CALC_ATR_WU_ATK] += f.heroAtr[enumErrCode.ATR_WULI] * 10
		f.heroAtr[enumErrCode.CALC_ATR_WU_DEF] += f.heroAtr[enumErrCode.ATR_WULI] * 10
		f.heroAtr[enumErrCode.CALC_ATR_ZHI_ATK] += f.heroAtr[enumErrCode.ATR_ZHILI] * 10
		f.heroAtr[enumErrCode.CALC_ATR_ZHI_DEF] += f.heroAtr[enumErrCode.ATR_ZHILI] * 10
		f.heroAtr[enumErrCode.CALC_ATR_BLEED] += f.heroAtr[enumErrCode.ATR_TONGSHAI] * 30
		//TODO Buff

		if f.heroAtr[enumErrCode.CALC_ATR_BLEED] > f.heroAtr[enumErrCode.CALC_ATR_BLEED_MAX] {
			f.heroAtr[enumErrCode.CALC_ATR_BLEED] = f.heroAtr[enumErrCode.CALC_ATR_BLEED_MAX]
		}
	}
}
func (f *fighter_t) releaseManSkill() bool {
	seelog.Trace(f.pos, "releaseManSkill")
	if f == nil {
		return false
	}
	if !f.beALive() {
		return false
	}
	if f.bHasStatue(enumErrCode.STATUE_UNABLE_MAN_SKILL) {
		return false
	}
	if f.q.skillPoint < f.cfgHero.ManCfgHeroSkill.SkillPoint {
		return false
	}

	if f.releaseTriggerSkill(f.cfgHero.ManCfgHeroSkill) {
		f.q.skillPoint -= f.cfgHero.ManCfgHeroSkill.SkillPoint
		return true
	}
	return false
}

func (f *fighter_t) getAtr(targetF *fighter_t, atr int32) int32 {
	seelog.Trace(f.pos, "getAtr", targetF.pos, atr)
	f.updateHeroAtr()
	value := f.heroAtr[atr]

	var targetArmyType int32 = 0
	if targetF != nil {
		targetArmyType = targetF.cfgHero.ArmyType
	}
	for _, v := range f.cfgArmySkillValueArr {
		value += v.GetValue(targetArmyType, atr)
	}
	return value
}

func (f *fighter_t) addEffTimesStatue(rf *fighter_t, skillValue, lastNum int32) {
	seelog.Trace(f.pos, "addEffTimesStatue", rf.pos, skillValue, lastNum)
	var bitMask uint64 = 1 << uint64(skillValue)
	switch lastNum {
	case 2:
		f.statueEffect2Times |= bitMask
	case 1:
		f.statueEffect1Times |= bitMask
	default:
	}
}

func (f *fighter_t) addStatue(rf *fighter_t, skillValue, lastNum int32) {
	seelog.Trace(f.pos, "addStatue", rf.pos, skillValue, lastNum)
	var bitMask uint64 = 1 << uint64(skillValue)
	switch lastNum {
	case 2:
		f.statueLast2Step |= bitMask
	case 1:
		f.statueLast1Step |= bitMask
	default:
		f.statue |= bitMask
	}
}

func (f *fighter_t) clearStatue(bGoodStatue bool) {
	seelog.Trace(f.pos, "clearStatue", bGoodStatue)
	var bitMask uint64
	if bGoodStatue {
		bitMask = enumErrCode.GOOD_STATUE_BIT_MASK
	} else {
		bitMask = enumErrCode.BAD_STATUE_BIT_MASK
	}
	bitMask ^= 0
	f.statue &= bitMask
	f.statueLast1Step &= bitMask
	f.statueLast2Step &= bitMask
	f.statueEffect1Times &= bitMask
	f.statueEffect2Times &= bitMask
}

func (f *fighter_t) clearBuffer(bGoodBuffer bool) {
	seelog.Trace(f.pos, "clearBuffer", bGoodBuffer)
	rmPosArr := make([]int32, 0)
	bt := enumErrCode.BUF_TY_GOOD_BUFFER
	if !bGoodBuffer {
		bt = enumErrCode.BUF_TY_BAD_BUFFER
	}
	for pos, v := range f.posMapbuf {
		if v.cfg.BuffType == bt {
			rmPosArr = append(rmPosArr, pos)
		}
	}
	for _, v := range rmPosArr {
		delete(f.posMapbuf, v)
	}
	f.bHeroAtrDirty = true
}

func (f *fighter_t) addBuffer(rf *fighter_t, cfgHeroSkillBuff *cfgHeroMgr.CfgHeroSkillBuff) {
	seelog.Trace(f.pos, "addBuffer", rf.pos, cfgHeroSkillBuff)
	if cfgHeroSkillBuff == nil {
		return
	}
	es := cfgHeroSkillBuff.LastNum + f.q.b.round
	if cfgHeroSkillBuff.LastNum == 0 {
		es = enumErrCode.STEP_MAX
	}
	f.posMapbuf[rf.pos] = &buf_t{
		cfg:      cfgHeroSkillBuff,
		endRound: es,
	}
	f.bHeroAtrDirty = true
}

func (f *fighter_t) releaseTriggerSkillInSort(cfgHeroSkill *cfgHeroMgr.CfgHeroSkill) bool {
	seelog.Trace(f.pos, "releaseTriggerSkillInSort", cfgHeroSkill)
	if !f.beALive() {
		return false
	}
	if f.bHasStatue(enumErrCode.STATUE_UNABLE_SKILL) {
		return false
	}
	return f.releaseTriggerSkill(cfgHeroSkill)
}

func (f *fighter_t) bHasStatue(statue uint64) bool {
	seelog.Trace(f.pos, "bHasStatue", statue)
	var bitMask uint64 = 1 << statue
	if (f.statue & bitMask) != 0 {
		return true
	}
	if (f.statueLast1Step & bitMask) != 0 {
		return true
	}
	if (f.statueLast2Step & bitMask) != 0 {
		return true
	}
	setBitMask := 0 ^ bitMask
	if f.statueEffect1Times&bitMask != 0 {
		f.statueEffect1Times &= setBitMask
		return true
	}
	if f.statueEffect2Times&bitMask != 0 {
		f.statueEffect2Times &= setBitMask
		f.statueEffect1Times |= bitMask
		return true
	}
	return false
}

func (f *fighter_t) endOfThisStep() {
	seelog.Trace(f.pos, "endOfThisStep")
	f.statueLast1Step = f.statueLast2Step
	f.statueLast2Step = 0
}

func (f *fighter_t) platformHeroOneActionBeDead_others(triggerTime int32) bool {
	seelog.Trace(f.pos, "platformHeroOneActionBeDead_others", triggerTime)
	if !f.bInPlatform {
		return false
	}
	oppo_q := f.q.oppo_q
	defencer := oppo_q.attacker
	if !f.beALive() || (defencer != nil && !defencer.beALive()) {
		return true
	}
	for i, _ := range f.q.teamer {
		teamf := f.q.teamer[i]
		if teamf != nil {
			if teamf.pos != f.pos {
				teamf.addTriggerSkill(triggerTime, enumErrCode.TRIGGER_TARGET_TEAMER)
			}
		}
	}
	for i, _ := range oppo_q.teamer {
		teamf := oppo_q.teamer[i]
		if teamf != nil {
			teamf.addTriggerSkill(triggerTime, enumErrCode.TRIGGER_TARGET_ENEMY)
		}
	}
	return false
}

func (f *fighter_t) thePlatformHeroOneActionBeDead(triggerTime int32) bool {
	seelog.Trace(f.pos, "thePlatformHeroOneActionBeDead", triggerTime)
	if !f.bInPlatform {
		return false
	}
	f.addTriggerSkill(triggerTime, enumErrCode.TRIGGER_TARGET_SELF)
	return f.platformHeroOneActionBeDead_others(triggerTime)
}

func (f *fighter_t) beforeAction() {
	seelog.Trace(f.pos, "beforeAction")
	f.addTriggerSkill(enumErrCode.TRIGGER_TIME_BEFOR_ALLACTION, enumErrCode.TRIGGER_TARGET_SELF)
}

func (f *fighter_t) attack(defencer *fighter_t) {
	seelog.Trace(f.pos, "attack")
	defencer.beAttackBy(f)
}

func (f *fighter_t) beforeAttack(defencer *fighter_t) {
	seelog.Trace(f.pos, "beforeAttack")
	f.thePlatformHeroOneActionBeDead(enumErrCode.TRIGGER_TIME_BEFORE_ATK)
}

func (f *fighter_t) afterAttack(defencer *fighter_t) {
	seelog.Trace(f.pos, "afterAttack")
	f.thePlatformHeroOneActionBeDead(enumErrCode.TRIGGER_TIME_AFTER_ATK)
}

func (f *fighter_t) beAttackBy(af *fighter_t) {
	seelog.Trace(f.pos, "beAttackBy", af.pos)
	if f.bHasStatue(enumErrCode.STATUE_UNABLE_BE_ATK) {
		return
	}
	if af.bHasStatue(enumErrCode.STATUE_UNABLE_ATK) {
		return
	}
	//fmt.Printf("beAttackBy,%d,%d\n", f.getBattleUnitPos(), af.getBattleUnitPos())
	atkTblHeroSkillVO := cfgHeroMgr.NormalATK_CfgHeroSkill()
	f.q.b.addFighterSkill(af, atkTblHeroSkillVO)
}

/*
func (f *fighter_t) getBattleUnitPos() int32 {
	if &f.q.b.q1 == f.q {
		return f.pos
	}
	return f.pos + 5
}
func (f *fighter_t) beLive() bool {
	return f.getBleed() > 0
}
func (f *fighter_t) getHeroAtr() []int32 {
	f.updateHeroAtr()
	return f.heroAtr[:enumErrCode.ALL_ATR_MAX]
}


func (f *fighter_t) beforeAction() {
	//只对自己触发，因为每个人都会同时触发
	f.addTriggerSkill(enumErrCode.TRIGGER_TIME_BEFOR_ALLACTION, enumErrCode.TRIGGER_TARGET_SELF)
}


func (f *fighter_t) thePlatformHeroOneActionBeDead_forRead(triggerTime int32, bNeedDoTriggerSkill bool) bool {
	if !f.bInPlatform {
		return false
	}
	if bNeedDoTriggerSkill {
		f.q.b.setEmptySortFighterSkill()
	}
	f.addTriggerSkill(triggerTime, enumErrCode.TRIGGER_TARGET_SELF)
	oppo_q := f.q.oppo_q
	defencer := oppo_q.attacker
	if !f.beLive() || (defencer != nil && !defencer.beLive()) {
		return true
	}
	for i, _ := range f.q.teamer {
		teamf := &f.q.teamer[i]
		if teamf.bValid {
			if teamf.pos != f.pos {
				teamf.addTriggerSkill(triggerTime, enumErrCode.TRIGGER_TARGET_TEAMER)
			}
		}
	}
	for i, _ := range oppo_q.teamer {
		teamf := &oppo_q.teamer[i]
		if teamf.bValid {
			teamf.addTriggerSkill(triggerTime, enumErrCode.TRIGGER_TARGET_ENEMY)
		}
	}
	if bNeedDoTriggerSkill {
		f.q.b.sortedFighterSkillAction()
		if !f.beLive() || (defencer != nil && !defencer.beLive()) {
			return true
		}
	}
	return false
}



func (f fighter_t) traceMeAtr() {
	ui := f.getBattleUnitPos()
	var atrStrArr [enumErrCode.ALL_ATR_MAX]string
	curAtr := f.getHeroAtr()
	for i, v := range curAtr {
		atrStrArr[i] = strconv.Itoa(v)
	}
	str := strings.Join(atrStrArr[:enumErrCode.ALL_ATR_MAX], ",")
	str = strconv.Itoa(ui) + "," + str
	fmt.Println(str)
}

func (f *fighter_t) enterBattle() {
	f.addTriggerSkill(enumErrCode.TRIGGER_TIME_ENTER_BATTLE, enumErrCode.TRIGGER_TARGET_SELF)
}


*/

func (f *fighter_t) leftThePlatForm() {
	seelog.Trace(f.pos, "leftThePlatForm")
	f.thePlatformHeroOneActionBeDead(enumErrCode.TRIGGER_TIME_LEFT_PLATFORM)
	f.bInPlatform = false
	leftSkillVO := cfgHeroMgr.Leftplatform_CfgHeroSkill()
	f.q.b.addFighterSkill(f, leftSkillVO)
}
func (f *fighter_t) sortAction() {
	seelog.Trace(f.pos, "sortAction")
	//只对自己触发，因为每个人都会同时触发
	f.addTriggerSkill(enumErrCode.TRIGGER_TIME_SORT, enumErrCode.TRIGGER_TARGET_SELF)
}
