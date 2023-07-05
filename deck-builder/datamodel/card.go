package datamodel

// Title, Cost, Type, Misc [Strength, Points, Trash], Subtype, Inf, Set
type Card struct {
	Title    string
	Cost     *int `json:",omitempty"`
	Type     string
	Strength *int    `json:",omitempty"`
	Points   *int    `json:",omitempty"`
	Trash    *int    `json:",omitempty"`
	Subtype  *string `json:",omitempty"`
	Inf      *int    `json:",omitempty"`
	Set      string
	Money    *string  `json:",omitempty"`
	Tags     []string `json:",omitempty"`
}
