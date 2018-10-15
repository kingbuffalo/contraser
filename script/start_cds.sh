go clean -cache
killall kinggw
killall kingcmd
cd ..
rm kinggw
sleep 1
go build kinggw.go
cd cmdServer
rm kingcmd
sleep 1
go build kingcmd.go
cd ..
./cmdServer/kingcmd script/start_cds.json 127.0.0.1:21000 &
sleep 1
./kinggw script/start_cds.json 10.0.0.211:9701 &
