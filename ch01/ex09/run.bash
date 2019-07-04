cd $(dirname $0) && pwd

echo "https://httpbin.org/status/200"
go run main.go https://httpbin.org/status/200

echo ""
echo "https://httpbin.org/status/404"
go run main.go https://httpbin.org/status/404

echo ""
echo "https://httpbin.org/status/500"
go run main.go https://httpbin.org/status/500
