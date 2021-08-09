package main

import (
	"reflect"
	"testing"
)

func TestGenerateSliceContainZero(t *testing.T) {
	type args struct {
		sliceLength int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{3}, []int{0, 0, 0}},
		{"test 2", args{10}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSliceContainZero(tt.args.sliceLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSliceContainZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHighestNumber(t *testing.T) {
	type args struct {
		numberSlice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{100, 200, 300, 100}}, 300},
		{"test 1", args{[]int{100, 200, 300, 1000}}, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHighestNumber(tt.args.numberSlice); got != tt.want {
				t.Errorf("GetHighestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
