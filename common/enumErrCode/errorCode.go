package enumErrCode

var ErrCode = map[int]string{
	1: "hhhhhhhhhh",
	//common error
	101: "通信读取数据出错",
	102: "协议长度不对",
	103: "调用rpc出错",
	104: "反序列化s2c_login出错",
	105: "此协议不存在",
	106: "此协议不存在2",
	107: "此协议已经关闭",
	108: "此协议不存在",
	111: "非推送压包错误",
	112: "推送压包错误",
	113: "报错了",
	//Login
	1001: "proto反序列化出错",
	1002: "后台出错",
	1003: "取用户信息出错",
	1004: "保存信号令出错",
	1005: "从后台那里取数据出错",
	1006: "鸡数据压包出错",
	1007: "后台返回错误码",
	1008: "open_id len = 0",
	//PlayerInfo
	1011: "proto反序列化出错",
	1012: "获取玩家信息出错",
	1013: "序列化出错",
	1014: "保存玩家信息出错",
	//GetSkill
	1021: "proto反序列化出错",
	1022: "获取玩定技能出错",
	1023: "proto序列化出错",
	//Ping
	1031: "proto反序列化出错",
	1032: "proto序列化出错",
	//Heart
	1041: "proto反序列化出错",
	1042: "proto序列化出错",
	//GetSigninTime
	1051: "proto反序列化出错",
	1052: "proto序列化出错",
	//ReConnect
	1061: "proto反序列化出错",
	1062: "proto序列化出错",
	1063: "获取玩家信息出错",
	1064: "获取玩家信息empty",
	1065: "获取信号令出错",
	1066: "信号令为空",
	1067: "信号令不相等",
	//UploadPlayerInfo
	1071: "proto反序列化出错",
	1072: "proto序列化出错",

	//GetInitGame
	2001: "proto反序列化出错",
	2002: "取城市信息出错",
	2003: "取建筑信息出错",
	2004: "取英雄信息出错",
	2005: "取用户信息出错",
	2006: "取军队信息出错",
	2007: "取signInfo信息出错",

	//C2S_GetMapCitys
	2021: "proto反序列化出错",
	2022: "取城池信息出错",
	2023: "获取资源出错",
	2024: "获取城池信息出错",
	//C2S_RefreshMapCitys
	2060: "proto反序列化出错",
	2061: "刷新城池信息出错",
	2062: "取城池信息出错",
	2063: "获取资源出错",
	2064: "获取城池皮肤出错",
	2065: "获取自己的城池信息出错",
	2066: "get armys in one map error",
	2067: "GetMapCityVO error",

	//C2S_AttackCity
	2081: "proto反序列化出错",
	2082: "取城池信息出错",
	2083: "城池信息为空",
	2084: "取军队信息出错",
	2085: "军队为空",
	2086: "该军队在行军中",
	2087: "保存军队信息出错",
	2088: "此军队有英雄已挂",
	2089: "cityFoundErr",
	2090: "cityNotFound",
	2091: "path not found",

	//C2S_AttackCityResult  not need
	2101: "proto反序列化出错",
	2102: "取军队信息出错",
	2103: "armyVO is nil",
	2104: "armyVO don't has target",
	2105: "armyVO time lack",
	2106: "getAtkHeroHeroArr error",
	2107: "getMapCityVOError",
	2108: "mapcity not found",
	2109: "targetPlayerId is 0",
	2110: "cityVO get ERROR",
	2111: "city not found",
	2112: "get defenceArmyError",
	2113: "defenceArmy is nil",
	2114: "getDefenceHerosVO error",
	2115: "save herovo error",
	2116: "defence heroVO is nil",

	//C2S_TimerRefreshMapCitys
	2120: "proto反序列化出错",
	2121: "取城池信息出错",
	2122: "获取资源出错",

	//C2S_GetAllBattleResult
	2141: "proto反序列化出错",

	//C2S_GetAllBattleResult
	2161: "proto反序列化出错",
	2162: "battleResultVO not exist",

	//C2S_Sign
	3001: "proto反序列化出错",
	3002: "拉取签到信息失败",
	3003: "今天已签到",
	3004: "签到失败",
	3005: "签到类型未定义",
	3006: "获取礼物失败",

	//C2S_GetSignInfo
	3021: "proto反序列化出错",
	3022: "拉取签到信息失败",

	//3021: "proto反序列化出错",
	//3022: "拉取签到信息失败",

	//C2S_BuyHero
	4011: "proto反序列化出错",
	4012: "GetHeroError",

	//C2S_EnQueue
	4031: "proto反序列化出错",
	4032: "get1 army error",
	4033: "get1 hero error",
	4034: "get2 army error",
	4035: "get2 hero error",
	4036: "get3 army error",
	4037: "get3 hero error",
	4038: "保存军队出错",
	4039: "保存英雄出错",

	//C2S_SetDefArmy
	4051: "proto反序列化出错",
	4052: "获取军队信息出错",
	4053: "不存在此军队",
	4054: "获取城市信息出错",
	4055: "不存在此城市",
	4056: "保存城市信息出错",

	4071: "proto反序列化出错",
	4072: "pos invalid",
	4073: "getArmy error",
	4074: "getHero error",
	4075: "hero save error",
	4076: "army save error",

	7011: "proto反序列化出错",
	7012: "get armyvo error",
	7013: "get pvelevel vo error",
	7014: "army is empty",
	7015: "cfg is nil",
	7016: "level can't be bigger than maxlevel + 1",
	7017: "get hero error",
	7018: "you must start the first level first",

	//TODO
	40111: "proto反序列化出错",
}

func GetErrorCode(errCodeMsg []byte) int32 {
	int3 := int32(errCodeMsg[0]) << 24
	int2 := int32(errCodeMsg[1]) << 16
	int1 := int32(errCodeMsg[2]) << 8
	int0 := int32(errCodeMsg[3])
	return int3 | int2 | int1 | int0
}
