package microPipeline

type Micro interface {
	Input() <-chan interface{}
	Output() chan<- interface{}
}
