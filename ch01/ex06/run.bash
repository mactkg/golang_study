cd $(dirname $0) && pwd

go run ./main.go > `date +%y%m%d_%H_%M_%S.gif`
