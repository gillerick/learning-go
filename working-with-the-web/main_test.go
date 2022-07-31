package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func Test_checkError(t *testing.T) {
	type args struct {
		err error
	}

	_, err := ioutil.ReadFile("non-existing-file.txt")
	_, errr := ioutil.ReadFile("./fromString.txt")
	tests := []struct {
		name string
		args args
	}{
		{"return error for non existing file", args{err}},
		{"return error for non existing file", args{errr}},
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

	content, err := ioutil.ReadFile("./working-with-the-web/sampleJsonResponse.json")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name string
		args args
		want []Tour
	}{
		{"correctly parse JSON to Tour struct", args{string(content)}, []Tour{{
			Name:  "2 Days Adrift the Salton Sea",
			Price: "350",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toursFromJson(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toursFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
