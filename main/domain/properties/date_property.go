package properties

import "time"

type DateProperty struct {
	AbstractProperty `bson:"metadata"`
	From             time.Time `bson:"from"`
	To               time.Time `bson:"to"`
}

func NewDateProperty(name string, from time.Time, to time.Time) (*DateProperty, error) {
	err := validateDate(name, from, to)
	if err != nil {
		return nil, err
	}

	return &DateProperty{
		AbstractProperty: newAbstractProperty(name, Date),
		From:             from,
		To:               to,
	}, nil
}

func validateDate(name string, from time.Time, to time.Time) error {
	if from.Before(to) {
		return NewPropertyNotValid(name, "from date should be after than to date")
	}
	return nil
}
