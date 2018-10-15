package battle

import (
	"github.com/kingbuffalo/seelog"
)

func atkHurtDef(atk, def int32) int32 {
	seelog.Trace("atkHurtDef", atk, def)
	var atkf float64 = float64(atk)
	var deff float64 = float64(def)
	sum := (atkf + deff)
	if sum == 0 {
		return 1
	}
	var hurt float64 = (atkf * atkf) / sum
	if hurt < 1 {
		hurt = 1
	}
	return int32(hurt)
}

func calcSkillValue(calcId int32, rf, tf *fighter_t) int32 {
	return 0
	//targetAtr := rf.getHeroAtr()
	//releaseAtr := tf.getHeroAtr()
	//value := 0
	//switch calcId {
	//case tblMgr.CID_ARROW_HURT:
	//value = -atkHurtDef(releaseAtr[tblMgr.CALC_ATR_ARROW_ATK], targetAtr[tblMgr.CALC_ATR_WU_DEF])
	//case tblMgr.CID_PAOXIAO_HURT:
	//value = -atkHurtDef(releaseAtr[tblMgr.ATR_WULI], targetAtr[tblMgr.ATR_WULI]) * 2
	//case tblMgr.CID_YUANJUN_ADD:
	//value = releaseAtr[tblMgr.CALC_ATR_BLEED_MAX] / 8
	//case tblMgr.CID_YUANJUN_SUB:
	//value = -releaseAtr[tblMgr.CALC_ATR_BLEED_MAX] / 8
	//case tblMgr.CID_HORSE_ARROW_HURT:
	//value = -3 * atkHurtDef(releaseAtr[tblMgr.CALC_ATR_ARROW_ATK], targetAtr[tblMgr.CALC_ATR_WU_DEF])
	//case tblMgr.CID_CEFAN_HURT:
	//value = -atkHurtDef(targetAtr[tblMgr.CALC_ATR_WU_ATK], targetAtr[tblMgr.CALC_ATR_WU_DEF])
	//case tblMgr.CID_LIANNU_HURT:
	//value = -2 * atkHurtDef(releaseAtr[tblMgr.CALC_ATR_ARROW_ATK], targetAtr[tblMgr.CALC_ATR_WU_DEF])
	//case tblMgr.CID_ANGYANG_ADD:
	//value = releaseAtr[tblMgr.CALC_ATR_BLEED_MAX] / 4
	//}
	//return value
}
