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
BenchmarkF1_OneBytesReplace-4             	 2000000	       917 ns/op	    6144 B/op	       1 allocs/op
BenchmarkF1_StringsOneReplace-4           	 1000000	      1588 ns/op	   12288 B/op	       2 allocs/op
BenchmarkF1_StringsOneReplaceConvert-4    	 1000000	      2231 ns/op	   18432 B/op	       3 allocs/op
BenchmarkF1_BytesOneRegexp-4              	  200000	      6697 ns/op	   44992 B/op	      29 allocs/op
BenchmarkF1_BytesManyRegexp-4             	    3000	    449161 ns/op	   59616 B/op	      53 allocs/op
BenchmarkF1_BytesManyReplace-4            	  200000	      7532 ns/op	   30720 B/op	       5 allocs/op
BenchmarkF1_StringsManyReplacer-4         	  100000	     23285 ns/op	   12320 B/op	       3 allocs/op
BenchmarkF1_StringsManyReplace-4          	  200000	     10917 ns/op	   61440 B/op	      10 allocs/op
BenchmarkF1_StringsManyReplaceConvert-4   	  200000	     11596 ns/op	   67584 B/op	      11 allocs/op
PASS
ok  	github.com/rformoso/bytes-x-strings	18.329s
Success: Benchmarks passed.
```

## Donate

Good Vibrations