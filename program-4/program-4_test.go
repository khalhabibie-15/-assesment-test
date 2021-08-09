package main

import (
	"testing"
)

func Test_stackLetter(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"JACK", "DANIEL"}, "DAJACKNIEL"},
		{"test 2", args{"ABACABA", "ABACABA"}, "AABABACABACABA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stackLetter(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("stackLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 1", args{"MAKAN"}, true},
		{"test 2", args{"MAkAN"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUpper(tt.args.s); got != tt.want {
				t.Errorf("IsUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
