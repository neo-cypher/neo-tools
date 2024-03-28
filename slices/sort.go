package slices

// Partition function partitions the slice s in the range [low, high) using the element at index high as pivot.
// Elements smaller than the pivot are moved to the left side of the pivot, and larger elements to the right.
// Returns the index where the pivot is placed.
func partition[T Ordered](s []T, low, high int) int {
	for j := low; j < high; j++ {
		if s[j] < s[high] {
			s[low], s[j] = s[j], s[low]
			low++
		}
	}
	s[low], s[high] = s[high], s[low]
	return low
}

// QuickSort function sorts the slice s in the range [low, high) in ascending order using the quicksort algorithm.
// If the slice's length is less than 12, it switches to using the insertionSort function for sorting.
func quickSort[T Ordered](s []T, low, high int) {
	if low < high {
		if high-low < 12 {
			insertionSort(s, low, high)
		} else {
			p := partition(s, low, high)
			quickSort(s, low, p-1)
			quickSort(s, p+1, high)
		}
	}
}

// InsertionSort function sorts the slice s in the range [a, b) in ascending order using the insertion sort algorithm.
func insertionSort[T Ordered](s []T, a, b int) {
	for i := 1; i < b-a+1; i++ {
		j := i
		for j > 0 {
			if s[a+j] < s[a+j-1] {
				s[a+j-1], s[a+j] = s[a+j], s[a+j-1]
			}
			j--
		}
	}
}

// PartitionFunc function is similar to the partition function but uses a custom comparison function provided by 'less' to determine the ordering.
func partitionFunc[T any](s []T, low, high int, less func(T, T) bool) int {
	for j := low; j < high; j++ {
		if less(s[j], s[high]) {
			s[low], s[j] = s[j], s[low]
			low++
		}
	}
	s[low], s[high] = s[high], s[low]
	return low
}

// QuickSortFunc function is a variation of the quickSort function that is based on a custom comparison function provided by 'less'.
func quickSortFunc[T any](s []T, low, high int, less func(T, T) bool) {
	if low < high {
		if high-low < 12 {
			insertionSortFunc(s, low, high, less)
		} else {
			p := partitionFunc(s, low, high, less)
			quickSortFunc(s, low, p-1, less)
			quickSortFunc(s, p+1, high, less)
		}
	}
}

// InsertionSortFunc function is a variant of the insertionSort function that uses a custom comparison function 'less'.
func insertionSortFunc[T any](s []T, a, b int, less func(T, T) bool) {
	for i := 1; i < b-a+1; i++ {
		j := i
		for j > 0 {
			if less(s[a+j], s[a+j-1]) {
				s[a+j-1], s[a+j] = s[a+j], s[a+j-1]
			}
			j--
		}
	}
}
