package functional_extensions

import "github.com/joaquimcalvogubianas/energy-pricing-parser/domain"

type PricesList []domain.Price

func (sl PricesList) First() domain.Price {
	return sl[0]
}
