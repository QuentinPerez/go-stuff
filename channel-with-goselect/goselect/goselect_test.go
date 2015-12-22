package goselect

import "testing"

// go test -bench='.' -benchmem -cover -cpu=1,2,3,4
// testing: warning: no tests to run
// PASS
// BenchmarkStart10000  	      20	  93723669 ns/op	 2809683 B/op	  112617 allocs/op
// BenchmarkStart10000-2	      20	  86137894 ns/op	 2880550 B/op	  113718 allocs/op
// BenchmarkStart10000-3	      20	  82514322 ns/op	 2897494 B/op	  113977 allocs/op
// BenchmarkStart10000-4	      20	  79759312 ns/op	 2866640 B/op	  113501 allocs/op
// coverage: 100.0% of statements
// ok  	github.com/QuentinPerez/go-stuff/channel-with-goselect/goselect	7.655s

func BenchmarkStart10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Start(10000)
	}
}
