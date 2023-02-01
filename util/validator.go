package util

// constant transaction in the USD 
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// check the supported currency in user transaction
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
  return false
}
