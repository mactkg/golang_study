cd $(dirname $0) && pwd

echo "gopl.io"
go run main.go http://gopl.io

echo ""
echo "bad.gopl.io"
go run main.go http://bad.gopl.io
