local M = {}
local utilsFunc = require("buffaloutils/utilsFunc")
local json = require("json")
local enum = require("enum")

local skynet = require("skynet")
local SERVER_NAME = "buffalo/webclient"
--local URL_PREFIX = "https://pay.heywoods.cn/api/"
--local URL_PREFIX = "http://pay_dev.eenpay.cn/api/"
local URL_PREFIX
local APP_ID = enum.APP_ID

local function getUrlPrefix()
	if URL_PREFIX == nil then
		local phpurl_test = skynet.getenv("phpurl_test") or 0
		phpurl_test = tonumber(phpurl_test)
		if phpurl_test == 1 then
			URL_PREFIX = "http://pay_dev.eenpay.cn/api/"
		else
			URL_PREFIX = "https://pay.heywoods.cn/api/"
		end
	end
	return URL_PREFIX
end

local function httppost(player_id,url,postT,bLog)
	local no_reply = false
	local wscAddr = skynet.queryservice(SERVER_NAME)
	local isOk,str = skynet.call(wscAddr,"lua","request",url,nil,postT,no_reply)
	local ret = str
	if type(str) == "string" then
		ret = json.decode(str)
	end
	if bLog then
		utilsFunc.chargePrint(player_id,"url,postT:",url,postT,"isOk,retstr",isOk,str)
		if ret.code ~= 10000 then
			utilsFunc.chargeErrPrint(player_id,"url,postT:",url,postT,"isOk,retstr",isOk,str)
		end
	end
	return isOk,ret
end

