package main

func main() {
	// 1) method
	//
	// ~3 950 000 msgs/s
	//
	// rw := NewRW()
	// input := rw.Input()
	// link := make(chan interface{})
	// go func() {
	// 	for d := range rw.Output(NewReader().Output(nil), NewWriter().Output(nil)) {
	// 		link <- d
	// 	}
	// }()
	// for {
	// 	input <- link
	// }

	// 2) method  1 + buffer
	//
	// ~11 116 000 msgs/s
	//
	// rw := NewRWBuffered()
	// input := rw.Input()
	// link := make(chan interface{}, 100000)
	// go func() {
	// 	for d := range rw.Output(NewReaderBuffered().Output(nil), NewWriterBuffered().Output(nil)) {
	// 		link <- d
	// 	}
	// }()
	// for {
	// 	input <- link
	// }

	// 3) method
	//
	// ~ 1 050 000 msgs/s
	//
	// rw := NewRW()
	// input := rw.Input()
	// for d := range rw.Output(NewReader().Output(nil), NewWriter().Output(nil)) {
	// 	input <- d
	// }

	// 4) method 3 + buffer
	//
	// ~ 3 359 500 msgs/s
	//
	rw := NewRWBuffered()
	input := rw.Input()
	for d := range rw.Output(NewReaderBuffered().Output(nil), NewWriterBuffered().Output(nil)) {
		input <- d
	}
}
