package main

//
// go version go1.5.2 darwin/amd64
// 4 cores
// 32 Go RAM
//
// go run main.go  0.49s user 0.07s system 206% cpu 0.270 total

// TODO use go bench

// no limit
const nbOfSalves = 10000

type shudown struct{}

type goSelect struct {
	ID   int
	Data interface{}
}

func event(id int, c chan<- goSelect) {
	i := 0

	for {
		c <- goSelect{
			ID:   id,
			Data: i,
		}
		i++
		if i == 10 {
			close(c)
			return
		}
	}
}

func main() {
	channels := make([]chan goSelect, nbOfSalves)
	aggregate := make(chan goSelect)

	for i := range channels {
		channels[i] = make(chan goSelect)
		go event(i, channels[i])
	}
	for _, ch := range channels {
		go func(agg chan<- goSelect, ch <-chan goSelect) {
			for {
				select {
				case msg, ok := <-ch:
					if ok {
						agg <- msg
					} else {
						agg <- goSelect{Data: shudown{}}
						return
					}
				}
			}
		}(aggregate, ch)
	}
	for nb := nbOfSalves; nb > 0; {
		select {
		case msg := <-aggregate:
			if _, ok := msg.Data.(shudown); ok {
				// fmt.Printf("Done\n")
				nb--
			} else {
				// fmt.Printf("%v <- %v\n", msg.ID, msg.Data.(int))
			}
		}
	}
}
