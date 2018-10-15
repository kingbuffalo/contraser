package email_db

import (
	"encoding/json"
	"fmt"
	"github.com/kingbuffalo/seelog"
	"strconv"
	"buffalo/king/common/gameutil"
	//"buffalo/king/dbOper"
)

const (
	keyFmt = "h:sg:EmailVOKey:%d"
)

type EmailVO struct {
	PlayerId int32 `json:"player_id" gorm:"primary_key"`
	Id       int32 `json:"id" gorm:"primary_key"`

	Type    int32     `json:"type"`
	Content string    `json:"content"`
	Gift    [][]int32 `json:"gift"`
}

/*
func (vo *EmailVO) SetPlayerIdId(playerId, id int32) {
	vo.PlayerId = playerId
	vo.Id = id
}
func (vo *EmailVO) BExistInMysql() bool {
	return vo.PlayerId != 0
}
func (vo *EmailVO) GetPlayerIdId() (int32, int32) {
	return vo.PlayerId, vo.Id
}
func (vo *EmailVO) BEqual(playerId, id int32) bool {
	return vo.PlayerId == playerId && vo.Id == id
}

func getEmailVOKey(playerId int32) string {
	return fmt.Sprintf("h:sg:EmailVOKey:%d", playerId)
}

func (vo *EmailVO) Clone() dbOper.Player2key {
	return &EmailVO{
		PlayerId: vo.PlayerId,
		Id:       vo.Id,
		Type:     vo.Type,
		Content:  vo.Content,
		Gift:     vo.Gift,
	}
}

func (vo *EmailVO) SetValue(v dbOper.Player2key) {
	ev := v.(*EmailVO)
	vo.PlayerId = ev.PlayerId
	vo.Id = ev.Id
	vo.Type = ev.Type
	vo.Gift = ev.Gift
	vo.Content = ev.Content
}

func GetEmailVOs(playerId int32) ([](*EmailVO), error) {
	vos := make([](dbOper.Player2key), 0)
	vo := &EmailVO{}
	err := dbOper.Get2KVOs(keyFmt, playerId, vos, vo)
	if err != nil {
		return nil, err
	}
	return vos, nil
}

*/

func GetEmailVOs(playerId int32) ([](*EmailVO), error) {
	key := fmt.Sprintf(keyFmt, playerId)
	keyMapJson := gameutil.HGetAll(key)
	if keyMapJson == nil {
		var vos []*EmailVO = [](*EmailVO){}
		db := gameutil.GetDB()
		db.Where("player_id=?", playerId).Find(&vos)
		for _, v := range vos {
			if err := v.SaveVO(false); err != nil {
				return nil, err
			}
		}
		return vos, nil
	}
	ret := make([](*EmailVO), len(keyMapJson))
	idx := 0
	for _, v := range keyMapJson {
		var vo EmailVO
		if err := json.Unmarshal([]byte(v), &vo); err != nil {
			seelog.WarnWE("json.Unmarshal error=", err)
			return nil, err
		}
		ret[idx] = &vo
		idx++
	}
	return ret, nil
}

func GetEmailVO(playerId, Id int32) (*EmailVO, error) {
	key := fmt.Sprintf(keyFmt, playerId)
	fieldStr := strconv.Itoa(int(Id))
	b := gameutil.HGet(key, fieldStr)
	if b == nil {
		if !gameutil.GExists(key) {
			if vos, err := GetEmailVOs(playerId); err != nil {
				return nil, err
			} else {
				for _, v := range vos {
					if v.Id == Id {
						return v, nil
					}
				}
			}
		} else {
			db := gameutil.GetDB()
			var vo EmailVO
			db.Where("player_id=? AND /*TODO Idinmysqlcolumn*/=?", playerId, Id).Find(&vo)
			if vo.PlayerId != 0 {
				err := vo.SaveVO(false)
				return &vo, err
			}
		}
		return nil, nil
	}
	var vo EmailVO
	if err := json.Unmarshal(b, &vo); err != nil {
		seelog.WarnWE("json.Unmarshal error=", err)
		return nil, err
	}
	return &vo, nil
}

func getEmailKey(playerId, emailId int32) (string, string) {
	return fmt.Sprintf("s:sg:EmailVO:%d", playerId), strconv.Itoa(int(emailId))
}

func (vo *EmailVO) getRedisKey() (string, string) {
	key, field := getEmailKey(vo.PlayerId, vo.Id)
	return key, field
}
func (vo *EmailVO) save_helper(bSaveToMysql bool, bNew bool) error {
	key, field := vo.getRedisKey()
	b, err := json.Marshal(vo)
	if err != nil {
		return err
	}
	gameutil.HSet(key, field, b)
	if bSaveToMysql {
		db := gameutil.GetDB()
		if bNew {
			db.Create(vo)
		} else {
			db.Save(vo)
		}
	}
	return nil
}

func (vo *EmailVO) SaveVO(bSaveToMysql bool) error {
	return vo.save_helper(bSaveToMysql, false)
}
