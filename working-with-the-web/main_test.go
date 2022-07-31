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
	content := `[
	{
		"tourId": "14",
		"packageId": "5",
		"packageTitle": "From Desert to Sea",
		"name": "2 Days Adrift the Salton Sea",
		"blurb": "The Salton Sea, 25% saltier than the Pacific, has been a tourist destination since the 1920s. See what attracts people to this desert oasis.",
		"description": "The Salton Sea is saltier than the Pacific, an unusual feat for inland body of water. And even though its salinity has risen over the years, due in part to lack of outflows and pollution from agricultural runoff, it has attracted a small, but dedicated population. The sea itself offers recreational opportunities including boating, camping, off-roading, hiking, use of personal watercraft, photography and bird watching. The sea has been termed a \"crown jewel of avian biodiversity,\" being a major resting stop on the Pacific Flyway, a migratory path for birds. 2 Days Adrift the Salton Sea includes two nights accommodations at the Bombay Beach Inn, boat rental at the Salton City Harbor, and a guided fishing tour.",
		"price": "350",
		"difficulty": "2",
		"length": "2",
		"graphic": "map_saltonsea.gif",
		"region": "Southern California"
	}
]`

	tests := []struct {
		name    string
		content string
		want    []Tour
	}{
		{"correctly parse JSON to Tour struct", content, []Tour{{
			Name:  "2 Days Adrift the Salton Sea",
			Price: "350",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toursFromJson(content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toursFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
