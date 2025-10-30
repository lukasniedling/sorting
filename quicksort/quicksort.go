package mergesort

// Partition teilt das Array arr in zwei Teile, basierend auf dem ersten Element
// (dem Pivot). Alle Elemente kleiner oder gleich dem Pivot werden nach links,
// alle größeren nach rechts einsortiert. Gibt die neue Position des Pivotelements zurück.
func Partition(arr []int) int {
	pivot := arr[0]
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[0], arr[i-1] = arr[i-1], arr[0]
	return i - 1
}

// QuickSort sortiert die übergebene Liste mittels des Quick-Sort-Algorithmus.
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivotIndex := Partition(arr)
	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}

