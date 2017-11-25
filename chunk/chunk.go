package chunk

// --- Directions
// Given an array and chunk size, divide the array into many subarrays
// where each subarray is of length size
// --- Examples
// chunk([1, 2, 3, 4], 2) --> [[ 1, 2], [3, 4]]
// chunk([1, 2, 3, 4, 5], 2) --> [[ 1, 2], [3, 4], [5]]
// chunk([1, 2, 3, 4, 5, 6, 7, 8], 3) --> [[ 1, 2, 3], [4, 5, 6], [7, 8]]
// chunk([1, 2, 3, 4, 5], 4) --> [[ 1, 2, 3, 4], [5]]
// chunk([1, 2, 3, 4, 5], 10) --> [[ 1, 2, 3, 4, 5]]

func chunk(arr []int, size int) [][]int {

	result := [][]int{}

	arrLen := len(arr)
	numChunks := 0
	if arrLen%size == 0 {
		numChunks = arrLen / size
	} else {
		numChunks = (arrLen / size) + 1
	}

	//fmt.Printf("NumChunks: %v\n", numChunks)
	index := 0
	// loop through chunks
	for i := 0; i < numChunks; i++ {
		slice := []int{}

		// iterate on original array
		for j := 0; j < size && index < len(arr); j++ {
			slice = append(slice, arr[index])
			index++
		}

		result = append(result, slice)
	}

	return result
}
