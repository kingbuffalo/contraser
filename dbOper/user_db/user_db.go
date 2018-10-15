package user_db

import (
	"encoding/json"
	"github.com/kingbuffalo/seelog"
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"buffalo/king/dbOper/playerInfo_db"
)

type UserVO struct {
	OpenId string `json:"open_id" gorm:primary_key`

	PlayerId int32  `json:"player_id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

func (UserVO) TableName() string {
	return "user"
}
func getUserKey(open_id string) string {
	return "s:sg:playerInfo:" + open_id
}

func newUser(player_id int32, open_id, name, url string) UserVO {
	return UserVO{
		PlayerId: player_id,
		Name:     name,
		Url:      url,
		OpenId:   open_id,
	}
}

func GetUser(open_id, name, url string) (*UserVO, error) {
	seelog.Trace("GetUser")
	key := getUserKey(open_id)
	rep := gameutil.GGet(key)
	var user UserVO
	if rep == nil {
		db := gameutil.GetDB()
		db.Where("open_id=?", open_id).First(&user)
		if user.PlayerId != 0 {
			err := user.Save(false)
			if err != nil {
				return nil, err
			}
			return &user, nil
		}
	}

	if rep == nil {
		player_id := gameutil.GIncr(enumErrCode.PLAYER_ID_INC_KEY)
		user := newUser(int32(player_id), open_id, name, url)
		err := user.createDbRecode()
		if err != nil {
			return nil, err
		}
		_, err = playerInfo_db.CreatePlayerInfoVO(int32(player_id), open_id, name, url)
		return &user, nil
	}
	err := json.Unmarshal(rep, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (user *UserVO) createDbRecode() error {
	key := getUserKey(user.OpenId)
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	gameutil.GSet(key, b)

	db := gameutil.GetDB()
	db.Create(user)

	return nil
}

func (user *UserVO) Save(bSaveToMysql bool) error {
	key := getUserKey(user.OpenId)
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	gameutil.GSet(key, b)

	if bSaveToMysql {
		db := gameutil.GetDB()
		db.Save(user)
	}

	return nil
}
