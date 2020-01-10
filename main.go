package main

func main() {

}

func bonusCounter(sales []int) int {
	const percentOfCashback = 5
	const borderOfBonus = 10_000
	const hundredPercent = 100
	sumOfCashback := 0
	for _, sale := range sales {

		if sale > borderOfBonus {
			bonus := sale - borderOfBonus
			cashback := percentOfCashback * bonus / hundredPercent
			sumOfCashback += cashback
		}

	}

	return sumOfCashback

}
