package pkg

type ErrorType string

func (e ErrorType) Error() string {
	return string(e)
}

const (
	ErrSdkCantBeEmpty     = ErrorType("Sdk can't be empty!")
	ErrWrongTypeAssertion = ErrorType("Wrong type assertion")
)
