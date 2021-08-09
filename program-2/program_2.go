package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	key   int
	Value int
}

func countSlice(sliceOfNumber []int) map[int]int {

	MapOfItemAppearCount := map[int]int{}

	for _, number := range sliceOfNumber {

		//checkeying, is number exist in map?
		currentValue, isExist := MapOfItemAppearCount[number]

		if isExist {
			//next value is add 1
			nextValue := currentValue + 1

			//update data
			MapOfItemAppearCount[number] = nextValue

		} else {
			//push new data equal 1
			MapOfItemAppearCount[number] = 1
		}

	}

	return MapOfItemAppearCount

}

func main() {

	// initializing value
	sliceOfNumber := []int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6}

	// get map data and count
	MapOfItemAppearCount := countSlice(sliceOfNumber)

	var slicePair []Pair
	for key, value := range MapOfItemAppearCount {
		slicePair = append(slicePair, Pair{key, value})
	}

	//sort slice
	sort.Slice(slicePair, func(i, j int) bool {
		return slicePair[i].Value > slicePair[j].Value
	})

	fmt.Println("number --> total")

	result := map[int]int{}
	for _, pair := range slicePair {
		fmt.Printf("%v --> %v\n", pair.key, pair.Value)
		result[pair.key] = pair.Value

	}

}
