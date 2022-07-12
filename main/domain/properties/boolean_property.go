package properties

type BooleanProperty struct {
	AbstractProperty `bson:"metadata"`
	Rate             float32 `bson:"rate"`
}

func NewBooleanProperty(name string, rate float32) *BooleanProperty {
	return &BooleanProperty{
		AbstractProperty: newAbstractProperty(name, Boolean),
		Rate:             rate,
	}
}
