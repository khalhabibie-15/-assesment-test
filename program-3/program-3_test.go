package main

import "testing"

func Test_mergerString(t *testing.T) {
	type args struct {
		string1 string
		string2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"saya", "kamu"}, "skaaymau"},
		{"test 2", args{"kosong", "ada"}, "kaodsaong"},
		{"test 2", args{"ada", "kosong"}, "akdoasong"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergerString(tt.args.string1, tt.args.string2); got != tt.want {
				t.Errorf("mergerString() = %v, want %v", got, tt.want)
			}
		})
	}
}
