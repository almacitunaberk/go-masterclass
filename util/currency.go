package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	CHF = "CHF"
	TRY = "TRY"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, CHF, TRY:
		return true
	}
	return false
}