package designanATMMachine

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	obj := Constructor()
	banknotesCount1 := []int{0, 0, 1, 2, 1}
	obj.Deposit(banknotesCount1)
	param_2 := obj.Withdraw(600)
	fmt.Println(param_2)
}
