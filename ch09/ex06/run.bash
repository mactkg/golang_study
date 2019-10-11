echo "1..."
GOMAXPROCS=1 go test -bench=.
echo "2..."
GOMAXPROCS=2 go test -bench=.
echo "4..."
GOMAXPROCS=4 go test -bench=.
