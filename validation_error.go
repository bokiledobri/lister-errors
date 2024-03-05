package lister_errors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct{
    Error string `json:"error"`
    Field string `json:"field,omitempty"`
}

func ValidateFields(err error) []ValidationError{
    m := make([]ValidationError, 0)
	if errors, ok := err.(validator.ValidationErrors); ok {
		for _, v := range errors {
			switch v.ActualTag() {
			case "required":
				m = append(m, createMessage("is required", v.Field()))
			case "email":
				m = append(m, createMessage("is not a valid email", v.Field()))
			case "gt":
				switch v.Type().String() {
				case "string":
					m = append(m, createMessage(fmt.Sprintf("should have more than %s characters", v.Param()), v.Field()))
				default:
					m = append(m, createMessage(fmt.Sprintf("should be greater than %s", v.Param()), v.Field()))
				}
			default:
				m = append(m, createMessage(v.Field(), v.ActualTag()))
			}
		}
	}
    return m
}
func createMessage(e string, f string) ValidationError {
	m := new(ValidationError)
	m.Field = f
	m.Error = e
	return *m
}

