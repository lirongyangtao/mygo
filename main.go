package main

import (
	"awesomeProject5/base"
	"fmt"
)

func main() {

	//rax := redis.NewRax()
	//rax.Add("ANNIBALE", 1)
	//rax.Add("ANNIBALExxxxx", 2)
	//rax.Add("BNNIBALE", 3)
	//rax.Add("AGO", 4)
	//rax.Add("ANNIBALI", 5)
	//rax.Add("AGO1", 6)
	////fmt.Println(rax.Find("AGO1"))
	////rax.Add("ANN", 6)
	////fmt.Println(rax.Find("AGO"))
	//fmt.Println(rax.Find("BNNIBALE"))
	//fmt.Println(rax.Find("ANNIBALExxxxx"))
	//fmt.Println(rax.Find("ANNIBALI"))
	//fmt.Println(rax.Find("AGO"))

	arr := &base.SortInt{99, 1, 2, 3, 7, 4, 5}
	base.HeapSort(arr)
	fmt.Println(arr)
}
