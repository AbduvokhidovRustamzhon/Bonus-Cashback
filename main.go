package main

func main() {

}

func bonusCounter(sales []int) int {
	const percentOfCashback = 5
	const borderOfBonus = 10_000
	const HundredPercent = 100
	sumOfCashback := 0
	for _, sale := range sales {

		if sale > borderOfBonus {
			bonus := sale - borderOfBonus
			cashback := percentOfCashback * bonus / HundredPercent
			sumOfCashback += cashback
		}

	}

	return sumOfCashback

}
