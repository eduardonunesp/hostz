package generator

import (
	"bytes"
	"fmt"
	"os"

	"github.com/eduardonunesp/hostz/internals"
	"github.com/eduardonunesp/hostz/internals/model"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type HostsGenerator interface {
	BuildHostsFromProfileName(name string) (string, error)
	BuildHostsFromProfile(profile model.Profile) string
}

type hostsGenerator struct{}

func NewHostsGenerator() HostsGenerator {
	return hostsGenerator{}
}

func (hg hostsGenerator) BuildHostsFromProfileName(name string) (string, error) {
	homeDir, err := internals.GetHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "fatal error on obtain home dir")
	}

	bs, err := os.ReadFile(fmt.Sprintf("%s/%s", homeDir, name))

	if err != nil {
		return "", errors.Wrap(err, "failed to get profile file")
	}

	var profile model.Profile
	err = yaml.Unmarshal(bs, &profile)
	if err != nil {
		return "", errors.Wrap(err, "failed to unmarshall profile")
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("## %s\n", name))
	for _, host := range profile.HostList {
		buffer.WriteString(fmt.Sprintf("%s %s %s\n", host.IP, host.Name, host.Alias))
	}
	return buffer.String(), nil
}

func (hg hostsGenerator) BuildHostsFromProfile(profile model.Profile) string {
	output := fmt.Sprintf("## %s\n", profile.Name)
	for _, host := range profile.HostList {
		output += fmt.Sprintf("%s %s %s\n", host.IP, host.Name, host.Alias)
	}

	return output
}
