package main

import (
	"buffalo/king/common"
	"buffalo/king/king"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/kingbuffalo/seelog"
	"net/url"
	"os"
	"time"
)

var createNewPlayer bool = true

func sendFakeCmd(times int) []byte {
	if createNewPlayer {
		f := newerIdMapFunc[times]
		if f != nil {
			pb, str := f()
			id := common.GetProtoId(str)
			b, _ := common.Pb2ByteArrClient(playerId, id, pb)
			return b
		}
		return nil

	} else {
		f := idMapFunc[times]
		if f != nil {
			pb, str := f()
			id := common.GetProtoId(str)
			b, _ := common.Pb2ByteArrClient(playerId, id, pb)
			return b
		}
		return nil
	}
}

var playerId int = 20

func login() (proto.Message, string) {
	seelog.Trace("login")
	return &king.C2S_Login{
		OpenId: proto.String("cds062"),
		Name:   proto.String("cds"),
		Url:    proto.String("url"),
	}, "C2S_Login"
}

func heroRecruit() (proto.Message, string) {
	seelog.Trace("heroRecruit")
	return &king.C2S_HeroRecruit{
		Type: proto.Int32(1),
	}, "C2S_HeroRecruit"

}

func getInitGame() (proto.Message, string) {
	seelog.Trace("getInitGame")
	return &king.C2S_GetInitGame{}, "C2S_GetInitGame"
}

//func attackCity() (proto.Message, string) {
//seelog.Trace("attackCity")
//return &king.C2S_AttackCity{
//MapCityIdx: proto.Int32(101),
//ArmyId:     proto.Int32(101),
//}, "C2S_AttackCity"
//}

//func getMapCitys() (proto.Message, string) {
//seelog.Trace("getMapCitys")
//return &king.C2S_GetMapCitys{
//MapId: proto.Int32(1),
//}, "C2S_GetMapCitys"
//}
//func addTroops() (proto.Message, string) {
//seelog.Trace("addTroops")
////TODO
//return &king.C2S_AddTroops{
//HeroId:    proto.Int32(1001),
//AddTroops: proto.Int32(2),
//}, "C2S_AddTroops"
//}

func buyHero() (proto.Message, string) {
	seelog.Trace("BuyHero")
	return &king.C2S_BuyHero{
		ShopId: proto.Int32(1),
	}, "C2S_BuyHero"
}

//func collectRes() (proto.Message, string) {
//seelog.Trace("CollectRes")
//return &king.C2S_CollectRes{
//BuildIdx: proto.Int32(1),
//}, "C2S_CollectRes"
//}
func enQueue() (proto.Message, string) {
	seelog.Trace("EnQueue")
	return &king.C2S_EnQueue{
		HeroId1: proto.Int32(1),
		HeroId2: proto.Int32(1),
		HeroId3: proto.Int32(1),
		ArmyId:  proto.Int32(1),
	}, "C2S_EnQueue"
}
func enQueueSingle() (proto.Message, string) {
	seelog.Trace("EnQueueSingle")
	return &king.C2S_EnQueueSingle{
		HeroId: proto.Int32(1),
		Pos:    proto.Int32(1),
		ArmyId: proto.Int32(1),
	}, "C2S_EnQueueSingle"
}
func freeHeroRecruit() (proto.Message, string) {
	seelog.Trace("FreeHeroRecruit")
	return &king.C2S_FreeHeroRecruit{}, "C2S_FreeHeroRecruit"
}

