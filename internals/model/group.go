package model

type Group struct {
	Name    string `yaml:"name"`
	HostMap `yaml:"hostmap"`
}
