package main

import (
	"fmt"
	"github.com/whimax/wildberries"
	"log"
	"time"
)

func main() {
	var wbClient wildberries.Interface
	wbClient = wildberries.New("")
	lastMonth := GetLastMonthDate()
	// выполним запрос на получение всех заказов за последний месяц

	orders, err := wbClient.Orders().UntilDone(1*time.Second, 3).Report(false, lastMonth)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range orders {
		fmt.Println(v)
	}

	stocks, err := wbClient.Stocks().UntilDone(1*time.Second, 3).Report(lastMonth)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range stocks {
		fmt.Println(v)
	}

	sales, err := wbClient.Sales().UntilDone(1*time.Second, 3).Report(false, lastMonth)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range sales {
		fmt.Println(v)
	}

	incomes, err := wbClient.Incomes().UntilDone(1*time.Second, 3).Report(lastMonth)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range incomes {
		fmt.Println(v)
	}
}

func GetLastMonthDate() time.Time {
	var current time.Time
	current = time.Now().AddDate(0, -1, 0)
	return time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, current.Location())
}
