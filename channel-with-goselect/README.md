## GoSelect

GoSelect is an alternative method at `reflect.Select`

## Benchmark

```console
go test -bench='.' -benchmem -cover -cpu=1,2,3,4 ./...
?   	github.com/QuentinPerez/go-stuff/channel-with-goselect	[no test files]
testing: warning: no tests to run
PASS
BenchmarkStart10000  	      20	  97719785 ns/op	 2825312 B/op	  112866 allocs/op
BenchmarkStart10000-2	      20	  82234619 ns/op	 2885856 B/op	  113797 allocs/op
BenchmarkStart10000-3	      20	  83770198 ns/op	 2843811 B/op	  113139 allocs/op
BenchmarkStart10000-4	      20	  82114092 ns/op	 2844860 B/op	  113156 allocs/op
coverage: 100.0% of statements
ok  	github.com/QuentinPerez/go-stuff/channel-with-goselect/goselect	8.962s
```
