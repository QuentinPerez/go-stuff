package goselect

import "sync"

type GoSelect struct {
	ID   int
	Data interface{}
}

type Shudown struct{}

func event(id int, c chan<- GoSelect) {
	i := 0

	for {
		c <- GoSelect{
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

func Start(nbOfSalves int) {
	channels := make([]chan GoSelect, nbOfSalves)
	aggregate := make(chan GoSelect)
	// just for checking if the goroutines are closed
	wait := sync.WaitGroup{}

	for i := range channels {
		channels[i] = make(chan GoSelect)
		wait.Add(1)
		go event(i, channels[i])
	}
	for _, ch := range channels {
		wait.Add(1)
		go func(agg chan<- GoSelect, ch <-chan GoSelect) {
			for {
				select {
				case msg, ok := <-ch:
					if ok {
						agg <- msg
					} else {
						agg <- GoSelect{Data: Shudown{}}
						wait.Done()
						return
					}
				}
			}
		}(aggregate, ch)
	}
	for nb := nbOfSalves; nb > 0; {
		select {
		case msg := <-aggregate:
			if _, ok := msg.Data.(Shudown); ok {
				// fmt.Printf("Done\n")
				wait.Done()
				nb--
			} else {
				// fmt.Printf("%v <- %v\n", msg.ID, msg.Data.(int))
			}
		}
	}
	wait.Wait()
}
