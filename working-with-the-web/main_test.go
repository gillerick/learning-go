package main

import (
	"reflect"
	"testing"
)

func Test_checkError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_toursFromJson(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want []Tour
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toursFromJson(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toursFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
