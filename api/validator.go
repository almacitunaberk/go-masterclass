package api

import (
	"github.com/almacitunaberk/go_masterclass/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// Check if the currency is supported or not
		return util.IsSupportedCurrency(currency)
	}
	return false

}