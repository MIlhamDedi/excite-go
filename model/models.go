package model

// Event Struct (Model)
type Event struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Venue       string        `json:"venue"`
	EO          *Organization `json:"eo"`
}

// Organization Struct (Model)
type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
