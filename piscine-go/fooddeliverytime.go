package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var time food
	switch {
	case order == "burger":
		time.preptime = 15
	case order == "chips":
		time.preptime = 10
	case order == "nuggets":
		time.preptime = 12
	default:
		time.preptime = 404
	}
	return time.preptime
}
