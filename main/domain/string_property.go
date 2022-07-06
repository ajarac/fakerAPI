package domain

type PropertyString struct {
	AbstractProperty
}

func (s *PropertyString) GetType() Type {
	return String
}

func NewPropertyString() *PropertyString {
	return &PropertyString{}
}
