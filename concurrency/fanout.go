package concurrency

func Fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In {
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
}
