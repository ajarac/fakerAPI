package domain

type BooleanProperty struct {
	AbstractProperty
	Rate float32
}

func (b *BooleanProperty) GetType() Type {
	return Boolean
}

func NewBooleanProperty(rate float32) *BooleanProperty {
	return &BooleanProperty{Rate: rate}
}