local function sign(t)
	local allStrT = {}
	for k,v in pairs(t) do
		allStrT[#allStrT+1] = k .. "=".. v
	end
	table.sort(allStrT)
	return table.concat(allStrT,"&")
end

--------------------------------------------------------------------charge begin
function M.clearBalance(player_id,session_id,amount)
	local urlprefix = getUrlPrefix()
	local url = urlprefix.."orders/midas/create"
	local ct = os.time()
	local randStr = utilsFunc.genRandString(16)
	local getT= {
		appid = APP_ID,
		nonce_str = randStr,
		timestamp = ct,
		secret = enum.MINI_GAME_SECRET,
		amount=amount,
		session_id= session_id,
		platform="android",
		goods_info ="clearBalance",
	}

	local signStr = sign(getT)
	local md5 = require("md5")
	local md5sign = md5.sumhexa(signStr)

	local postT = {
		session_id = session_id,
		amount = amount,
		nonce_str = randStr,
		platform="android",
		goods_info ="clearBalance",

		timestamp = ct,
		appid = APP_ID,
		sign = md5sign,
	}

	local _,creatT= httppost(player_id,url,postT,true)
	local transaction_id = creatT.data.transaction_id

	url = urlprefix.."orders/midas/pay"
	ct = os.time()
	randStr = utilsFunc.genRandString(16)
	postT= {
		appid = APP_ID,
		nonce_str = randStr,
		session_id = session_id,
		timestamp = ct,
		secret = enum.MINI_GAME_SECRET,
		transaction_id=tostring(transaction_id),
	}

	signStr = sign( postT)
	md5 = require("md5")
	md5sign = md5.sumhexa(signStr)

	postT.sign = md5sign
	postT.secret = nil

	local isOk,ret= httppost(player_id,url,postT,true)
	if isOk == false then return nil end
	return ret
end


function M.createCharge(player_id,session_id,goods_info,amount,platform)
	local urlprefix = getUrlPrefix()
	local url = urlprefix.."orders/midas/create"
	local ct = os.time()
	local randStr = utilsFunc.genRandString(16)
	local getT= {
		appid = APP_ID,
		nonce_str = randStr,
		timestamp = ct,
		secret = enum.MINI_GAME_SECRET,
		goods_info = goods_info,
		amount=amount,
		session_id= session_id,
		platform=platform,
	}

	local signStr = sign(getT)
	local md5 = require("md5")
	local md5sign = md5.sumhexa(signStr)

	local postT = {
		session_id = session_id,
		goods_info = goods_info,
		amount = amount,
		nonce_str = randStr,
		platform = platform,

		timestamp = ct,
		appid = APP_ID,
		sign = md5sign,
	}

	local isOk,ret = httppost(player_id,url,postT,true)
	if isOk == true then
		return ret
	end
	return nil
end

function M.consumeCharge(player_id,session_id,transaction_id)
	local urlprefix = getUrlPrefix()
	local url = urlprefix.."orders/midas/pay"
	local ct = os.time()
	local randStr = utilsFunc.genRandString(16)
	local postT= {
		appid = APP_ID,
		nonce_str = randStr,
		session_id = session_id,
		timestamp = ct,
		secret = enum.MINI_GAME_SECRET,
		transaction_id=transaction_id,
	}

	local signStr = sign( postT)
	local md5 = require("md5")
	local md5sign = md5.sumhexa(signStr)

	postT.sign = md5sign
	postT.secret = nil

	local isOk,ret = httppost(player_id,url,postT,true)
	if isOk == false then return nil end
	return ret
end

function M.querySum(player_id,session_id,transaction_id)
	local urlprefix = getUrlPrefix()
	local url = urlprefix.."orders/midas/balance"
	local ct = os.time()
	local randStr = utilsFunc.genRandString(16)
	local postT= {
		appid = APP_ID,
		nonce_str = randStr,
		timestamp = ct,
		secret = enum.MINI_GAME_SECRET,
		session_id = session_id,
		platform = "android",
	}

	local signStr = sign( postT)
	local md5 = require("md5")
	local md5sign = md5.sumhexa(signStr)

	postT.sign = md5sign
	postT.secret = nil

	local _,ret =httppost(player_id,url, postT,true)
	return ret
end

function M.queryCharge(player_id,transaction_id)
	local urlprefix = getUrlPrefix()
	local url = urlprefix.."result/query"
	local ct = os.time()
	local randStr = utilsFunc.genRandString(16)
	local postT= {
		appid = APP_ID,
		nonce_str = randStr,
		timestamp = ct,
		secret = enum.MINI_GAME_SECRET,
		transaction_id=transaction_id,
	}

	local signStr = sign( postT)
	local md5 = require("md5")
	local md5sign = md5.sumhexa(signStr)

	postT.sign = md5sign
	postT.secret = nil

	local _,ret=httppost(player_id,url, postT,true)
	return ret
end
------------------------------------------------------------------charge end
--
local PHP_USR_URL_PREFIX = "https://sdkapi.heywoodsminiprogram.com/user/"
local SALT = "SpPKOMbtVOqh8oLD"
function M.loginCheck(sessionid)
	local url = PHP_USR_URL_PREFIX .. "mini_session_check"
	local ct = os.time()
	local postT = {
		appid = APP_ID,
		sessionid = sessionid,
		timestamp = ct,
		version = "0.1.0",
		sign_salt = SALT,
	}

	local signStr = sign( postT)
	local md5 = require("md5")
	local md5sign = md5.sumhexa(signStr)
	utilsFunc.debugPrint("loginCheck md5",signStr,md5sign)
	postT.signature = md5sign
	postT.sign_salt = nil

	local _,ret=httppost(0,url, postT,false)
	return ret
end

function M.shareCheck(token)
	local url = PHP_USR_URL_PREFIX .. "mini_share_check"
	local ct = os.time()
	local postT = {
		appid = APP_ID,
		token_type = 1,
		share_token = token,
		timestamp = ct,
		version = "0.1.0",
		sign_salt = SALT,
	}

	local signStr = sign( postT)
	local md5 = require("md5")
	local md5sign = md5.sumhexa(signStr)
	postT.signature = md5sign
	postT.sign_salt = nil

	local _,ret=httppost(0,url, postT,false)
	return ret
end

return M
