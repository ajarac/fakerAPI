package domain

import "time"

type DateProperty struct {
	AbstractProperty
	Between *DateBetweenProperty
}

type DateBetweenProperty struct {
	from time.Time
	to   time.Time
}

func (d *DateProperty) GetType() Type {
	return Date
}

func NewDateProperty(between *DateBetweenProperty) *DateProperty {
	return &DateProperty{
		Between: between,
	}
}
