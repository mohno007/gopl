package pingpong

func makePingpong(done chan struct{}) (start func(), result <-chan uint) {
	ch1 := make(chan uint)
	ch2 := make(chan uint)
	res := make(chan uint)

	go func() {
		var v uint
	loop:
		for {
			select {
			case v = <-ch1:
				ch2 <- v + 1
			case <-done:
				break loop
			}
		}
		res <- v
	}()
	go func() {
		var v uint
	loop2:
		for {
			select {
			case v = <-ch2:
				ch1 <- v + 1
			case <-done:
				break loop2
			}
		}
		res <- v
	}()

	return func() { ch1 <- 0 }, res
}
