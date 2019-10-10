curl localhost:8080 --data-urlencode "expr=a+10" --data-urlencode "env=a=10"
curl localhost:8080 --data-urlencode "expr=a+10"
curl localhost:8080 --data-urlencode "expr=a*b+10" --data-urlencode "env=a=12,b=20"
curl localhost:8080 --data-urlencode "expr=10/2" --data-urlencode "env=a=10"
