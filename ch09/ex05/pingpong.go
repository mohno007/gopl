package pingpong

func makePingpong(done chan struct{}) (start func(), result <-chan uint) {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	res := make(chan uint)
	count := uint(0)

	go func() {
	loop:
		for {
			select {
			case <-ch1:
				count++
				ch2 <- struct{}{}
			case <-done:
				break loop
			}
		}
		close(ch2)
		res <- count
	}()
	go func() {
	loop:
		for {
			select {
			case <-ch2:
				count++
				ch1 <- struct{}{}
			case <-done:
				break loop
			}
		}
		close(ch1)
		res <- count
	}()

	return func() { ch1 <- struct{}{} }, res
}
