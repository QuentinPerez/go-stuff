package goselect

import "testing"

// go test -bench='.' -benchmem -cover -cpu=1,2,3,4
// ?   	github.com/QuentinPerez/go-stuff/channel-with-goselect	[no test files]
// testing: warning: no tests to run
// PASS
// BenchmarkStart5000  	      30	  43595542 ns/op	 1419512 B/op	   56541 allocs/op
// BenchmarkStart5000-2	      50	  39558259 ns/op	 1415775 B/op	   56476 allocs/op
// BenchmarkStart5000-3	      50	  34770516 ns/op	 1385317 B/op	   56000 allocs/op
// BenchmarkStart5000-4	      50	  36051506 ns/op	 1394619 B/op	   56146 allocs/op
// coverage: 100.0% of statements
// ok  	github.com/QuentinPerez/go-stuff/channel-with-goselect/goselect	7.552s

func BenchmarkStart5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Start(5000)
	}
}
