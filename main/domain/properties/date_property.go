package properties

import "time"

type DateProperty struct {
	AbstractProperty
	From time.Time
	To   time.Time
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
	if from.After(to) {
		return NewPropertyNotValid(name, "from date should be after than to date")
	}
	return nil
}
