package main

import (
	"errors"
	"fmt"
)

func withdraw(balance, amount, dailyLimit int) (int, error) {
	if amount > balance {
		return balance, errors.New("insufficient funds")
	}
	if amount > dailyLimit {
		return balance, fmt.Errorf("withdrawal exceeds the daily limit of %d", dailyLimit)
	}
	balance -= amount
	return balance, nil
}

func main() {
	balance := 1000
	dailyLimit := 500
	newBalance, err := withdraw(balance, 600, dailyLimit)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Balance:", newBalance)
	}
}
