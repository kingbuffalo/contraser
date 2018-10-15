package battle

import (
	"buffalo/king/cfgmgr/cfgPVELevelMgr"
	"buffalo/king/dbOper/hero_db"
	"github.com/kingbuffalo/seelog"
)

func NewBattlePVP(atk_heroVOs, def_heroVOs []*hero_db.HeroVO) *Battle {
	seelog.Trace("NewBattlePVP", atk_heroVOs, def_heroVOs)
	b := &Battle{
		round:    0,
		subRound: 0,
	}
	b.aq = newQueue(b, atk_heroVOs, true)
	b.dq = newQueue(b, def_heroVOs, false)
	b.aq.oppo_q = b.dq
	b.dq.oppo_q = b.aq
	return b
}
func NewBattlePVELevel(heroVOs []*hero_db.HeroVO, cfg *cfgPVELevelMgr.CfgPVELevel) *Battle {
	seelog.Trace("NewBattlePVELevel", heroVOs, cfg)
	var b Battle
	//attack_q := newQueuePVELevelPlayer(&b, heroVOs)
	attack_q := newQueue(&b, heroVOs, true)
	defence_q := newQueuePVENpc(&b, cfg)
	attack_q.oppo_q = defence_q
	defence_q.oppo_q = attack_q
	b.aq = attack_q
	b.dq = defence_q
	/*
		b.mapCity = &mapCity_t{
			durability: 0,
			level:      1,
			pos:        enumErrCode.BATTLE_POS_DEFENCER_CITY,
		}*/
	return &b
}
