package main

import (
	"fmt"
	"chj/algorithm/algorithm"
	"chj/algorithm/tools"
)

func main() {
	testShellSort()
}

func testShellSort()  {
	numArray:=[]int{5,3,9,12,6,1,7,2,4,11,8,10}
	algorithm.ShellSort(numArray)
	fmt.Println(numArray)
}

func testInsertSort()  {
	numArray:=[]int{12,1,3,46,5,0,-3,12,35,16}
	algorithm.InsertSort(numArray)
	fmt.Println(numArray)
}

func testMergeSort()  {
	numArray:=[]int{5,8,6,3,9,1,7}
	algorithm.MergeSort(numArray)
	//algorithm.MergeSort(numArray,0,len(numArray)-1)
	fmt.Println(numArray)
}

func testFindRemovedIndex() {
	numArray := []int{-4, 3, -5, -7, -21, 9, -1, 5, 6}
	fmt.Println(algorithm.FindRemovedIndex(numArray))

	fmt.Println(numArray)
	numArray = tools.RemoveOneFromArray(numArray,algorithm.FindRemovedIndex(numArray))
	fmt.Println(numArray)

	numArray = []int{4, 3, 5, -7, -21, 9, -1, -5, 6, 0}
	fmt.Println(algorithm.FindRemovedIndex(numArray))

	fmt.Println(numArray)
	numArray = tools.RemoveOneFromArray(numArray,algorithm.FindRemovedIndex(numArray))
	fmt.Println(numArray)

	numArray = []int{-4, -3, -5, -7, -21, 0, -1, -8}
	fmt.Println(algorithm.FindRemovedIndex(numArray))

	fmt.Println(numArray)
	numArray = tools.RemoveOneFromArray(numArray,algorithm.FindRemovedIndex(numArray))
	fmt.Println(numArray)

}

func testRadixSort() {
	strArray := []string{"qdd", "abc", "qwe", "h", "ae", "cis", "oye"}
	fmt.Println(algorithm.RadixSort(strArray, 3))
}

func testCountSort() {
	array := []int{5, 8, 6, 3, 9, 2, 7, 1, 5, 0, 6, 9, 2, 1, 7}
	lackNum := algorithm.CountSort(array)
	fmt.Println(lackNum)
}

func testLackNum() {
	array := []int{5, 8, 6, 3, 9, 2, 7, 1, 5, 0, 6, 9, 2, 1, 7}
	lackNum := algorithm.LackNum(array)
	fmt.Println(lackNum)
}

func testPowerOf2() {
	for num := 0; num < 1024; num++ {
		if algorithm.IsPowerOf2(num) {
			fmt.Println(num, true)
		}
	}
}

func testBubbleSort() {
	array := []int{5, 8, 6, 3, 9, 2, 1, 7}
	algorithm.BubbleSort(array)
	fmt.Println(array)
}
