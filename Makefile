benchpprof:
	go test -run=^$ -bench=^BenchmarkPrimeNumbers$ -benchmem -cpu=1,2,3,4 -o=/tmp/foo.test -cpuprofile=/tmp/cpuprofile.out ./test/somefunc_test.go

pprof:
	go tool pprof -http=:5000 /tmp/cpuprofile.out

benchtrace:
	go test -run=^$ -bench=^BenchmarkPrimeNumbers$ -trace=/tmp/trace.out ./test/

trace:
	go tool trace /tmp/trace.out