package properties

type NumberProperty struct {
	AbstractProperty `bson:"metadata"`
	Min              int `bson:"min"`
	Max              int `bson:"max"`
}

func NewNumberProperty(name string, min int, max int) *NumberProperty {
	return &NumberProperty{
		AbstractProperty: newAbstractProperty(name, Number),
		Min:              min,
		Max:              max,
	}
}
