package solves

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	ratios := make([]float64, n)
	for i := 0; i < n; i++ {
		ratios[i] = float64(wage[i]) / float64(quality[i])
	}
	quickSort(ratios, 0, n-1)
	heap := make([]int, 0)
	sum := 0
	for i := 0; i < k; i++ {
		heap = append(heap, quality[i])
		sum += quality[i]
	}
	floatHeap := make([]float64, k)
	for j := 0; j < k; j++ {
		floatHeap[j] = float64(heap[j])
	}

	quickSort(floatHeap, 0, k-1)
	min := float64(sum) * ratios[k-1]
	for i := k; i < n; i++ {
		if quality[i] < heap[k-1] {
			sum -= heap[k-1]
			sum += quality[i]
			heap[k-1] = quality[i]
			quickSort(floatHeap, 0, k-1)
		}
		if float64(sum)*ratios[i] < min {
			min = float64(sum) * ratios[i]
		}
	}
	return min
}

func quickSort(arr []float64, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}

func partition(arr []float64, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
