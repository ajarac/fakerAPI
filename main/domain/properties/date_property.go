package properties

import "time"

type DateProperty struct {
	AbstractProperty `bson:"metadata"`
	From             time.Time `bson:"from"`
	To               time.Time `bson:"to"`
}

func NewDateProperty(name string, from time.Time, to time.Time) *DateProperty {
	return &DateProperty{
		AbstractProperty: newAbstractProperty(name, Date),
		From:             from,
		To:               to,
	}
}
