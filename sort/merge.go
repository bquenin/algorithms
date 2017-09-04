package sort

func MergeSort(a []int) {
	tmp := make([]int, len(a))
	mergeSort(a, tmp, 0, len(a)-1)
}

func mergeSort(a, tmp []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(a, tmp, left, mid)
	mergeSort(a, tmp, mid+1, right)
	merge(a, tmp, left, mid, right)
}

func merge(a, tmp []int, left, mid, right int) {
	for k := left; k <= right; k++ {
		tmp[k] = a[k]
	}
	for i, j, k := left, mid+1, left; k <= right; k++ {
		if i > mid {
			a[k] = tmp[j]
			j++
		} else if j > right {
			a[k] = tmp[i]
			i++
		} else if tmp[i] > tmp[j] {
			a[k] = tmp[j]
			j++
		} else {
			a[k] = tmp[i]
			i++
		}
	}
}
