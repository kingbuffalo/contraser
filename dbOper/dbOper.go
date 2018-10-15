package dbOper

import (
	"buffalo/king/common/gameutil"
	"encoding/json"
	"fmt"
	"github.com/kingbuffalo/seelog"
	"strconv"
)

///////////////////////////////////1k///////////////////////////////////////////////////// begin
type Player1key interface {
	SetPlayerId(playerId int32)
	BExistInMysql() bool
	GetPlayerId() int32
}

func SaveVO(vo Player1key, keyFmt string, bSaveToMysql bool) error {
	return save_helper(vo, keyFmt, bSaveToMysql, false)
}

func SaveNewVO(vo Player1key, keyFmt string) error {
	return save_helper(vo, keyFmt, true, true)
}

func save_helper(vo Player1key, keyFmt string, bSaveToMysql bool, bNew bool) error {
	key := fmt.Sprintf(keyFmt, vo.GetPlayerId())
	b, err := json.Marshal(vo)
	if err != nil {
		return err
	}
	gameutil.GSet(key, b)
	if bSaveToMysql {
		db := gameutil.GetDB()
		if bNew {
			db.Create(vo)
		} else {
			db.Model(vo).Updates(vo)
		}
	}
	return nil
}

func GetOneKeyVO(keyFmt string, playerId int32, vo Player1key) error {
	key := fmt.Sprintf(keyFmt, playerId)
	rep := gameutil.GGet(key)
	if rep == nil {
		db := gameutil.GetDB()
		db.Where("player_id=?", playerId).Find(vo)
		if !vo.BExistInMysql() {
			vo.SetPlayerId(playerId)
			if err := save_helper(vo, keyFmt, true, true); err != nil {
				return err
			}
			return nil
		} else {
			err := SaveVO(vo, keyFmt, false)
			return err
		}
	}
	err := json.Unmarshal(rep, &vo)
	if err != nil {
		seelog.Trace(" Unmarshal error")
		return err
	}
	return nil
}

///////////////////////////////////1k///////////////////////////////////////////////////// end

///////////////////////////////////2k///////////////////////////////////////////////////// end
type Player2key interface {
	SetPlayerIdId(playerId, id int32)
	BExistInMysql() bool
	GetPlayerIdId() (int32, int32)
	BEqual(playerId, id int32) bool
	Clone() Player2key
	SetValue(Player2key)
}

func Get2KVOs(keyFmt string, playerId int32, vos []Player2key, vo Player2key) error {
	key := fmt.Sprintf(keyFmt, playerId)
	keyMapJson := gameutil.HGetAll(key)
	if keyMapJson == nil {
		db := gameutil.GetDB()
		db.Where("player_id=?", playerId).Find(&vos)
		for _, v := range vos {
			if err := Save2VO(v, keyFmt, false); err != nil {
				return err
			}
		}
		return nil
	}
	idx := 0
	for _, v := range keyMapJson {
		cloneVO := vo.Clone()
		if err := json.Unmarshal([]byte(v), cloneVO); err != nil {
			seelog.WarnWE("json.Unmarshal error=", err)
			return err
		}
		vos[idx] = cloneVO
		idx++
	}
	return nil
}

func Get2kVO(keyFmt string, playerId, id int32, vos []Player2key, vo Player2key) error {
	key := fmt.Sprintf(keyFmt, playerId)
	fieldStr := strconv.Itoa(int(id))
	b := gameutil.HGet(key, fieldStr)
	if b == nil {
		if !gameutil.GExists(key) {
			if err := Get2KVOs(keyFmt, playerId, vos, vo); err != nil {
				return err
			} else {
				for _, v := range vos {
					if v.BEqual(playerId, id) {
						vo.SetValue(v)
						return nil
					}
				}
			}
		}
		return nil
	}
	if err := json.Unmarshal(b, vo); err != nil {
		seelog.WarnWE("json.Unmarshal error=", err)
		return err
	}
	return nil
}

func save_2_helper(vo Player2key, keyFmt string, bSaveToMysql bool, bNew bool) error {
	v1, v2 := vo.GetPlayerIdId()
	key := fmt.Sprintf(keyFmt, v1)
	field := strconv.Itoa(int(v2))
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
			db.Model(vo).Updates(vo)
		}
	}
	return nil
}

func Save2VO(vo Player2key, keyFmt string, bSaveToMysql bool) error {
	return save_2_helper(vo, keyFmt, bSaveToMysql, false)
}
