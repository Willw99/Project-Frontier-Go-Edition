package validator

type Validator interface {
	Validate(body string) (bool, error)
}
