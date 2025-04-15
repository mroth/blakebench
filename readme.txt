$ go test -bench=. -cpu=1
goos: darwin
goarch: arm64
pkg: github.com/mroth/blakebench
cpu: Apple M2 Pro
BenchmarkXCryptoBlake/fcaddr-payload             6447487               177.2 ns/op
BenchmarkXCryptoBlake/fcaddr-chksum              6785844               175.1 ns/op
BenchmarkMinioBlakeSIMD/fcaddr-payload           5803861               205.6 ns/op
BenchmarkMinioBlakeSIMD/fcaddr-chksum            5179191               226.6 ns/op
PASS
ok      github.com/mroth/blakebench     5.674s

$ GOAMD64=v3 go test -bench=. -benchmem -cpu=1
goos: linux
goarch: amd64
pkg: github.com/mroth/blakebench
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkXCryptoBlake/fcaddr-payload         	 6237770	           193.0 ns/op
BenchmarkXCryptoBlake/fcaddr-chksum          	 6213296	           192.9 ns/op
BenchmarkMinioBlakeSIMD/fcaddr-payload       	 4714474	           254.8 ns/op
BenchmarkMinioBlakeSIMD/fcaddr-chksum        	 4257506	           283.1 ns/op
PASS
ok  	github.com/mroth/blakebench	    5.747s
