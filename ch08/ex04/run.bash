go run reverb.go &
sleep 1
go run netcat.go

pkill reverb
