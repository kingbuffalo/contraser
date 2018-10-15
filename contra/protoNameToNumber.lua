local function string_split(str,sep)
	local t = {}
	local p = string.format("([^%s]+)",sep)
	string.gsub(str,p,function(c)t[#t+1]=c end)
	return t
end

local f = assert(io.open("king.proto", "r"))
local allstr = f:read("*a")
f:close()

local t = string_split(allstr,"\n")

local numToName = {}
local packageName = nil
--local packageId = nil

local num = nil
for i,v in ipairs(t) do
	if string.find(v,"//--") then
		local numstr = string.sub(v,5)
		num = tonumber(numstr)
		print(num)
	end
	if num ~= nil then
		if string.find(v,"package") then
			local endIdx = string.find(v,";")
			local pkname = string.sub(v,9,endIdx-1)
			packageName = pkname
			--packageId = num
			num = nil
		end
		if string.find(v,"message") then
			local endIdx = string.find(v,"{")
			local messageName = string.sub(v,9,endIdx-1)
			numToName[num] = messageName
			num = nil
		end
	end
end

local wstrt = {}
local server_wstrt = {}

local cstrt = {}
local cstrt_r = {}

local cstrtEx = {}
local cstrt_rEx = {}

--local packagePrefix = packageId << 9
for k,v in pairs(numToName) do
	--local dig = packagePrefix | k
	local dig = k
	dig = dig & 0xffff
	local value = packageName.."."..v
	wstrt[#wstrt+1] = {
		key = dig,
		value = value
	}
	server_wstrt[dig] = value
	server_wstrt[value] = dig

	if string.find(value,"S2C_") ~= nil then
		cstrt[#cstrt+1] = dig  .. ":".. value
	else
		cstrt_r[#cstrt_r+1] =  "\""..value .. "\":"..dig
	end

	local cvalue = "protobuf.roots.laya." .. value
	cstrtEx[#cstrtEx+1] = dig  .. ":".. cvalue
	cstrt_rEx[#cstrt_rEx+1] =  "\""..cvalue .. "\":"..dig
end

local serpent = require("serpent")
print(serpent.dump(wstrt))

local clientKeyValueStr = table.concat(cstrt,",\n")
local clientValueKeyStr = table.concat(cstrt_r,",\n")

local clientKeyValueStrEx = table.concat(cstrtEx,",\n")
local clientValueKeyStrEx = table.concat(cstrt_rEx,",\n")

local clientstrt = {
	"import{",
	packageName,
	[[} from "./bundle";
const {ccclass, property} = cc._decorator;

@ccclass
export class Procnetpack {

	public static S2C_MethonKeyGen= {
	]],
	clientKeyValueStr,
	[[};
]],
	[[ public static C2S_MethonKeyGen = { ]],
	clientValueKeyStr,
	[[};
]],


[[}]]
}

local clientstrtEx = {
	"module laya{\n",
	[[export class LayaProcnetpack{
	public static S2C_MethonKeyGen= {
	]],
	clientKeyValueStrEx,
	[[};
]],
	[[ public static C2S_MethonKeyGen = { ]],
	clientValueKeyStrEx,
	[[};
]],


[[}
}]]
}


local wcstr = table.concat(clientstrt,"")
local wsstr =serpent.dump(server_wstrt)

local protoNumMapNamef = io.open("protoNumMapName.lua","w")
protoNumMapNamef:write(wsstr)
protoNumMapNamef:close()

local clientf = io.open("Procnetpack.ts","w")
clientf:write(wcstr)
clientf:close()


local wcstrEx = table.concat(clientstrtEx,"")
local clientfEx = io.open("LayaProcnetpack.ts","w")
clientfEx:write(wcstrEx)
clientfEx:close()
