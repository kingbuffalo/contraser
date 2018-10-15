killall fakeclient
rm fakeclient
go build fakeclient.go
cd ../../..
./script/tools/fakeclient/fakeclient script/start_cds.json 127.0.0.1:1001 10.0.0.211:9701 &
