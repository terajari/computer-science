package arrays

func sum(arr []int) int {

	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Equilibrium finds the index of an element in the given array where the sum of the elements to the left of the index is equal to the sum of the elements to the right of the index.
//
// arr: the input array of integers
// Returns the index of the equilibrium point if found, otherwise -1.
func Equilibrium(arr []int) int {

	var leftSum int
	rightSum := sum(arr)

	for i := 0; i < len(arr); i++ {
		rightSum -= arr[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += arr[i]
	}
	return -1
}
