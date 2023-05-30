package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPreferredPrice(t *testing.T) {
	tests := []struct {
		name                  string
		key                   string
		price                 Price
		expectedPriceValue    float64
		expectedIsMarketPrice bool
	}{
		{
			name: "it should return price when the market price is 0 even if prefer price source is Coingecko",
			key:  "address",
			price: Price{
				Address:           "address",
				Price:             1,
				Liquidity:         100000,
				LpAddress:         "lpAddress",
				MarketPrice:       0,
				PreferPriceSource: PriceSourceCoingecko,
			},
			expectedPriceValue:    1,
			expectedIsMarketPrice: false,
		},
		{
			name: "it should return market price when the prefer price source is Coingecko",
			key:  "address",
			price: Price{
				Address:           "address",
				Price:             1,
				Liquidity:         100000,
				LpAddress:         "lpAddress",
				MarketPrice:       2,
				PreferPriceSource: PriceSourceCoingecko,
			},
			expectedPriceValue:    2,
			expectedIsMarketPrice: true,
		},
		{
			name: "it should return price when the prefer price source is lyfebloc",
			key:  "address",
			price: Price{
				Address:           "address",
				Price:             1,
				Liquidity:         100000,
				LpAddress:         "lpAddress",
				MarketPrice:       2,
				PreferPriceSource: PriceSourcelyfebloc,
			},
			expectedPriceValue:    1,
			expectedIsMarketPrice: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			priceValue, isMarketPrice := test.price.GetPreferredPrice()
			assert.Equal(t, test.expectedPriceValue, priceValue)
			assert.Equal(t, test.expectedIsMarketPrice, isMarketPrice)
		})
	}
}
