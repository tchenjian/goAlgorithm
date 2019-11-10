package algorithm

//判断是否2的n次幂
func IsPowerOf2(num int) bool {
	return num&(num-1)==0
}

//寻找奇数个数的数
func LackNum(numArray []int)  int{
	tmp:=numArray[0]
	for i:=1;i<len(numArray);i++{
		tmp^=numArray[i]
	}
	return tmp
}

//删除一个元素，使得剩余乘积最大
func FindRemovedIndex(numArray []int) int  {
	//负数个数
	negativeCount:=0
	for i:=0;i<len(numArray);i++{
		if numArray[i]<0{
			negativeCount++
		}
	}

	tmpIndex:=0
	//负数个数为奇数
	if negativeCount&1 ==1{
		for i:=1;i<len(numArray);i++{
			if numArray[i]<0{
				if numArray[tmpIndex]>=0||numArray[i]>numArray[tmpIndex]{
					tmpIndex=i
				}
			}
		}
		return tmpIndex
	}
	//负数个数为偶数，且全为负数
	if len(numArray)==negativeCount{
		for i:=0;i<len(numArray);i++{
			if numArray[i]<numArray[tmpIndex]{
				tmpIndex = i
			}
		}
		return tmpIndex
	}
	//负数个数为偶数
	for i:=0;i<len(numArray);i++{
		if numArray[i]>=0{
			if numArray[tmpIndex]<0||numArray[i]<numArray[tmpIndex]{
				tmpIndex = i
			}
		}
	}
	return tmpIndex
}

