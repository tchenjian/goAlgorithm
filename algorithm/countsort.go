package algorithm

func CountSort(numArray []int)  []int{
	max:=numArray[0]
	min:=numArray[0]
	for i:=0;i<len(numArray);i++{
		if numArray[i]>max{
			max =numArray[i]
		}
		if numArray[i]<min{
			min = numArray[i]
		}
	}

	d:=max-min
	countArray :=make([]int,d+1,d+1)
	for i:=0;i<len(numArray);i++{
		countArray[numArray[i]-min]++
	}

	sum:=0
	for i:=0;i<len(countArray);i++{
		sum+=countArray[i]
		countArray[i] = sum
	}

	d=len(numArray)
	sortedArray :=make([]int,d,d)
	for i:=len(numArray)-1;i>=0;i--{
		sortedArray[countArray[numArray[i]-min]-1] = numArray[i]
		countArray[numArray[i]-min]--
	}
	return sortedArray
}
