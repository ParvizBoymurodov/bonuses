package main

func main() {

}
func bonuses(sales [] int) int {
	sum := 0
	const percent = 5
	const minAmountForBonus = 10_000
	for _, sale := range sales {
		if sale > minAmountForBonus {
			result := (sale - minAmountForBonus) * percent / 100
			sum = sum + result
		}
	}
	return sum
}
