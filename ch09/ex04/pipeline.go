package pipeline

func makePipeline(length uint, done <-chan struct{}) (input chan<- int, output <-chan int) {
	firstIn := make(chan int)
	out := make(chan int)

	in := firstIn
	for i := uint(0); i < length; i++ {
		go func(in <-chan int, out chan<- int) {
			var v int
			select {
			case v = <-in:
			case <-done:
				return
			}

			select {
			case out <- v:
			case <-done:
				return
			}
		}(in, out)
		in = out
		out = make(chan int)
	}

	return firstIn, out
}
