package main

import "testing"

func Test_bonuses(t *testing.T) {
	tests:=[] struct {
		name string
		sales []int
		want int
	}{
		{name:"Have bonus",sales:[]int{12_000,8_000,15_000,9_000},want:350},
		{name:" No bonus ",sales:[]int{10_000},want:0},
	}

	for _, test:= range tests {
		got:=bonuses(test.sales)
		if got!= test.want {
			t.Error("For bonus test:",test.name, "got:",got, "want:", test.want)
		}

	}
}