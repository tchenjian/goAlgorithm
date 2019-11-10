package algorithm

func InsertSort(numArray []int) {
	for i := 1; i < len(numArray); i++ {
		insertNum := numArray[i]
		j := i - 1
		for ; j >= 0 && insertNum < numArray[j]; j-- {
			numArray[j+1] = numArray[j]
		}
		numArray[j+1] = insertNum
	}
}

func ShellSort(numArray []int) {
	d := len(numArray)
	for ; d > 1; {
		d = d / 2
		for x := 0; x < d; x++ {
			for i := x + d; i < len(numArray); i = i + d {
				tmp := numArray[i]
				j := 0
				for j = i - d; j >= 0 && numArray[j] > tmp; j = j - d {
					numArray[j+d] = numArray[j]
				}
				numArray[j+d] = tmp
			}
		}
	}
}
