package main

import (
	"testing"
)

func Test_addNumbers(t *testing.T) {
	type args struct {
		first  float64
		second float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"correctly add two integers", args{first: 98, second: 2}, 100},
		{"correctly add two floating points", args{first: 98.45, second: 2.46}, 100.91},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addNumbers(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("addNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divideNumbers(t *testing.T) {
	type args struct {
		first  float64
		second float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"correctly divide two integers", args{first: 98, second: 2}, 49},
		{"correctly divide two floating points", args{first: 98.45, second: 2.46}, 40.020325203252035},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideNumbers(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("divideNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplyNumbers(t *testing.T) {
	type args struct {
		first  float64
		second float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"correctly multiply two integers", args{first: 98, second: 2}, 196},
		{"correctly multiply two floating points", args{first: 98.45, second: 2.46}, 242.187},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplyNumbers(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("multiplyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtractNumbers(t *testing.T) {
	type args struct {
		first  float64
		second float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"correctly subtract two integers", args{first: 98, second: 2}, 96},
		{"correctly subtract two floating points", args{first: 98.45, second: 2.46}, 95.99000000000001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractNumbers(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("subtractNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
