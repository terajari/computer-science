package arrays

import "sort"

// TripletSumZero finds all triplets in the given array that sum to zero.
//
// Parameters:
// - arr: An array of integers.
//
// Returns:
// - A 2D array of integers representing the triplets that sum to zero.
func TripletSumZero(arr []int) [][]int {

	if len(arr) > 0 {
		var triplets [][]int

		sort.Ints(arr)

		for i := 0; i < len(arr); i++ {
			if i > 0 && arr[i] == arr[i-1] {
				continue
			}

			left := i + 1
			right := len(arr) - 1

			for left < right {
				sum := arr[i] + arr[left] + arr[right]
				if sum == 0 {
					triplet := []int{arr[i], arr[left], arr[right]}
					triplets = append(triplets, triplet)
					for left < right && arr[left] == arr[left+1] {
						left++
					}
					for left < right && arr[right] == arr[right-1] {
						right--
					}

					left++
					right--
				} else if sum < 0 {
					left++
				} else {
					right--
				}
			}
		}
		return triplets
	}

	return [][]int{}
}
