package fifth

func GetSliceWithMaxSum(values []int) []int {
	startIndex, endIndex, possibleStartIndex, maxSum := 0, 0, 0, 0
	sums := make([]int, 0)

	endIndex = 0
	maxSum = 0
	if values[0] < 0 {
		possibleStartIndex = 1
	} else {
		maxSum = values[0]
	}
	sums = append(sums, maxSum)

	for index, value := range values[1:] {
		if value < 0 && sums[index] + value <= 0 {
			value = 0
			possibleStartIndex = index + 2
		}
		sum := sums[index] + value
		if sum >= maxSum {
			maxSum = sum
			startIndex = possibleStartIndex
			endIndex = index + 1
		}
		sums = append(sums, sum)
	}
	return values[startIndex:endIndex+1]
}
