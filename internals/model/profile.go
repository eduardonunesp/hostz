package model

type Profile struct {
	Name   string  `yaml:"name"`
	Groups []Group `yaml:"groups"`
}