//func getAllBattleResult() (proto.Message, string) {
//seelog.Trace("GetAllBattleResult")
//return &king.C2S_GetAllBattleResult{}, "C2S_GetAllBattleResult"
//}
//func getBattleResultDetail() (proto.Message, string) {
//seelog.Trace("GetBattleResultDetail")
//return &king.C2S_GetBattleResultDetail{
//BattleResultId: proto.Int32(1),
//}, "C2S_GetBattleResultDetail"
//}
//func getRankInfoList() (proto.Message, string) {
//seelog.Trace("GetRankInfoList")
//return &king.C2S_GetRankInfoList{}, "C2S_GetRankInfoList"
//}
/*
func getRedPoint() (proto.Message, string) {
	seelog.Trace("GetRedPoint")
	return &king.C2S_GetRedPoint{}, "C2S_GetRedPoint"
}
func getSelfRankInfo() (proto.Message, string) {
	seelog.Trace("GetSelfRankInfo")
	return &king.C2S_GetSelfRankInfo{}, "C2S_GetSelfRankInfo"
}
func getSignInfo() (proto.Message, string) {
	seelog.Trace("GetSignInfo")
	return &king.C2S_GetSignInfo{}, "C2S_GetSignInfo"
}
func getTopRankInfoList() (proto.Message, string) {
	seelog.Trace("GetTopRankInfoList")
	return &king.C2S_GetTopRankInfoList{}, "C2S_GetTopRankInfoList"
}
func heart() (proto.Message, string) {
	seelog.Trace("Heart")
	return &king.C2S_Heart{}, "C2S_Heart"
}
func levelupBuilding() (proto.Message, string) {
	seelog.Trace("LevelupBuilding")
	return &king.C2S_LevelupBuilding{
		BuildIdx: proto.Int32(102),
	}, "C2S_LevelupBuilding"
}
func oneKeySetQueue() (proto.Message, string) {
	seelog.Trace("OneKeySetQueue")
	return &king.C2S_OneKeySetQueue{}, "C2S_OneKeySetQueue"
}
func ping() (proto.Message, string) {
	seelog.Trace("Ping")
	return &king.C2S_Ping{
		ClientTimestamp: proto.Int32(1234),
	}, "C2S_Ping"
}
func reConnect() (proto.Message, string) {
	seelog.Trace("ReConnect")
	return &king.C2S_ReConnect{
		OpenId:   proto.String("567890"),
		PlayerId: proto.Int32(1),
		Token:    proto.String("ygfadghjsfghduf"),
	}, "C2S_ReConnect"
}

func unlockBuilding() (proto.Message, string) {
	seelog.Trace("UnlockBuilding")
	return &king.C2S_UnlockBuilding{
		BuildIdx: proto.Int32(102),
	}, "C2S_UnlockBuilding"
}
func timerRefreshMapCitys() (proto.Message, string) {
	seelog.Trace("TimerRefreshMapCitys")
	return &king.C2S_TimerRefreshMapCitys{
		MapId: proto.Int32(1),
	}, "C2S_TimerRefreshMapCitys"
}
func setDefArmy() (proto.Message, string) {
	seelog.Trace("SetDefArmy")
	return &king.C2S_SetDefArmy{
		ArmyId: proto.Int32(1),
		CityId: proto.Int32(1),
	}, "C2S_SetDefArmy"
}
func refreshMapCitys() (proto.Message, string) {
	seelog.Trace("RefreshMapCitys")
	return &king.C2S_RefreshMapCitys{
		MapId: proto.Int32(1),
	}, "C2S_RefreshMapCitys"
}
func sign() (proto.Message, string) {
	seelog.Trace("Sign")
	return &king.C2S_Sign{
		SignType: proto.Int32(1),
	}, "C2S_Sign"
}


func attackUnlockBuilding() (proto.Message, string) {
	seelog.Trace("attackUnlockBuilding")
	return &king.C2S_AttackUnlockBuilding{
		BuildIdx: proto.Int32(110),
		ArmyId:   proto.Int32(101),
	}, "C2S_AttackUnlockBuilding"
}

func getBuildings() (proto.Message, string) {
	//time.Sleep(time.Second * 20)
	seelog.Trace("getBuildings")
	return &king.C2S_GetBuildings{
		CityId: proto.Int32(1),
	}, "C2S_GetBuildings"
}

func getShopInfo() (proto.Message, string) {
	seelog.Trace("getShopInfo")
	return &king.C2S_GetShopInfo{}, "C2S_GetShopInfo"
}

func buyItem() (proto.Message, string) {
	seelog.Trace("buyItem")
	return &king.C2S_BuyItem{
		ShopId: proto.Int32(1101),
		Count:  proto.Int32(1),
	}, "C2S_BuyItem"
}

func getHeroRecruitInfo() (proto.Message, string) {
	seelog.Trace("getHeroRecruitInfo")
	return &king.C2S_GetHeroRecruitInfo{}, "C2S_GetHeroRecruitInfo"
}

func getPveInfoId() (proto.Message, string) {
	seelog.Trace("getPveInfoId")
	return &king.C2S_GetPVELevelId{
		MapId: proto.Int32(1),
	}, "C2S_GetPVELevelId"
}


func getDailyTask() (proto.Message, string) {
	seelog.Trace("getDailyTask")
	return &king.C2S_GetDailyTask{}, "C2S_GetDailyTask"
}
*/

