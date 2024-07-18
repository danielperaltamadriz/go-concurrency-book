package chapter1

func deadlock() {

	deadlockChan := make(chan int)
	<-deadlockChan

	deadlockChan <- 1
}
