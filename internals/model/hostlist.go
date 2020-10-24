package model

type Host struct {
	IP    string
	Name  string
	Alias string
}

type HostList []Host
