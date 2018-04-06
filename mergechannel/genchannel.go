package mergechannel

// genIntChan creates int channels
func genIntChan(values ...int) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, value := range values {
			c <- value
			//time.Sleep(1 * time.Millisecond)
		}
	}()
	return c
}
