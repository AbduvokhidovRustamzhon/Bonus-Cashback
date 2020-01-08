package main

import "testing"

func Test_bonusCounter(t *testing.T) {

	tests := []struct {
		name string
		sales []int
		want int
	}{
		{name:"Cashback",sales:[]int{12_000, 15_000, 8_000, 9_000}, want:350},
	}
	for _, test := range tests {
		got := bonusCounter(test.sales)
			if  got != test.want {
				t.Error("Bonus Amount", test.name, "got", got, "want", test.want)
			}
		}
	}
