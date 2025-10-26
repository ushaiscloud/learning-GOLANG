package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	item, exists := menu[order]
	if !exists {
		// 404 means the item is not found in the menu
		return 404
	}
	return item.preptime
}
