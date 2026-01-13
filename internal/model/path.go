package model

type DotmanPath struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Linked      bool   `json:"linked"`
}
