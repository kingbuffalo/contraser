local json = require("json")
local t = require("protoNumMapName")
local wt = {}
for k,v in pairs(t) do
	wt[#wt+1] = {Key=tostring(k),Value=tostring(v)}
end
local allStr = json.encode(wt)
local f = io.open("proto.json","w")
f:write(allStr)
f:close()
