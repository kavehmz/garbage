package main

import "fmt"

type Item struct {
	quantity int
}

type Location struct {
}

func loadItem(location Location) Item {
	return Item{}
}
func isDamaged(location Location) bool {
	return false
}

func checkItemsQuantity(items map[Location]Item) {
	for loc, item := range items {
		id := loadItem(loc)
		if getQuantity(id) != items[loc].quantity {
			refillRequest(id, items)
		}
	}
}

func checkDamagedItems(items map[Location]Item) {
	for loc, item := range items {
		id := loadItem(loc)
		if isDamaged(id) {
			damageAlert(location)
		}
	}
}

func check(check)

func main() {

	var items []Items = []items{}

	for l := range []Location{} {
		checkItemsQuantity(items, l)
		checkDamagedItems(items, l)
	}

	var quantity = func(item) {
		if car.doorID%2 == 0 {
			fmt.Println(car.ID, "Oops, all even id door need return!")
		}
	}
	for l := range []Location{} {
		checkItemsQuantity(items, l)
		checkDamagedItems(items, l)
	}
}
