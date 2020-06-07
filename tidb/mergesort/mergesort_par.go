package main

import (
	"runtime"
	"sync"
	//"sort" 
)


//Function to help schedule the go subroutines. Buffer used to avoid over-provision and waiting of resources on each other
func MergeSortPar(src []int64) {
	procs := runtime.NumCPU()-1 //declare and intialize variable 
	bufChannels := make(chan struct{}, 2*procs) //give enough work to system to that its not deprived of work
	defer close(bufChannels) //defer runs after the return of this function. close is a signal to the receiver that the buffer is not longer being send data to
	temp := make([]int64, len(src))
	mergesort(src, temp, bufChannels)
}


func mergesort(src, temp []int64, bufChannels chan struct{}){
	arr_len := len(src)

	if(arr_len <= 1000) { //Base case use sequential algorithm 
		MergeSort(src)
		//sort.Slice(src, func(i, j int) bool { return src[i] <= src[j] }) //inbuilt quicksort, considerably more memory efficient due to in-place sorting with similar runtime complexity. Parameters: array and sorting function
		return
	}

	mid := arr_len/2
	
	left, ltemp := src[:mid], temp[:mid] //less memory allocation than earlier, use of the same temp array by appropriate partitioning of it
	right, rtemp := src[mid:], temp[mid:] 

	wg := sync.WaitGroup{} //Wait group, acts as a semaphore

	/*
	In cases, where the work is small semaphore logic is an extra overhead, we are just better of spawning go subroutines as many as needed as there are free slots in the trace. 
	However, as work increases its necessary to limit the number of subroutines to avoid large number of subroutines as compared to the resources available.
	*/
	select { //select runs one of the cases if possible else runs the default case
		case bufChannels <- struct{}{}: //Go does not allow untyped channels hence struct{}{} is one of the ways that can be exploited
			wg.Add(1) //add 1 to semaphore
			go func() { //works as a fork() in C, where this function is run on another thread
				mergesort(left, ltemp, bufChannels)
				<- bufChannels
				wg.Done() //substract 1, or remove from semaphore value
			}()
		default:
			mergesort(left, ltemp, bufChannels)
	}

	mergesort(right, rtemp, bufChannels) //mergesort the right part in this thread

	wg.Wait() //equivalent to join() in C++, wait for threads to finish
	merge_opt(src, temp, left, right)  //used the same function in sequential merge, called here
}

func merge_opt(src, temp, left, right []int64) {
	var lindex, rindex, counter int //default value is 0 or null in Go

	for lindex < len(left) && rindex < len(right) { //while loop in Go
		if left[lindex] <= right[rindex] {
			temp[counter] = left[lindex]
			lindex++
		} else {
			temp[counter] = right[rindex]
			rindex++
		}
		counter++
	}

	for lindex < len(left) {
		temp[counter] = left[lindex]
		lindex++
		counter++
	}

	for rindex < len(right) {
		temp[counter] = right[rindex]
		rindex++
		counter++
	}

	copy(src, temp)
}

