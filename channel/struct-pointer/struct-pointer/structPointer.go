package structPointer

type huge struct {
	_ [4192]int64
	_ [1]uint8
	_ [2]uint16
	// not aligned
}

func consumerStruct(recv <-chan huge, send chan<- huge) {
	for {
		data, ok := <-recv
		if !ok {
			return
		}
		send <- data
	}
}

func StartWithStruct(nbOfTransfer int) {
	input := make(chan huge)
	output := make(chan huge)

	go consumerStruct(input, output)
	for ; nbOfTransfer > 0; nbOfTransfer-- {
		input <- huge{}
		<-output
	}
	close(input)
	close(output)
}

func consumerPointer(recv <-chan *huge, send chan<- *huge) {
	for {
		data, ok := <-recv
		if !ok {
			return
		}
		send <- data
	}
}

func StartWithPointer(nbOfTransfer int) {
	input := make(chan *huge)
	output := make(chan *huge)

	go consumerPointer(input, output)
	for ; nbOfTransfer > 0; nbOfTransfer-- {
		input <- &huge{}
		<-output
	}
	close(input)
	close(output)
}
