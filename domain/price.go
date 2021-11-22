package domain

import "time"

type Price struct {
	Date  time.Time
	Price float64
}

type Prices []interface{}

func (prices Prices) remove(index int) []interface{} {
	copy(prices[index:], prices[index+1:])
	prices[len(prices)-1] = Price{}
	return prices[:len(prices)-1]
}
