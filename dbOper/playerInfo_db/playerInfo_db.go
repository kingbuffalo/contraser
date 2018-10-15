package playerInfo_db

import (
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"buffalo/king/dbOper/army_db"
	"buffalo/king/dbOper/build_db"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kingbuffalo/seelog"
	"strconv"
)

const (
	new_player_cutter_id = 10101
	lc_KEY_ABC           = "set:sg:palyerIdSet"
)

type PlayerInfoVO struct {
	PlayerId int32 `gorm:"primary_key"`
	Name     string
	Url      string
	OpenId   string
	Coin     int32
	Diamond  int32
	Wood     int32
	Mineral  int32
	Grain    int32
	Level    int32

	//五星卡累计不获得次数
	FiveStarHeroRecruitTimes int32
	//每10次就会有特殊抽取的次数
	TenTimesHeroRecuritTimes int32
	//free
	FreeHeroRecruitTimestamp int32
}

func (PlayerInfoVO) TableName() string {
	return "player"
}
func getPlayerInfoKey(playerId int32) string {
	return "playerInfo:" + strconv.Itoa(int(playerId))
}

func newPlayerInfo(playerId int32, open_id, name, url string) PlayerInfoVO {
	gameutil.SAdd(lc_KEY_ABC, playerId)
	return PlayerInfoVO{
		PlayerId: playerId,
		Name:     name,
		Url:      url,
		OpenId:   open_id,
		Coin:     enumErrCode.NEW_BORN_PLAYER_COIN,
		Diamond:  enumErrCode.NEW_BORN_PLAYER_DIAMOND,
		Wood:     enumErrCode.NEW_BORN_PLAYER_WOOD,
		Mineral:  enumErrCode.NEW_BORN_PLAYER_MINERAL,
		Level:    enumErrCode.NEW_BORN_PLAYER_LEVEL,
	}
}

//数据库获取数据
func getPlayerInfoWithDb(playerId int32) (*PlayerInfoVO, error) {
	var playerInfo PlayerInfoVO
	db := gameutil.GetDB()
	db.Where("player_id=?", playerId).Find(&playerInfo)
	if playerInfo.PlayerId == 0 {
		errMsg := fmt.Sprintf("getPlayerInfoWithDb,playerId=%d", playerId)
		return nil, errors.New(errMsg)
	}
	return &playerInfo, nil
}

//newplayer db
func CreatePlayerInfoVO(playerId int32, open_id, name, url string) (PlayerInfoVO, error) {
	playerInfoVO := newPlayerInfo(playerId, open_id, name, url)
	db := gameutil.GetDB()
	db.Create(&playerInfoVO)
	err := playerInfoVO.Save()
	if err != nil {
		seelog.WarnWE("SetPlayerInfo err=", err)
	}
	build_db.NewPlayerBuildingVO(playerId)
	army_db.NewPlayerArmy(playerId)

	return playerInfoVO, nil
}

func GetPlayerInfoVOBy(playerId int32) (*PlayerInfoVO, error) {
	seelog.Trace("playerId:", playerId)
	key := getPlayerInfoKey(playerId)
	rep := gameutil.GGet(key)
	if rep == nil { //&& err == nil
		//quyer from mysql
		dPlayerInfo, derr := getPlayerInfoWithDb(playerId)
		if derr != nil {
			return nil, derr
		}
		//set 到 redis中
		err := dPlayerInfo.Save()
		return dPlayerInfo, err
	}
	var playerInfo PlayerInfoVO
	err := json.Unmarshal(rep, &playerInfo)
	if err != nil {
		seelog.Trace(" Unmarshal error")
		return nil, err
	}
	return &playerInfo, nil
}

func (playerInfo *PlayerInfoVO) Save() error {
	key := getPlayerInfoKey(playerInfo.PlayerId)
	b, err := json.Marshal(playerInfo)
	if err != nil {
		return nil
	}
	gameutil.GSet(key, b)

	db := gameutil.GetDB()
	db.Save(playerInfo)
	return nil
}

func SetMaxPlayId() {
	initAddPlayerIdToSet()
	var playerInfo PlayerInfoVO
	db := gameutil.GetDB()
	db.Last(&playerInfo)
	gameutil.GSetAllType(enumErrCode.PLAYER_ID_INC_KEY, playerInfo.PlayerId)
}

//随机获取num个玩家信息 //此处只需要playerId
func GetRandPlayerIds(num, playerId int32) ([]int32, error) {
	num++
	strArr := gameutil.SRandMemberN(lc_KEY_ABC, int64(num))
	if strArr != nil {
		ret := make([]int32, int(num))
		if len(strArr) >= int(num) {
			for i, str := range strArr {
				v, err := strconv.Atoi(str)
				if err != nil {
					return nil, err
				}
				ret[i] = int32(v)
			}
		} else {
			return nil, errors.New("SRandMemberN number < num")
		}

		theSamePidIdx := -1
		for i, v := range ret {
			if playerId == v {
				theSamePidIdx = i
			}
		}
		maxIdx := int(num) - 1
		if theSamePidIdx != -1 {
			ret[theSamePidIdx], ret[maxIdx] = ret[maxIdx], ret[theSamePidIdx]
		}

		return ret[:maxIdx], nil
	}
	return nil, errors.New("SRandMemberN error")
}

func base64Name(name string) string {
	return base64.StdEncoding.EncodeToString([]byte(name))
}

func initAddPlayerIdToSet() {
	db := gameutil.GetDB()
	rows, err := db.Raw("select player_id from player").Rows()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()
	if !rows.Next() {
		urlArr := []string{"https://wx.qlogo.cn/mmopen/vi_32/BBca90dxj0Q004YlOyoXk40TNLIib0QY2oiaogoichXprjiaib4IaYDffVsoAtiblVtktiatBAHrZtUfySibJgoqCRkSdA/132",
			"https://idleherores.heywoodsminiprogram.com/common/defaultIcon.png"}
		_, _ = CreatePlayerInfoVO(1, "open_id1", base64Name("name1"), urlArr[0])
		_, _ = CreatePlayerInfoVO(2, "open_id2", base64Name("name2"), urlArr[1])
		_, _ = CreatePlayerInfoVO(3, "open_id3", base64Name("name3"), urlArr[1])
		_, _ = CreatePlayerInfoVO(4, "open_id4", base64Name("name4"), urlArr[1])
		_, _ = CreatePlayerInfoVO(5, "open_id5", base64Name("name5"), urlArr[1])
		_, _ = CreatePlayerInfoVO(6, "open_id6", base64Name("name6"), urlArr[1])
		_, _ = CreatePlayerInfoVO(7, "open_id7", base64Name("name7"), urlArr[1])
	} else {
		pids := make([]interface{}, 0)
		for rows.Next() {
			var pid int
			if err := rows.Scan(&pid); err != nil {
				panic(err)
			}
			pids = append(pids, pid)
		}
		if len(pids) < int(enumErrCode.TOTAL_NEIGHBOR) {
			panic("len(members) < enumErrCode.TOTAL_NEIGHBOR")
		}

		gameutil.SAdd(lc_KEY_ABC, pids...)
	}
}
