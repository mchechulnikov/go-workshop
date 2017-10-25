package fifth

func GetSliceWithMaxSum(values []int) []int {
	startIndex, endIndex, possibleStartIndex, maxSum := 0, 0, 0, 0
	sums := make([]int, 0)

	for index, value := range values {
		if index == 0 {
			endIndex = 0
			maxSum = 0
			if values[0] < 0 {
				possibleStartIndex = 1
			} else {
				maxSum = values[0]
			}
			sums = append(sums, maxSum)
			continue
		}

		if value < 0 && sums[index - 1] + value <= 0 {
			value = 0
			possibleStartIndex = index + 1
		}
		sum := sums[index - 1] + value
		if sum >= maxSum {
			maxSum = sum
			startIndex = possibleStartIndex
			endIndex = index
		}
		sums = append(sums, sum)
	}
	return values[startIndex:endIndex+1]
}
