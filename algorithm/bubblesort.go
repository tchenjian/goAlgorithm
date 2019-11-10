package algorithm

func BubbleSort(numArray []int)  {
	tmp:=0
	for i:=0;i<len(numArray);i++{
		for j:=0;j<len(numArray)-1;j++{
			if numArray[j]>numArray[j+1]{
				tmp = numArray[j]
				numArray[j]=numArray[j+1]
				numArray[j+1]=tmp
			}
		}
	}
}