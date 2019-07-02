buffer="n,times,type.time,nsop,byte,allock, "
for N in 1 10 100 1000 10000; do
    DATA=`cat /dev/urandom | LC_CTYPE=C tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n $N | tr '\n' ' ' | sort | uniq`
    echo "Data size: $N"

    echo "BadEcho($N):"
    go test -bench=BadEcho -benchmem -- $DATA | tail -n3 | head -n1

    echo "GoodEcho($N):"
    go test -bench=GoodEcho -benchmem -- $DATA | tail -n3 | head -n1
    
    echo ""
done
