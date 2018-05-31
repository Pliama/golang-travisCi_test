package main

import "testing"

func TestThreeInN(t *testing.T){
	var tests = []struct{
		number uint32
		expected bool
	}{
		{18,true},
		{19,true },
		{111, true}}

		for _,test:=range tests{
			answer:= searchThreeInN2(test.number)
			if answer != test.expected {
				t.Error("Test Failed, expected",test.expected," got :", searchThreeInN2(test.number))
			}
		}
}
