package arrays

func sum(arr []int) int {

	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

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
