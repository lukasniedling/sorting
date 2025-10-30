package bubblesort

// BubbleUp geht einmal durch das Array arr und tauscht benachbarte Elemente,
// wenn das linke größer ist als das rechte.
// Gibt true zurück, wenn mindestens ein Tausch stattgefunden hat.
func BubbleUp(arr []int) bool {
    swapped := false
    n := len(arr)
    for i := 0; i < n-1; i++ {
        if arr[i] > arr[i+1] {
            arr[i], arr[i+1] = arr[i+1], arr[i]
            swapped = true
        }
    }
    return swapped
}


// BubbleSort sortiert die übergebene Liste mittels des Bubble-Sort-Algorithmus.
func BubbleSort(arr []int) {
    for {
        hasSwapped := BubbleUp(arr)

        if !hasSwapped {
            break
        }
    }
}
