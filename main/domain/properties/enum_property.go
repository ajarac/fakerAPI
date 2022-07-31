package properties

type EnumProperty struct {
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Enums []string `json:"enums"`
}

func (e *EnumProperty) GetType() Type {
	return e.Type
}

func (e *EnumProperty) GetName() string {
	return e.Name
}

func NewEnumProperty(name string, enums []string) (*EnumProperty, error) {
	return &EnumProperty{Name: name, Type: Enum, Enums: enums}, nil
}
