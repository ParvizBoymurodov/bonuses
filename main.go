package main

import "fmt"

func main() {
	sales:=[]int{12_000,8_000,15_000,9_000}
	fmt.Println(bonuses(sales))

}
func bonuses(sales[] int) int{
	sum:=0
	const procent=5
	const minAmountForBonus =10_000
	for _, sale:= range sales{
		if sale > minAmountForBonus {
			result:=(sale- minAmountForBonus)*procent/100
			sum=sum+result
		}
	}
	return sum
}
