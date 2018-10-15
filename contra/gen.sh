lua protoNameToNumber.lua
sleep 1
protoc --go_out=. *.proto
lua luat2json.lua
cp luat2json.lua ~/Desktop/windowsshare/king/
cp Procnetpack.ts ~/Desktop/windowsshare/king/
cp LayaProcnetpack.ts ~/Desktop/windowsshare/king/
cp king.proto ~/Desktop/windowsshare/king/
cd /home/cds/fun/go/src/buffalo/king/script/tools/
go build errMsgToJson.go
./errMsgToJson
mv /tmp/errCode.json ~/Desktop/windowsshare/king/
cd /home/cds/fun/go/src/buffalo/king/king/
sshpass -p XrzV9WjJ9RhJjY4w scp king.proto game@10.0.0.251:/home/game/
