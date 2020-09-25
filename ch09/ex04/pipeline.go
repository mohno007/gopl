package pipeline

func makePipeline(length uint) (input chan<- int, output <-chan int) {
	in := make(chan int)
	out := make(chan int)

	for i := uint(0); i < length; i++ {
		go func(in <-chan int, out chan<- int) {
			v := <-in
			out <- v
		}(in, out)
		output = make(chan int)
	}

	return in, out
}
