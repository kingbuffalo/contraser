package battleResult_db

import (
	"buffalo/king/common/gameutil"
	"buffalo/king/dbOper/commonVO"
	"buffalo/king/king"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	"strconv"
)

type HeroResultVO struct {
	Pos        int32 `json:"pos"`
	HeroId     int32 `json:"heroId"`
	Troops     int32 `json:"troops"`
	InitTroops int32 `json:"init_troops"`
	Level      int32 `json:"level"`
	Exp        int32 `json:"exp"`
}

func NewRobAbleResVO(wood, mineral, grain, coin int32) *commonVO.RobAbleResVO {
	return &commonVO.RobAbleResVO{
		Wood:    wood,
		Mineral: mineral,
		Grain:   grain,
		Coin:    coin,
	}
}

type BattleEventVO struct {
	Id   int32 `json:"id"`
	Pos  int   `json:"pos"`
	Type int   `json:"type"`
}

type BattleTroopsChangeVO struct {
	Pos   int32 `json:"pos"`
	Value int32 `json:"value"`
}

type BattleBuffVO struct {
	Pos      int32 `json:"pos"`
	BuffId   int32 `json:"buffId"`
	EndRound int32 `json:"endRound"`
}

type BattleActionVO struct {
	AttackPos             int                     `json:"attack_pos"`
	TargetPos             int                     `json:"target_pos"`
	SkillId               int32                   `json:"skill_id"`
	BattleTroopsChangeVOs []*BattleTroopsChangeVO `json:"battleTroopsChangeVOs"`
	BattleBuffVOs         []*BattleBuffVO         `json:"battleBuffVOs"`
	EventIds              []int32                 `json:"event_id"`
}

func (ba *BattleActionVO) GenBattleTroopsChange() []*king.TargetBeAction {
	ret := make([]*king.BattleTroopsChange, len(ba.BattleTroopsChangeVOs))
	for i, v := range ba.BattleTroopsChangeVOs {
		btv := &king.BattleTroopsChange{
			Pos:   proto.Int32(v.Pos),
			Value: proto.Int32(v.Value),
		}
		ret[i] = btv
	}
	return ret
}

type RoundVO struct {
	Round           int32               `json:"attack_pos"`
	BattleActionVOs [](*BattleActionVO) `json:"battle_actionvo_s"`
	BattleEventVOs  [](*BattleEventVO)  `json:"battle_eventvo_s"`
}

type BattleCityInfoVO struct {
	Level          int32 `json:"level"`
	InitDurability int32 `json:"init_durability"`
	Durability     int32 `json:"durability"`
}

func getKeyStr(playerId int32) string {
	return fmt.Sprintf("h:sg:battleResult:%d", playerId)
}

func NewRoundVO(round int32) *RoundVO {
	r := &RoundVO{
		Round:           round,
		BattleEventVOs:  make([](*BattleEventVO), 0),
		BattleActionVOs: make([](*BattleActionVO), 0),
	}
	return r
}

func (rr *RoundVO) AddAction(attackPos int, targetPos int, skillId int32, battleTroopsChangeVOs []*BattleTroopsChangeVO, bb []*BattleBuffVO,
	eventPoss []int, eventType int, eventIds []int32) {
	seelog.Tracef("r addAction,apos=%d,tpos=%d,skillId=%d,value=%d,eventPoss=%v,eventType=%d,eventIds=%v", attackPos, targetPos, skillId, 0, eventPoss, eventType, eventIds)
	if len(eventPoss) != 0 {
		for i, _ := range eventPoss {
			be := &BattleEventVO{
				Id:   eventIds[i],
				Pos:  eventPoss[i],
				Type: eventType,
			}
			rr.BattleEventVOs = append(rr.BattleEventVOs, be)
		}
	}
	ba := &BattleActionVO{
		AttackPos:             attackPos,
		TargetPos:             targetPos,
		SkillId:               skillId,
		BattleTroopsChangeVOs: battleTroopsChangeVOs,
		BattleBuffVOs:         bb,
		EventIds:              eventIds,
	}
	rr.BattleActionVOs = append(rr.BattleActionVOs, ba)
}

type BattleResultVO struct {
	PlayerId        int32                  `json:"player_id"`
	BattleResultId  int32                  `json:"battle_result_id"`
	ArmyId          int32                  `json:"army_id"`
	HeroResultVOs   []*HeroResultVO        `json:"hero_resultvos"`
	Result          int32                  `json:"result"`
	Res             *commonVO.RobAbleResVO `json:"res"`
	Rounds          []*RoundVO             `json:"rounds"`
	MapCityIdx      int32                  `json:"map_city_idx"`
	DecampTimestamp int32                  `json:"decamp_timestamp"`
	Timestamp       int32                  `json:"timestamp"`
	CityInfoVO      *BattleCityInfoVO      `json:"battle_city_info"`
	AtkId           int32                  `json:"atk_id"`
	DefId           int32                  `json:"def_id"`
}

func getKey(playerId int32) string {
	return fmt.Sprintf("h:sg:BattleResult:%d", playerId)
}

func getRedisKey(playerId, battleResultId int32) (string, string) {
	key := getKey(playerId)
	field := strconv.Itoa(int(battleResultId))
	return key, field

}

func (b *BattleResultVO) getRedisKey() (string, string) {
	return getRedisKey(b.PlayerId, b.BattleResultId)
}

