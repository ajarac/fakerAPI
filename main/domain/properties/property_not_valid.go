package properties

import "fmt"

type PropertyNotValid struct {
	msg string
}

func NewPropertyNotValid(name string, msg string) *PropertyNotValid {
	return &PropertyNotValid{msg: fmt.Sprintf("Property %s not valid: %s", name, msg)}
}

func (e *PropertyNotValid) Error() string {
	return e.msg
}
