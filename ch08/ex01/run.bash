go build -o clock clock.go

TZ=US/Eastern ./clock -port 8010 > /dev/null &
TZ=Asia/Tokyo ./clock -port 8020 > /dev/null &
TZ=Europe/London ./clock -port 8030 > /dev/null &
go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

pkill clock
