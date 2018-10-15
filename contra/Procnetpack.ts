import{king} from "./bundle";
const {ccclass, property} = cc._decorator;

@ccclass
export class Procnetpack {

	public static S2C_MethonKeyGen= {
	704:king.S2C_GetPVELevelId,
706:king.S2C_ContinueFight,
201:king.S2C_GetInitGame,
203:king.S2C_AttackNpc,
205:king.S2C_AttackNpcReleaseSkill,
402:king.S2C_BuyHero,
404:king.S2C_EnQueue,
406:king.S2C_SetDefArmy,
408:king.S2C_EnQueueSingle,
410:king.S2C_HeroRecruit,
412:king.S2C_FreeHeroRecruit,
101:king.S2C_Login,
102:king.S2C_ErrMsg,
104:king.S2C_Ping,
106:king.S2C_Heart,
301:king.S2C_Sign,
303:king.S2C_GetSignInfo,
305:king.S2C_GetRedPoint,
307:king.S2C_GetRankInfo,
108:king.S2C_ReConnect,
702:king.S2C_FightPVE};
 public static C2S_MethonKeyGen = { "king.C2S_ContinueFight":705,
"king.C2S_GetInitGame":200,
"king.C2S_AttackNpc":202,
"king.C2S_AttackNpcReleaseSkill":204,
"king.C2S_BuyHero":401,
"king.C2S_EnQueue":403,
"king.C2S_SetDefArmy":405,
"king.C2S_EnQueueSingle":407,
"king.C2S_HeroRecruit":409,
"king.C2S_FreeHeroRecruit":411,
"king.C2S_Login":100,
"king.C2S_Ping":103,
"king.C2S_Heart":105,
"king.C2S_ReConnect":107,
"king.C2S_Sign":300,
"king.C2S_GetSignInfo":302,
"king.C2S_GetRedPoint":304,
"king.C2S_GetRankInfo":306,
"king.C2S_FightPVE":701,
"king.C2S_GetPVELevelId":703};
}