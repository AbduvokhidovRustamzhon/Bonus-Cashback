package main

func main() {

}

func bonusCounter(sales []int) int {
	const percentOfBonus = 5
	const borderOfBonus = 10_000
	const hundredPercent = 100
	sumOfBonus := 0
	for _, sale := range sales {

		if sale > borderOfBonus {
			bonusFrom := sale - borderOfBonus
			bonus := percentOfBonus * bonusFrom / hundredPercent
			sumOfBonus += bonus
		}

	}

	return sumOfBonus

}
