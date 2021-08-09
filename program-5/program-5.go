package main

import (
	"fmt"
	"log"
)

type OperationInputStruct struct {
	a int
	b int
	k int
}

type InitialInputStruct struct {
	m int
	n int
}

func GenerateSliceContainZero(sliceLength int) []int {
	slice := []int{}

	for i := 0; i < sliceLength; i++ {
		slice = append(slice, 0)
	}

	return slice
}

func GetOperationInput(operationCount int) []OperationInputStruct {
	operationInput := OperationInputStruct{}
	sliceOperationInput := []OperationInputStruct{}

	for i := 0; i < operationCount; i++ {
		fmt.Scanf("%d %d %d", &operationInput.a, &operationInput.b, &operationInput.k)
		sliceOperationInput = append(sliceOperationInput, operationInput)
	}

	return sliceOperationInput
}

func UpdateSlice(operationInputSlice []OperationInputStruct, numberSlice []int) []int {

	for _, operation := range operationInputSlice {
		for i := operation.a - 1; i < operation.b; i++ {
			log.Println(numberSlice[i])
			numberSlice[i] += operation.k
		}
	}

	return numberSlice

}

func GetHighestNumber(numberSlice []int) int {
	if len(numberSlice) != 0 {
		highestNumber := 0
		for _, number := range numberSlice {
			if number > highestNumber {
				highestNumber = number
			}
		}

		return highestNumber
	}

	return 0

}

func main() {
	//declaration variabel
	initialInput := InitialInputStruct{}

	//get format input
	fmt.Scanf("%d %d ", &initialInput.m, &initialInput.n)

	//get slice
	numberSlice := GenerateSliceContainZero(initialInput.m)

	// generate operation
	operationInputSlice := GetOperationInput(initialInput.n)

	//update slice
	numberSlice = UpdateSlice(operationInputSlice, numberSlice)

	//get highest number
	highestNumber := GetHighestNumber(numberSlice)

	fmt.Println("slice data :", numberSlice)

	fmt.Println("highest number : ", highestNumber)

}
