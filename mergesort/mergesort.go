package mergesort

// Merge kombiniert zwei sortierte Listen zu einer sortierten Liste.
func Merge(a1 []int, a2 []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			result = append(result, a1[i])
			i++
		} else {
			result = append(result, a2[j])
			j++
		}
	}
	// Füge verbleibende Elemente aus einer der beiden Listen hinzu
	result = append(result, a1[i:]...)
	result = append(result, a2[j:]...)
	return result
}

// MergeSort sortiert die übergebene Liste mittels des Merge-Sort-Algorithmus.
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}
