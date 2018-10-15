package enumErrCode

const (
	PLAYER_ID_INC_KEY string = "PLAYER_ID_INC_KEY"
	NPC_IDX_INC_KEY   string = "NPC_IDX_INC_KEY"
	PLAYER_ID_SET     string = "PLAYER_ID_SET"
	//PROTO_PREFIX_LEN           int    = 9
	WS_READ_TIMEOUT_SECOND     = 70
	BATTLE_KEEP_ALIVE_TIMER    = 60
	MAPCITY_ID_TYPE_NPC        = 1
	MAPCITY_ID_TYPE_PLAYER     = 2
	ARMY_INCR_IDX              = 6
	BATTLE_PUSH_PROTO_ID       = 1
	BATTLE_PUSH_PROTO_LEN      = 0
	BATTLE_PUSH_KEEP_ALIVE_LEN = 0
	BATTLE_PUSH_KEEP_ALIVE_ID  = 0
	//ERR_CODE_SUC               = 0
	//PHPURL_TEST     = 1
	//DATA_LEN_EXTERN = 2
)

var GOOD_STATUE_BIT_MASK uint64 = (1 << STATUE_UNABLE_BE_ATK) | (1 << STATUE_UNABLE_BE_ENEMY_SKILL) | (1 << STATUE_UNABLE_BE_ENEMY_BUFFER) | (1 << STATUE_UNABLE_BE_ENEMY_STATUE) | (1 << STATUE_UNABLE_DEAD)
var BAD_STATUE_BIT_MASK uint64 = (1 << STATUE_DIZZY) | (1 << STATUE_UNABLE_ATK) | (1 << STATUE_UNABLE_SKILL) | (1 << STATUE_UNABLE_BE_TEAMER_SKILL) | (1 << STATUE_UNABLE_BE_TEAMER_STATUE) | (1 << STATUE_UNABLE_BE_TEAMER_BUFFER)