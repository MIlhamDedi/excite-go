package model

import "io/ioutil"

// Model Interface
type Model interface {
	saveCard() error
}

// Organization Struct (Model)
type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description []byte `json:"description"`
}

// Event Struct (Model)
type Event struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description []byte        `json:"description"`
	Venue       string        `json:"venue"`
	EO          *Organization `json:"eo"`
}

// SaveCard save event details as card on file text
func (e *Event) SaveCard() error {
	filename := e.Name + ".txt"
	return ioutil.WriteFile(filename, e.Description, 0600)
}

// ReadCard read saved card
func ReadCard(name string) (*Event, error) {
	filename := name + ".txt"
	description, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Event{Name: name, Description: description}, nil
}
