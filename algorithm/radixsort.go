package algorithm


const ASCIIRANGE = 128

func RadixSort(StrArray []string, maxLenth int) []string {
	sortedArray := make([]string, len(StrArray))
	for k := maxLenth-1; k >= 0; k-- {
		count := make([]int, ASCIIRANGE)
		for i := 0; i < len(StrArray); i++ {
			index := getCharIndex(StrArray[i], k)
			count[index]++
		}

		for i := 1; i < len(count); i++ {
			count[i] = count[i] + count[i-1]
		}

		for i := len(StrArray) - 1; i >= 0; i-- {
			index := getCharIndex(StrArray[i], k)
			sortedIndex := count[index] - 1
			sortedArray[sortedIndex] = StrArray[i]
			count[index]--
		}

		copy(StrArray,sortedArray)
	}

	return sortedArray
}

func getCharIndex(str string, k int) int {
	if len(str) < k+1 {
		return 0
	}
	return int(str[k])
}
