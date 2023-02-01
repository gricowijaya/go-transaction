package api

import (
  "github.com/go-playground/validator/v10"
  "github.com/gricowijaya/go-transaction/util"
)

// this valid currency is checking the validation for the currency in the API 
// of transaction, this api will override the validator function and return 
// a boolean value for returning the ok true
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
  if currency, ok := fieldLevel.Field().Interface().(string); ok {
    return util.IsSupportedCurrency(currency) // true
  }
  return false
}
