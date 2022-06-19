package mapper

type JsonSchema struct {
	Name       string `json:"name"`
	Properties []*JsonProperty
}

type JsonProperty struct {
	Name       string          `json:"name"`
	Type       string          `json:"type,omitempty"`
	Min        int             `json:"min,omitempty"`
	Max        int             `json:"max,omitempty"`
	Rate       float64         `json:"rate,omitempty"`
	Element    *JsonProperty   `json:"element,omitempty"`
	RangeSize  []int           `json:"rangeSize,omitempty"`
	Properties []*JsonProperty `json:"properties,omitempty" json:"properties"`
}
