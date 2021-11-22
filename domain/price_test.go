package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRemoveShouldDeleteFirstElement(t *testing.T) {
	prices := []interface{}{Price{
		Date: time.Now(),
		Price: 1670.1,
	},
	Price{
		Date: time.Now().Add(86400000000000),
		Price: 1670.1,
	}}

	var prices2 Prices = prices
	var updatedPrices = prices2.remove(0)
	assert.Condition(t, func() bool {return len(updatedPrices) == 1} , "No element has been removed")
}