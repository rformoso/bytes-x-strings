# bytes-x-strings
Bytes x Strings performance benchmark.

bytes-x-strings is a simple benchmark for verifying performance between replace types in a character set.


## How to use

``` sh
go test -bench=.
```

## Result

```sh
> go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/rformoso/bytes-x-strings
BenchmarkBytesReplace-4          1000000              1651 ns/op            6144 B/op          1 allocs/op
BenchmarkStringsReplace-4        1000000              1536 ns/op           12288 B/op          2 allocs/op
BenchmarkBytesRegexp-4            500000              2401 ns/op            7548 B/op         18 allocs/op
PASS
ok      github.com/rformoso/bytes-x-strings     4.467s
```

## Donate

Good Vibrations