func conFight() (proto.Message, string) {
	seelog.Trace("conFight")
	return &king.C2S_ContinueFight{
		SkillId1: proto.Int32(0),
		SkillId2: proto.Int32(0),
		SkillId3: proto.Int32(0),
	}, "C2S_ContinueFight"
}
func fightPVE() (proto.Message, string) {
	seelog.Trace("fightPVE")
	return &king.C2S_FightPVE{
		LevelId: proto.Int32(1),
		ArmyId:  proto.Int32(1),
	}, "C2S_FightPVE"
}

func enQueueSingle1() (proto.Message, string) {
	seelog.Trace("enQueueSingle")
	return &king.C2S_EnQueueSingle{
		ArmyId: proto.Int32(1),
		Pos:    proto.Int32(1),
		HeroId: proto.Int32(1),
	}, "C2S_EnQueueSingle"
}

func enQueueSingle2() (proto.Message, string) {
	seelog.Trace("enQueueSingle")
	return &king.C2S_EnQueueSingle{
		ArmyId: proto.Int32(1),
		Pos:    proto.Int32(2),
		HeroId: proto.Int32(2),
	}, "C2S_EnQueueSingle"
}

func enQueueSingle3() (proto.Message, string) {
	seelog.Trace("enQueueSingle")
	return &king.C2S_EnQueueSingle{
		ArmyId: proto.Int32(1),
		Pos:    proto.Int32(3),
		HeroId: proto.Int32(3),
	}, "C2S_EnQueueSingle"
}

var newerIdMapFunc = map[int](func() (proto.Message, string)){
	0:  login,
	1:  getInitGame,
	2:  heroRecruit,
	3:  heroRecruit,
	4:  heroRecruit,
	5:  heroRecruit,
	6:  heroRecruit,
	7:  heroRecruit,
	8:  heroRecruit,
	9:  heroRecruit,
	10: heroRecruit,
	11: heroRecruit,
	12: heroRecruit,
	13: heroRecruit,
	14: heroRecruit,
	15: heroRecruit,
	16: heroRecruit,
	17: enQueueSingle1,
	18: enQueueSingle2,
	19: enQueueSingle3,
	20: fightPVE,
	21: conFight,
}

var idMapFunc = map[int](func() (proto.Message, string)){
	0: login,
	1: getInitGame,
	2: conFight,
	//2: getMapCitys,
	//3: getDailyTask,
	//3: buyItem,
	//3: fightPVE,
	//4: addTroops,
	//3: refreshMapCitys,
	//4: getPveInfoId,
	//2: fightPVE,
	//3: getPveInfoId,
	//2: getMapCitys,
	//3: getBuildings,
	//4: getShopInfo,
	//5: buyItem,
	//6: heroRecruit,
	//7: getHeroRecruitInfo,
	//8: getPveInfoId,
	//3: attackCity,
}

//var idMapFunc = map[int](func() (proto.Message, string)){
//0:  login,
//1:  getInitGame,
//2:  heroRecruit,
//3:  getMapCitys,
//4:  attackCity,
//5:  addTroops,
//6:  buyHero,
//7:  collectRes,
//8:  enQueue,
//9:  enQueueSingle,
//10: reConnect,
//11: ping,
//12: freeHeroRecruit,
//13: getAllBattleResult,
//14: getBattleResultDetail,
//15: getRankInfoList,
//16: getRedPoint,
//17: getSelfRankInfo,
//18: getSignInfo,
//19: getTopRankInfoList,
//20: heart,
//21: unlockBuilding,
//22: oneKeySetQueue,
//23: refreshMapCitys,
//24: setDefArmy,
//25: sign,
//26: timerRefreshMapCitys,
//27: levelupBuilding,
//}

func main() {
	if len(os.Args) < 3 {
		panic("need addr from Args[3]")
	}
	addr := os.Args[3]
	times := 0
L:
	for {
		select {
		case <-time.After(time.Second):
			u := url.URL{Scheme: "ws", Host: addr, Path: "/"}
			c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				panic(err)
			}

			xx := sendFakeCmd(times)
			if xx == nil {
				break L
			}
			if err := c.WriteMessage(websocket.BinaryMessage, xx); err != nil {
				seelog.WarnWE("c.WriteMessage,err=", err)
				break L
			}
			_ = c.Close()
		}
		times++
	}

	time.Sleep(time.Second * 2)
	fmt.Println("all run")
}
