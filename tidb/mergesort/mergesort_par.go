package main

import (
	"runtime"
	"sync"
)


//Function to help schedule the go subroutines. Buffer used to avoid over-provision and waiting of resources on each other
func MergeSortPar(src []int64) {
	procs := runtime.NumCPU()-1 //declare and intialize variable 
	bufChannels := make(chan struct{}, procs)
	defer close(bufChannels) //defer runs after the return of this function. close is a signal to the receiver that the buffer is not longer being send data to
	mergesort(src, bufChannels)
}


func mergesort(src []int64, bufChannels chan struct{}){
	arr_len := len(src)

	if arr_len == 1 { //base case
		return
	}

	mid := arr_len/2
	left := make([]int64, mid)
	right := make([]int64, arr_len - mid)

	copy(left, src[:mid])
	copy(right, src[mid:])

	wg := sync.WaitGroup{} //Wait group, acts as a semaphore

	select { //select runs one of the cases if possible else runs the default case
		case bufChannels <- struct{}{}: //Go does not allow untyped channels hence struct{}{} is one of the ways that can be exploited
			wg.Add(1) //add 1 to semaphore
			go func() { //works as a fork() in C, where this function is run on another thread
				mergesort(left, bufChannels)
				<- bufChannels
				wg.Done() //substract 1, or remove from semaphore value
			}()
		default:
			mergesort(left, bufChannels)
	}

	mergesort(right, bufChannels) //mergesort the right part in this thread

	wg.Wait() //equivalent to join() in C++, wait for threads to finish
	merge(src, left, right)  //used the same function in sequential merge, called here
}

