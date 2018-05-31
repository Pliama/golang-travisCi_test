package main

import "testing"

func TestReverseOrder(t *testing.T){
	var tests = []struct{
		number uint32
		expected string
	}{
		{81,"18"},
		{1122,"2211" }}

	for _,test:=range tests{
		answer:= reverseOrder(test.number)
		if answer != test.expected {
			t.Error("Test Failed, expected",test.expected," got :", reverseOrder(test.number))
		}
	}
}