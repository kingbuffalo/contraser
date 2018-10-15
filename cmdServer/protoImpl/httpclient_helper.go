package protoImpl

import (
	"buffalo/king/common/enumErrCode"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kingbuffalo/seelog"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	php_user_url_prefix string = "https://sdkapi.heywoodsminiprogram.com/user/"
	salt                string = "SpPKOMbtVOqh8oLD"
	login_check_url     string = "https://sdkapi.heywoodsminiprogram.com/user/mini_session_check"
	appid               string = "wx4155e712c9244d5e"
)

func loginCheckRetGetPid(jsonM map[string]interface{}) (string, int) {
	errCode := jsonM["errcode"]
	var errCodeInt int
	k := reflect.TypeOf(errCode)
	switch k.Kind() {
	case reflect.String:
		var err error
		errCodeInt, err = strconv.Atoi(errCode.(string))
		if err != nil {
			errCodeInt = 1
		}
		break
	case reflect.Float64:
		errCodeInt = int(errCode.(float64))
		break
	}
	if errCodeInt == enumErrCode.ERR_CODE_SUC {
		data := jsonM["data"].(map[string]interface{})
		return data["openid"].(string), 0
	}
	return "", errCodeInt

}
func loginCheck(sessionid string) (map[string]interface{}, error) {
	ct := int(time.Now().Unix())
	cts := strconv.Itoa(ct)
	signArr := []string{
		"appid=" + appid,
		"sessionid=" + sessionid,
		"sign_salt=" + salt,
		"timestamp=" + cts,
		"version=" + "0.1.0",
	}
	signStr := strings.Join(signArr, "&")
	md5sum := md5.Sum([]byte(signStr))
	//md5sumStr := string(md5sum[:])
	md5sumStr := fmt.Sprintf("%x", md5sum)
	resp, err := http.PostForm(login_check_url,
		url.Values{
			"appid":     {appid},
			"sessionid": {sessionid},
			"timestamp": {cts},
			"version":   {"0.1.0"},
			"signature": {md5sumStr},
		})
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			seelog.Debug("respBodyCloseErr", err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	seelog.Trace("body", string(body))
	var jsonM map[string]interface{}
	err = json.Unmarshal(body, &jsonM)
	if err != nil {
		return nil, err
	}
	return jsonM, nil
}
