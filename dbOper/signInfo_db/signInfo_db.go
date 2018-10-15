package signInfo_db

import (
	"encoding/json"
	//"errors"
	"github.com/kingbuffalo/seelog"
	"strconv"
	"time"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
)

type SignInfoVO struct {
	PlayerId          int32 `gorm:"primary_key"`
	LastWeekSignTime  int32
	LastMonthSignTime int32
	WeekSignDays      int32
	MonthSignDays     int32
	CreateTimestamp   int32
	WeekId            int32
	MonthId           int32
	dirty             bool
}

func getSignInfoKey(player_id int32) string {
	return "signinfo:" + strconv.Itoa(int(player_id))
}

func newSignInfo(player_id int32) SignInfoVO {
	ct := time.Now()
	ctimestamp := ct.Unix()
	month := ct.Month()
	return SignInfoVO{
		PlayerId:          player_id,
		LastWeekSignTime:  0,
		LastMonthSignTime: 0,
		WeekSignDays:      0,
		MonthSignDays:     0,

		CreateTimestamp: int32(ctimestamp),
		WeekId:          1,
		MonthId:         int32(month),
		dirty:           true,
	}
}

func (s *SignInfoVO) BSignToday() (bWeekSignToday bool, bMonthSignToday bool) {
	bWeekSignToday = gameutil.BToday(s.LastWeekSignTime)
	bMonthSignToday = gameutil.BToday(s.LastMonthSignTime)
	return
}
func (s *SignInfoVO) SaveVO(bSaveToMysql bool) error {
	if !s.dirty {
		return nil
	}
	key := getSignInfoKey(s.PlayerId)
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	gameutil.GSet(key, b)
	return nil
}

func GetSignInfoById(player_id int32) (*SignInfoVO, error) {
	key := getSignInfoKey(player_id)
	rep := gameutil.GGet(key)
	if rep == nil { //&& err != nil
		signInfo := newSignInfo(player_id)
		lerr := signInfo.SaveVO(true)
		if lerr != nil {
			seelog.Trace("[signInfo_db:: GetSignInfoById] Set sign info err")
			return nil, lerr
		}
		seelog.Trace("[signInfo_db:: GetSignInfoById] new data")
		//这里的SignInfoVO 数据要进行一个处理
		return &signInfo, nil
	}

	var signInfo SignInfoVO
	lerr := json.Unmarshal(rep, &signInfo)
	if lerr != nil {
		seelog.Trace("[signInfo_db:: GetSignInfoById] Unmarshal SignInfoVO error")
		return nil, lerr
	}
	signInfo.update()
	seelog.Trace("[signInfo_db:: GetSignInfoById]ok de ya")
	return &signInfo, nil
}

func (s *SignInfoVO) SignToday(signType int32) {
	now := int32(time.Now().Unix())
	s.dirty = true
	if signType == enumErrCode.SIGN_TYPE_WEEK {
		s.LastWeekSignTime = now
		s.WeekSignDays++
	} else if signType == enumErrCode.SIGN_TYPE_MONTH {
		s.LastMonthSignTime = now
		s.MonthSignDays++
	}
}

func (s *SignInfoVO) update() {
	nowTime := time.Now()
	createTime := time.Unix(int64(s.CreateTimestamp), 0)
	createY, createM, createD := createTime.Date()
	loc := createTime.Location()
	createTimeAt0 := time.Date(createY, createM, createD, 0, 0, 0, 0, loc)
	d := nowTime.Sub(createTimeAt0)
	h := d.Hours()
	days := int(h / 24)
	weekId := int32(days / 7)
	if days%7 != 0 {
		weekId++
	}
	if weekId != s.WeekId {
		s.WeekId = weekId
		s.WeekSignDays = 0
		s.dirty = true
	}
	nowMonth := nowTime.Month()
	if s.MonthId != int32(nowMonth) {
		s.MonthId = int32(nowMonth)
		s.MonthSignDays = 0
		s.dirty = true
	}
}
