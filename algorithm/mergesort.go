package algorithm


//go语言专用归并，利用切片处理
func MergeSort(numArray []int) {
	l := len(numArray)
	if l > 1 {
		mid := l / 2
		MergeSort(numArray[:mid])
		MergeSort(numArray[mid:])
		merge(numArray,mid,l)
	}
}

//一般语言使用的归并，只能依靠index处理关系
func NormalMergeSort(numArray []int, start int, end int) {
	if start < end {
		mid := (start + end) / 2
		NormalMergeSort(numArray, start, mid)
		NormalMergeSort(numArray, mid+1, end)
		normalMerge(numArray, start, mid, end)
	}
}

func merge(numArray []int,mid,end int){

	p1 := 0
	p2 := mid

	tmpArray := make([]int, end)

	for p := 0; p<end; p++ {
		if p1 > mid-1 {
			tmpArray[p] = numArray[p2]
			p2++
			continue
		}
		if p2 > end-1 {
			tmpArray[p] = numArray[p1]
			p1++
			continue
		}

		if numArray[p1] > numArray[p2] {
			tmpArray[p] = numArray[p2]
			p2++
		} else {
			tmpArray[p] = numArray[p1]
			p1++
		}
	}

	copy(numArray, tmpArray)
}

func normalMerge(numArray []int, start int, mid int, end int) {
	tmpArray := make([]int, end-start+1)
	p1 := start
	p2 := mid + 1
	p := 0
	for ; p1 <= mid && p2 <= end; {
		if numArray[p1] <= numArray[p2] {
			tmpArray[p] = numArray[p1]
			p++
			p1++
		} else {
			tmpArray[p] = numArray[p2]
			p++
			p2++
		}
	}

	for ; p1 <= mid; {
		tmpArray[p] = numArray[p1]
		p++
		p1++
	}

	for ; p2 <= end; {
		tmpArray[p] = numArray[p2]
		p++
		p2++
	}

	for i := 0; i < len(tmpArray); i++ {
		numArray[i+start] = tmpArray[i]
	}
}
