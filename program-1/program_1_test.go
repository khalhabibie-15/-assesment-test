package main

import (
	"testing"
)

func Test_swapvalue(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"case 1", args{2, 3}, 3, 2},
		{"case 2", args{4, 5}, 5, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := swapvalue(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("swapvalue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("swapvalue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
