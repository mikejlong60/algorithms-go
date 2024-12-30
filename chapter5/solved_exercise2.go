package chapter5

import "math"

type DayStockPrice struct {
	day   int
	price int
}

var maxS = func(xs []DayStockPrice) DayStockPrice {
	var currMax = DayStockPrice{
		day:   math.MinInt,
		price: math.MinInt,
	}
	for _, y := range xs {
		if y.price >= currMax.price {
			currMax = y
		}
	}
	return currMax
}

var minS = func(xs []DayStockPrice) DayStockPrice {
	var currMin = DayStockPrice{
		day:   math.MaxInt,
		price: math.MaxInt,
	}
	for _, y := range xs {
		if y.price <= currMin.price {
			currMin = y
		}
	}
	return currMin
}

func MostProfit(xs []DayStockPrice) []DayStockPrice { //If no solution return the min
	if len(xs) == 0 {
		return []DayStockPrice{}
	} else if len(xs) == 2 {
		if xs[1].price > xs[0].price {
			return xs
		} else {
			return []DayStockPrice{}
		}
	} else if len(xs) == 3 {
		if xs[2].price > xs[1].price && xs[1].price > xs[0].price { //descending price
			return []DayStockPrice{xs[0], xs[2]}
		} else if xs[2].price > xs[1].price && xs[1].price < xs[0].price { //peak last element and 2nd element is smaller than first
			return []DayStockPrice{xs[1], xs[2]}
		} else if xs[2].price < xs[1].price && xs[1].price > xs[0].price { //peak middle element and 1st element is smaller than middle
			return []DayStockPrice{xs[0], xs[1]}
		} else { //no valid peak
			return []DayStockPrice{}
		}
	} else if len(xs) == 4 {
		if xs[3].price > xs[2].price &&
			xs[2].price > xs[1].price &&
			xs[1].price > xs[0].price { //descending price -- no solution
			return []DayStockPrice{xs[0], xs[3]}
		} else if xs[3].price > xs[2].price &&
			xs[3].price > xs[1].price &&
			xs[3].price > xs[0].price &&
			xs[1].price < xs[0].price &&
			xs[1].price < xs[2].price { //peak 4th element and 2nd element is min
			return []DayStockPrice{xs[1], xs[3]}
		} else if xs[3].price > xs[2].price &&
			xs[3].price > xs[1].price &&
			xs[3].price > xs[0].price &&
			xs[2].price < xs[1].price &&
			xs[2].price < xs[0].price { //peak 4th element and 3rd element is min
			return []DayStockPrice{xs[2], xs[3]}
		} else if xs[3].price > xs[2].price &&
			xs[3].price > xs[1].price &&
			xs[3].price > xs[0].price &&
			xs[0].price < xs[1].price &&
			xs[0].price < xs[2].price { //peak 4th element and 1st element is min
			return []DayStockPrice{xs[2], xs[3]}
		} else if xs[2].price > xs[3].price &&
			xs[2].price > xs[1].price &&
			xs[2].price > xs[0].price &&
			xs[1].price < xs[0].price { //peak 3rd element and 2nd element is min
			return []DayStockPrice{xs[1], xs[3]}
		} else if xs[2].price > xs[3].price &&
			xs[2].price > xs[1].price &&
			xs[2].price > xs[0].price { //peak 3rd element and 1st element is min
			return []DayStockPrice{xs[0], xs[2]}
		} else if xs[1].price > xs[2].price &&
			xs[1].price > xs[3].price &&
			xs[1].price > xs[0].price { //peak 2nd element and 1st element is min
			return []DayStockPrice{xs[0], xs[1]}
		} else { //no valid peak
			return []DayStockPrice{}
		}
	} else {
		//slice the array in half and send it off recursively
		i := len(xs) / 2
		left := xs[0:i]
		right := xs[i:]

		var a, b []DayStockPrice
		a = MostProfit(left)
		b = MostProfit(right)
		if len(a) == 0 {
			a = []DayStockPrice{minS(left)}
		}
		if len(b) == 0 {
			b = []DayStockPrice{maxS(right)}
		}
		return MostProfit(append(a, b...))
	}
}
