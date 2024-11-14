package main

import (
	"fmt"
	datatstructures "hammadExchange/datastructures"
)

type SampleData struct {
	rollNumber int
}

type UserDefinedRedBlackTreeComparator struct{}

func (comparator UserDefinedRedBlackTreeComparator) Compare(lhs, rhs *SampleData) int {
	if lhs.rollNumber < rhs.rollNumber {
		return -1
	} else if lhs.rollNumber > rhs.rollNumber {
		return 1
	}
	return 0
}

func main() {
	comparator := UserDefinedRedBlackTreeComparator{}
	tree := datatstructures.NewRedBlackTree[SampleData, string](comparator)
	data5 := &SampleData{rollNumber: 5}
	val5 := "e"
	tree.Insert(data5, &val5)

	data9 := &SampleData{rollNumber: 9}
	val9 := "i"
	tree.Insert(data9, &val9)

	data1 := &SampleData{rollNumber: 10}
	val1 := "a"
	tree.Insert(data1, &val1)

	data2 := &SampleData{rollNumber: 20}
	val2 := "b"
	tree.Insert(data2, &val2)

	data3 := &SampleData{rollNumber: 5}
	val3 := "c"
	tree.Insert(data3, &val3)

	data4 := &SampleData{rollNumber: 4}
	val4 := "d"
	tree.Insert(data4, &val4)

	data6 := &SampleData{rollNumber: 6}
	val6 := "f"
	tree.Insert(data6, &val6)

	data7 := &SampleData{rollNumber: 7}
	val7 := "g"
	tree.Insert(data7, &val7)

	data8 := &SampleData{rollNumber: 8}
	val8 := "h"
	tree.Insert(data8, &val8)

	tree.Erase(data2)
	sortedElements := tree.GetSortedElements(true)
	for i := 0; i < len(sortedElements); i++ {
		if key, ok := sortedElements[i].Key.(*SampleData); ok {
			fmt.Println(key.rollNumber)
		}
		if val, ok := sortedElements[i].Value.(*string); ok {
			fmt.Println(*val)
		}
	}
	fmt.Println("atyab")
}
