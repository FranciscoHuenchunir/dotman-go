package model

type DotmanPath struct {
	Name        string `json:"name"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Linked      bool   `json:"linked"`
}
