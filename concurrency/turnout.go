package concurrency

func Turnout(Quit <-chan int, InA, InB, OutA, OutB chan int) {
	var data int
	for {
		select {
		case data = <-InA:
		case data = <-InB:
		case <-Quit:
			close(InA)
			close(InB)
			Fanout(InA, OutA, OutB)
			Fanout(InB, OutA, OutB)
			return
		}
	}
}
