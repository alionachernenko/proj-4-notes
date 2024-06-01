package main

import (
	"fmt"
	"sort"
)

type Order struct {
	ID int
}

func main() {
	orders := []Order{
		{
			1,
		},
		{
			55,
		},
		{
			20,
		},
		{
			40,
		},
		{
			40,
		},
		{
			2,
		},
	}

	uniqueOrders := getUniqueOrders(orders)
	sortOrdersById(uniqueOrders)

	fmt.Print(uniqueOrders)
}

func getUniqueOrders(orders []Order) []Order {
	var uniqueOrders []Order

	for _, order := range orders {
		var duplicate bool

		for _, uniqueOrder := range uniqueOrders {
			if order.ID == uniqueOrder.ID {
				duplicate = true
				break
			}
		}

		if !duplicate {
			uniqueOrders = append(uniqueOrders, order)
		}
	}

	return uniqueOrders
}

func sortOrdersById(uniqueOrders []Order) {
	sort.Slice(uniqueOrders, func(i, j int) bool { return uniqueOrders[i].ID > uniqueOrders[j].ID })
}
