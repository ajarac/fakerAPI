package properties

import "time"

type DateProperty struct {
	Name string    `json:"name"`
	Type Type      `json:"type"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func (d *DateProperty) GetType() Type {
	return d.Type
}

func (d *DateProperty) GetName() string {
	return d.Name
}

func NewDateProperty(name string, from time.Time, to time.Time) (*DateProperty, error) {
	err := validateDate(name, from, to)
	if err != nil {
		return nil, err
	}

	return &DateProperty{Name: name, Type: Date, From: from, To: to}, nil
}

func validateDate(name string, from time.Time, to time.Time) error {
	if from.After(to) {
		return NewPropertyNotValid(name, "from date should be after than to date")
	}
	return nil
}
