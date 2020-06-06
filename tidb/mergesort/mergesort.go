package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	var arr_len int =len(src)
	if arr_len == 1 {
		return
	}

	mid := arr_len/2

	//allocate array for left and right slices
	left := make([]int64, mid)
	right := make([]int64, arr_len-mid)

	//copy array to left and right temp array
	copy(left, src[:mid])
	copy(right, src[mid:])

	//merge sort the parts and merge the sorted halves
	MergeSort(left)
	MergeSort(right)
	merge(src, left, right)
}

//call by reference as we pass array as parameters
func merge(result, left, right []int64) {
	var lindex, rindex, counter int //default value is 0 or null in Go

	for lindex < len(left) && rindex < len(right) { //while loop in Go
		if left[lindex] <= right[rindex] {
			result[counter] = left[lindex]
			lindex++
		} else {
			result[counter] = right[rindex]
			rindex++
		}
		counter++
	}

	for lindex < len(left) {
		result[counter] = left[lindex]
		lindex++
		counter++
	}

	for rindex < len(right) {
		result[counter] = right[rindex]
		rindex++
		counter++
	}
}

//Complexity: O(nlog(n)) time and O(n) space. Standard sequential mergersort. 

