package main

import (
	"testing"
)


var tests = []struct{
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
} {
	{"valid-data", 10, 2, 5, false},
	{"invalid-data", 10, 0, 0, true},
}


func TestDivide(t *testing.T){
	_, err := divide(10.0, 2.0)

	if err != nil {
		t.Error("Got an error we should not have")
	}
}

func TestBadDivide(t *testing.T){
	_, err := divide(10.0, 0)

	if err == nil {
		t.Error("Got an error we should not have")
	}
}

func testLoopDivision(t *testing.T){
	for _, tt := range tests{
		_, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error, but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one")
			}
		}

	}
}