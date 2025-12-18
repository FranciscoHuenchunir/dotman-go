package model

type DotmanPath struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
	Target string `yaml:"target"`
	Type   string `yaml:"type"`
	Linked bool   `yaml:"linked"`
}

type DotmanPathFile struct {
	Paths []DotmanPath `yaml:"paths"`
}
