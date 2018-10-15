package protoImpl

import (
	"buffalo/king/common/enumErrCode"
	"buffalo/king/common/gameutil"
	"crypto/md5"
	"encoding/json"
	"errors"
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
	//TODO 这个我不知道是多少
	mini_game_secret string = ""
	platform         string = "android"
)

func getUrlPrefix() string {
	URL_PREFIX := ""
	phpurl_test := enumErrCode.PHPURL_TEST
	if phpurl_test == 1 {
		URL_PREFIX = "http://pay_dev.eenpay.cn/api/"
	} else {
		URL_PREFIX = "https://pay.heywoods.cn/api/"
	}
	return URL_PREFIX
}

func getNow() string {
	now := int(time.Now().Unix())
	cts := strconv.Itoa(now)
	return cts
}
func getMd5(signStr string) string {
	md5sum := md5.Sum([]byte(signStr))
	md5sumStr := fmt.Sprintf("%x", md5sum)
	return md5sumStr
}
func shareCheck(token string) (map[string]interface{}, error) {
	url_ := php_user_url_prefix + "mini_share_check"
	cts := getNow()
	signArr := []string{
		"appid=" + appid,
		"share_token=" + token,
		"sign_salt=" + salt,
		"timestamp=" + cts,
		"token_type=" + "1",
		"version=" + "0.1.0",
	}
	signStr := strings.Join(signArr, "&")
	md5sumStr := getMd5(signStr)
	//发送的一堆数据
	resp, err := http.PostForm(url_,
		url.Values{
			"appid":       {appid},
			"share_token": {token},
			"timestamp":   {cts},
			"version":     {"0.1.0"},
			"token_type":  {"1"},
			"signature":   {md5sumStr},
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

func queryCharge(transaction_id string) (map[string]interface{}, error) {
	url_ := getUrlPrefix() + "result/query"
	cts := getNow()
	randStr := gameutil.GenRandStr(16)

	signArr := []string{
		"appid=" + appid,
		"nonce_str=" + randStr,
		"secret=" + mini_game_secret,
		"timestamp=" + cts,
		"transaction_id=" + transaction_id,
	}
	signStr := strings.Join(signArr, "&")
	md5sumStr := getMd5(signStr)
	//发送的一堆数据
	resp, err := http.PostForm(url_,
		url.Values{
			"appid":          {appid},
			"nonce_str":      {randStr},
			"timestamp":      {cts},
			"transaction_id": {transaction_id},
			"signature":      {md5sumStr},
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

func subClearBalance(transaction_id string, session_id string, cts string) (map[string]interface{}, error) {
	url_ := getUrlPrefix() + "orders/midas/pay"
	randStr := gameutil.GenRandStr(16)
	signArr := []string{
		"appid=" + appid,
		"nonce_str=" + randStr,
		"secret=" + mini_game_secret,
		"session_id=" + session_id,
		"timestamp=" + cts,
		"transaction_id=" + transaction_id,
	}
	signStr := strings.Join(signArr, "&")
	md5sumStr := getMd5(signStr)
	resp, err := http.PostForm(url_,
		url.Values{
			"appid":          {appid},
			"nonce_str":      {randStr},
			"session_id":     {session_id},
			"timestamp":      {cts},
			"signature":      {md5sumStr},
			"transaction_id": {transaction_id},
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
func clearBalance(session_id string, amount int) (map[string]interface{}, error) {
	url_ := getUrlPrefix() + "orders/midas/create"
	cts := getNow()
	randStr := gameutil.GenRandStr(16)
	signArr := []string{
		"amount=" + strconv.Itoa(amount),
		"appid=" + appid,
		"goods_info=" + "clearBalance",
		"nonce_str=" + randStr,
		"platform=" + platform,
		"secret=" + mini_game_secret,
		"session_id=" + session_id,
		"timestamp=" + cts,
	}
	signStr := strings.Join(signArr, "&")
	md5sumStr := getMd5(signStr)
	//发送的一堆数据
	resp, err := http.PostForm(url_,
		url.Values{
			"appid":      {appid},
			"amount":     {strconv.Itoa(amount)},
			"goods_info": {"clearBalance"},
			"nonce_str":  {randStr},
			"platform":   {platform},
			"secret":     {mini_game_secret},
			"session_id": {session_id},
			"timestamp":  {cts},
			"signature":  {md5sumStr},
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
	//对jsonM 进行处理, 然后再对
	errCode := jsonM["errCode"]
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
		transaction_id := data["transaction_id"].(string)
		return subClearBalance(transaction_id, session_id, cts)
	}
	errStr := "errCodeInt :" + strconv.Itoa(errCodeInt)
	return nil, errors.New(errStr)
}
func createCharge(session_id string, goods_info string, amount int, platform string) (map[string]interface{}, error) {
	url_ := getUrlPrefix() + "orders/midas/create"
	cts := getNow()
	randStr := gameutil.GenRandStr(16)
	signArr := []string{
		"amount=" + strconv.Itoa(amount),
		"appid=" + appid,
		"goods_info=" + goods_info,
		"nonce_str=" + randStr,
		"platform=" + platform,
		"secret=" + mini_game_secret,
		"session_id=" + session_id,
		"timestamp=" + cts,
	}
	signStr := strings.Join(signArr, "&")
	md5sumStr := getMd5(signStr)
	resp, err := http.PostForm(url_,
		url.Values{
			"appid":      {appid},
			"amount":     {strconv.Itoa(amount)},
			"nonce_str":  {randStr},
			"goods_info": {goods_info},
			"platform":   {platform},
			"timestamp":  {cts},
			"session_id": {session_id},
			"signature":  {md5sumStr},
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

func consumeCharge(session_id string, transaction_id string) (map[string]interface{}, error) {
	url_ := getUrlPrefix() + "orders/midas/pay"
	cts := getNow()
	randStr := gameutil.GenRandStr(16)
	signArr := []string{
		"appid=" + appid,
		"nonce_str=" + randStr,
		"secret=" + mini_game_secret,
		"session_id=" + session_id,
		"timestamp=" + cts,
		"transaction_id=" + transaction_id,
	}

	signStr := strings.Join(signArr, "&")
	md5sumStr := getMd5(signStr)
	resp, err := http.PostForm(url_,
		url.Values{
			"appid":          {appid},
			"nonce_str":      {randStr},
			"timestamp":      {cts},
			"session_id":     {session_id},
			"transaction_id": {transaction_id},
			"signature":      {md5sumStr},
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
func querySum(session_id string, transaction_id string) (map[string]interface{}, error) {
	url_ := getUrlPrefix() + "orders/midas/balance"
	cts := getNow()
	randStr := gameutil.GenRandStr(16)

	signArr := []string{
		"appid=" + appid,
		"nonce_str=" + randStr,
		"platform=" + platform,
		"secret=" + mini_game_secret,
		"session_id=" + session_id,
		"timestamp=" + cts,
	}
	signStr := strings.Join(signArr, "&")
	md5sumStr := getMd5(signStr)
	//发送的一堆数据
	resp, err := http.PostForm(url_,
		url.Values{
			"appid":      {appid},
			"nonce_str":  {randStr},
			"timestamp":  {cts},
			"platform":   {platform},
			"session_id": {session_id},
			"signature":  {md5sumStr},
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
