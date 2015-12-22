package structPointer

import "testing"

// go test -bench='StartWithStruct' -v -benchmem -cpu=1,2,3,4
// testing: warning: no tests to run
// PASS
// BenchmarkStartWithStruct100000  	       2	 585674136 ns/op	     416 B/op	       3 allocs/op
// BenchmarkStartWithStruct100000-2	       2	 844651197 ns/op	     256 B/op	       3 allocs/op
// BenchmarkStartWithStruct100000-3	       2	 798072712 ns/op	     224 B/op	       2 allocs/op
// BenchmarkStartWithStruct100000-4	       2	 792149718 ns/op	     384 B/op	       2 allocs/op
// ok  	github.com/QuentinPerez/go-stuff/channel-struct-pointer/struct-pointer	9.065s

func BenchmarkStartWithStruct100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartWithStruct(100000)
	}
}

// go test -bench='StartWithPointer' -v -benchmem -cpu=1,2,3,4
// testing: warning: no tests to run
// PASS
// BenchmarkStartWithPointer100000  	       2	 673373232 ns/op	4096000480 B/op	  100003 allocs/op
// BenchmarkStartWithPointer100000-2	       2	 807885851 ns/op	4096016608 B/op	  100253 allocs/op
// BenchmarkStartWithPointer100000-3	       2	 833188863 ns/op	4096041728 B/op	  100632 allocs/op
// BenchmarkStartWithPointer100000-4	       2	 846809760 ns/op	4096048000 B/op	  100746 allocs/op
// ok  	github.com/QuentinPerez/go-stuff/channel-struct-pointer/struct-pointer	9.501s

func BenchmarkStartWithPointer100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartWithPointer(100000)
	}
}
