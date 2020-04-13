package concurrency

func Funnel(InA <-chan int, OutA, OutB chan int) {
	for data := range InA {
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
}
