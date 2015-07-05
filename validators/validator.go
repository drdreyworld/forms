package validators

type Validator interface {
	Create(meta ValidatorMeta) Validator
	IsValid(value interface{}) bool
	GetError() string
}

type Validators []Validator

type ValidatorMeta struct {
	Type    string
	Options interface{}
}

type ValidatorsMeta []ValidatorMeta
