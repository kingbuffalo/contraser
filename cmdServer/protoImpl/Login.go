package protoImpl

import (
	"buffalo/king/common"
	"buffalo/king/common/gameutil"
	"buffalo/king/common/startcfg"
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"github.com/kingbuffalo/seelog"
	//"buffalo/king/dbOper/playerInfo_db"
	"buffalo/king/dbOper/user_db"
	"buffalo/king/king"
)

type Login struct {
	proto_t
}

func (L *Login) DoTheJob(rpcC2SProto *common.RpcC2SProto) *common.RpcS2CRet {
	if L.bClose {
		return common.NewErrorCodeRet(107)
	}
	c2s := &king.C2S_Login{}
	err := proto.Unmarshal(rpcC2SProto.Data, c2s)
	if err != nil {
		return common.NewErrorCodeRet(1001)
	}
	seelog.Info(c2s.String())

	openId := *c2s.OpenId
	newName := base64.StdEncoding.EncodeToString([]byte(*c2s.Name))
	newUrl := *c2s.Url
	if len(openId) == 0 {
		return common.NewErrorCodeRet(1008)
	}
	if len(openId) >= 10 && !startcfg.GetBUnitTest() {
		jsonM, err := loginCheck(openId)
		if err != nil {
			return common.NewErrorCodeRet(1002)
		}
		var errcode int
		openId, errcode = loginCheckRetGetPid(jsonM)
		if errcode != 0 {
			return common.NewErrorCodeRet(1007)
		}
	}
	user, err := user_db.GetUser(*c2s.OpenId, newName, newUrl)
	if err != nil {
		return common.NewErrorCodeRet(1003)
	}
	bNeedSave := false
	if newUrl != user.Url {
		bNeedSave = true
		user.Url = newUrl
	}
	if newName != user.Name {
		bNeedSave = true
		user.Name = newName
	}
	if bNeedSave {
		if err := user.Save(true); err != nil {
			seelog.WarnWE("user.Save error=", err)
		}
	}
	token := gameutil.GenToken(user.PlayerId)
	s2c := &king.S2C_Login{
		PlayerId: proto.Int32(user.PlayerId),
		Token:    proto.String(token),
	}
	return getProtoS2CDataWithoutPush(user.PlayerId, "S2C_Login", s2c)
}
