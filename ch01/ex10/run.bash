cd $(dirname $0) && pwd

go build -o fetchall main.go
./fetchall https://www.reddit.com/r/all/
./fetchall https://www.reddit.com/r/all/
