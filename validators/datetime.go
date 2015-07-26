package validators

import (
	"regexp"
)

func init() {
	Factory.Register("datetime", new(DateTime))
}

type DateTime struct {
	Format string
	Error  string
	Valid  bool
}

func (prototype DateTime) Create(meta ValidatorMeta) Validator {
	validator := new(DateTime)
	validator.Format = "^\\d{2}\\.\\d{2}\\.\\d{4} \\d{2}:\\d{2}$"
	return validator
}

func (validator *DateTime) GetError() string {
	return validator.Error
}

func (validator *DateTime) IsValid(value interface{}) bool {
	if len(value.(string)) == 0 {
		validator.Valid = true
		validator.Error = ""
	} else {
		valid, err := regexp.MatchString(validator.Format, value.(string))

		validator.Valid = valid
		validator.Error = ""

		if err != nil {
			validator.Error = err.Error()
		} else if !valid {
			validator.Error = "Не корректный формат!"
		}
	}
	return validator.Valid
}
