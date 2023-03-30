$ go test -bench=. -benchmem -cpu=1
goos: darwin
goarch: arm64
pkg: github.com/mroth/blakebench
cpu: Apple M2 Pro
BenchmarkXCryptoBlake/fcaddr-payload             6282751           180.0 ns/op         0 B/op          0 allocs/op
BenchmarkXCryptoBlake/fcaddr-chksum              6738614           176.3 ns/op         0 B/op          0 allocs/op
BenchmarkMinioBlakeSIMD/fcaddr-payload           5156140           231.2 ns/op         0 B/op          0 allocs/op
BenchmarkMinioBlakeSIMD/fcaddr-chksum            4713930           255.0 ns/op         0 B/op          0 allocs/op
PASS
ok      github.com/mroth/blakebench     5.689s

$ GOAMD64=v3 go test -bench=. -benchmem -cpu=1
goos: linux
goarch: amd64
pkg: github.com/mroth/blakebench
cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
BenchmarkXCryptoBlake/fcaddr-payload         	 5997702	       205.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkXCryptoBlake/fcaddr-chksum          	 6164827	       194.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinioBlakeSIMD/fcaddr-payload       	 4708561	       254.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkMinioBlakeSIMD/fcaddr-chksum        	 4252591	       281.9 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/mroth/blakebench	5.779s
