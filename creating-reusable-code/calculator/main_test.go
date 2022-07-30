package main

import "testing"

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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideNumbers(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("divideNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInput(t *testing.T) {
	type args struct {
		prompt string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInput(tt.args.prompt); got != tt.want {
				t.Errorf("getInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOperation(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOperation(); got != tt.want {
				t.Errorf("getOperation() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractNumbers(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("subtractNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
