## reflectSelect

`reflect.Select`'s Benchmark

## Benchmark

```console
// go test -bench='.' -benchmem -cover -cpu=1,2,3,4
// testing: warning: no tests to run
// PASS
// BenchmarkStart5000  	       1	154165258858 ns/op	24753351224 B/op	229255376 allocs/op
// BenchmarkStart5000-2	       1	112448337910 ns/op	24704789600 B/op	228790409 allocs/op
// BenchmarkStart5000-3	       1	108479248940 ns/op	24721059424 B/op	229021220 allocs/op
// BenchmarkStart5000-4	       1	107836330297 ns/op	24782490544 B/op	229545096 allocs/op
// coverage: 95.0% of statements
// ok  	github.com/QuentinPerez/go-stuff/channel-with-reflectSelect/reflectSelect	482.942s
```
