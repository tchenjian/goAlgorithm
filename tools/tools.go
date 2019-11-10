package tools

func RemoveOneFromArray(numArray []int,index int)  []int{
	tmpArray:=numArray[:index]
	for _,i:=range numArray[index+1:]{
		tmpArray=append(tmpArray,i)
	}
	return  tmpArray
}
