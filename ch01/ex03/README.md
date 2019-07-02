## time
- `$ cat INPUT_SOMETHING | go run main.go`
- or
- `$ go run main.go foo bar a b c 1234 1024 42`

## testing.Benchmark
- `$ ./run.bash`

### result
```
Data size: 1
BadEcho(1):
  500000              3942 ns/op             176 B/op          4 allocs/op
GoodEcho(1):
  300000              4608 ns/op              80 B/op          2 allocs/op

Data size: 10
BadEcho(10):
  200000             13627 ns/op            1216 B/op         13 allocs/op
GoodEcho(10):
  100000             19042 ns/op             176 B/op          2 allocs/op

Data size: 100
BadEcho(100):
   10000            133890 ns/op           63284 B/op        103 allocs/op
GoodEcho(100):
   10000            115177 ns/op            1169 B/op          2 allocs/op

Data size: 1000
BadEcho(1000):
    1000           1807382 ns/op         5890290 B/op       1007 allocs/op
GoodEcho(1000):
    2000            940536 ns/op           12372 B/op          2 allocs/op

Data size: 10000
BadEcho(10000):
      10         128048876 ns/op        582719102 B/op     10007 allocs/op
GoodEcho(10000):
     200           7816892 ns/op          229651 B/op          4 allocs/op
```