func NewBattleResultVO(playerId int32, r []*RoundVO, mapCityIdx int32) *BattleResultVO {
	return &BattleResultVO{
		PlayerId:   playerId,
		Rounds:     r,
		MapCityIdx: mapCityIdx,
	}
}

func GetBattleResultVO(playerId, battleResultId int32) (*BattleResultVO, error) {
	key, field := getRedisKey(playerId, battleResultId)
	b := gameutil.HGet(key, field)
	if b == nil {
		return nil, errors.New("not data exist")
	}
	var br BattleResultVO
	err := json.Unmarshal(b, &br)
	if err != nil {
		return nil, err
	}

	return &br, nil

}
func GetBattleResultVOs(playerId int32) ([](*BattleResultVO), error) {
	key := getKey(playerId)
	keyMapJson := gameutil.HGetAll(key)
	if keyMapJson == nil {
		return [](*BattleResultVO){}, nil
	}
	ret := make([]*BattleResultVO, len(keyMapJson))
	idx := 0
	for _, v := range keyMapJson {
		var vo BattleResultVO
		if err := json.Unmarshal([]byte(v), &vo); err != nil {
			return nil, err
		}
		ret[idx] = &vo
		idx++
	}
	return ret, nil
}

func (br *BattleResultVO) Save() error {
	key, field := br.getRedisKey()
	b, err := json.Marshal(br)
	if err != nil {
		return err
	}
	gameutil.HSet(key, field, b)

	return nil
}

func (b *BattleResultVO) GenBriefPushData() *king.BriefBattleResult {
	res := &king.RobableRes{
		Wood:    proto.Int32(b.Res.Wood),
		Mineral: proto.Int32(b.Res.Mineral),
		Grain:   proto.Int32(b.Res.Grain),
		Coin:    proto.Int32(b.Res.Coin),
	}

	br := &king.BriefBattleResult{
		ArmyId:         proto.Int32(b.ArmyId),
		Result:         proto.Int32(b.Result),
		Res:            res,
		MapCityIdx:     proto.Int32(b.MapCityIdx),
		Timestamp:      proto.Int32(b.Timestamp),
		BattleResultId: proto.Int32(b.BattleResultId),
		AtkId:          proto.Int32(b.AtkId),
		DefId:          proto.Int32(b.DefId),
	}
	return br
}

func (b *BattleResultVO) GenDetailPushData() *king.S2C_GetBattleResultDetail {
	ret := &king.S2C_GetBattleResultDetail{
		BattleResult: b.GenPushData(),
	}
	return ret
}

func (b *BattleResultVO) GenPushData() *king.BattleResult {
	rounds := make([](*king.Round), len(b.Rounds))
	for i, v := range b.Rounds {
		p_battleActions := make([](*king.BattleAction), len(v.BattleActionVOs))
		for ii, vv := range v.BattleActionVOs {
			btcvos := vv.GenBattleTroopsChange()
			p_battleActions[ii] = new(king.BattleAction)
			p_battleActions[ii].AttackPos = proto.Int32(int32(vv.AttackPos))
			p_battleActions[ii].TargetPos = proto.Int32(int32(vv.TargetPos))
			p_battleActions[ii].SkillId = proto.Int32(int32(vv.SkillId))
			p_battleActions[ii].BattleTroopsChanges = btcvos
			p_battleActions[ii].EventIds = vv.EventIds
		}
		p_battleEvents := make([](*king.BattleEvent), len(v.BattleEventVOs))
		for ii, vv := range v.BattleEventVOs {
			p_battleEvents[ii] = new(king.BattleEvent)
			p_battleEvents[ii].Id = proto.Int32(int32(vv.Id))
			p_battleEvents[ii].Pos = proto.Int32(int32(vv.Pos))
			p_battleEvents[ii].Type = proto.Int32(int32(vv.Type))
		}
		rounds[i] = new(king.Round)
		rounds[i].Round = proto.Int32(int32(v.Round))
		rounds[i].BattleEvents = p_battleEvents
		rounds[i].BattleActions = p_battleActions
	}

	heroResults := make([](*king.HeroResult), len(b.HeroResultVOs))
	for i, v := range b.HeroResultVOs {
		hr := &king.HeroResult{
			Pos:        proto.Int32(v.Pos),
			HeroId:     proto.Int32(v.HeroId),
			Troops:     proto.Int32(v.Troops),
			InitTroops: proto.Int32(v.InitTroops),
			Level:      proto.Int32(1),
			Exp:        proto.Int32(0),
		}
		heroResults[i] = hr
	}

	res := &king.RobableRes{
		Wood:    proto.Int32(100),
		Mineral: proto.Int32(101),
		Grain:   proto.Int32(102),
		Coin:    proto.Int32(103),
	}

	ba := &king.BattleCityInfo{
		Durability:     proto.Int32(b.CityInfoVO.Durability),
		InitDurability: proto.Int32(b.CityInfoVO.InitDurability),
		Level:          proto.Int32(b.CityInfoVO.Level),
	}

	br := &king.BattleResult{
		Res:             res,
		HeroResults:     heroResults,
		Result:          proto.Int32(b.Result),
		Rounds:          rounds,
		ArmyId:          proto.Int32(b.ArmyId),
		MapCityIdx:      proto.Int32(b.MapCityIdx),
		DecampTimestamp: proto.Int32(b.DecampTimestamp),
		BattleResultId:  proto.Int32(b.BattleResultId),
		BattleCityInfo:  ba,
		AtkId:           proto.Int32(b.AtkId),
		DefId:           proto.Int32(b.DefId),
	}
	return br
}
