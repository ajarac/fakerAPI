package mapper

import (
	"fakerAPI/main/domain/properties"
	"time"
)

type JsonSchema struct {
	Name       string `json:"name"`
	Properties []JsonProperty
}

type JsonProperty struct {
	Name       string          `json:"name"`
	Type       properties.Type `json:"type"`
	Min        int             `json:"min,omitempty"`
	Max        int             `json:"max,omitempty"`
	Rate       float32         `json:"rate,omitempty"`
	From       time.Time       `json:"from,omitempty"`
	To         time.Time       `json:"to,omitempty"`
	Element    *JsonProperty   `json:"element,omitempty"`
	RangeSize  [2]int          `json:"rangeSize,omitempty"`
	Properties []JsonProperty  `json:"properties,omitempty" json:"properties"`
}
