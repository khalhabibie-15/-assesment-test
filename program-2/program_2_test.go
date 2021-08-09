package main

import (
	"reflect"
	"testing"
)

func Test_countSlice(t *testing.T) {
	firstSlice := []int{1, 1, 1, 1, 2}
	firstMap := map[int]int{1: 4, 2: 1}

	secondSlice := []int{1, 1, 1, 1, 2, 2, 2, 3, 2}
	secondMap := map[int]int{1: 4, 2: 4, 3: 1}

	type args struct {
		sliceOfNumber []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{"test 1", args{firstSlice}, firstMap},
		{"test 2", args{secondSlice}, secondMap},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSlice(tt.args.sliceOfNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
