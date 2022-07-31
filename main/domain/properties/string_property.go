package properties

type StringProperty struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
}

func (s *StringProperty) GetType() Type {
	return s.Type
}

func (s *StringProperty) GetName() string {
	return s.Name
}

func NewStringProperty(name string) (*StringProperty, error) {
	return &StringProperty{Name: name, Type: String}, nil
}
