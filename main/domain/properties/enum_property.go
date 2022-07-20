package properties

type EnumProperty struct {
	AbstractProperty
	Enums []string
}

func NewEnumProperty(name string, enums []string) (*EnumProperty, error) {
	return &EnumProperty{
		AbstractProperty: newAbstractProperty(name, Enum),
		Enums:            enums,
	}, nil
}
