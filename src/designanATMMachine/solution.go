package designanATMMachine

/*
Runtime: 224 ms, faster than 95.90% of Go online submissions for Design an ATM Machine.
Memory Usage: 9.2 MB, less than 17.21% of Go online submissions for Design an ATM Machine.
*/
type ATM struct {
	denominations    []int
	denominationsCnt []int
}

func Constructor() ATM {
	return ATM{
		denominations:    []int{20, 50, 100, 200, 500},
		denominationsCnt: []int{0, 0, 0, 0, 0},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, v := range banknotesCount {
		this.denominationsCnt[i] += v
	}
}

func (this *ATM) Withdraw(amount int) []int {
	withdrawCnt := []int{0, 0, 0, 0, 0}

	for i := len(this.denominations) - 1; i >= 0; i-- {
		if amount == 0 {
			break
		}
		if this.denominations[i] == 0 || this.denominations[i] > amount {
			continue
		}

		if this.denominations[i]*this.denominationsCnt[i] >= amount {
			withdrawCnt[i] = amount / this.denominations[i]
			amount = amount % this.denominations[i]
		} else {
			withdrawCnt[i] = this.denominationsCnt[i]
			amount -= (this.denominations[i] * this.denominationsCnt[i])
		}
	}
	if amount != 0 {
		return []int{-1}
	}

	for i, v := range withdrawCnt {
		this.denominationsCnt[i] -= v
	}
	return withdrawCnt
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
