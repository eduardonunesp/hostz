package model

const ProfilesPath = ".hostz"

type Profile struct {
	Name    string `yaml:"name"`
	HostMap `yaml:"hostmap"`
}
