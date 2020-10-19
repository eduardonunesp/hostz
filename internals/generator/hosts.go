package generator

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"

	"github.com/eduardonunesp/hostz/internals/model"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type HostsGenerator interface {
	BuildHostsFromProfileName(name string) (string, error)
}

type hostsGenerator struct{}

func NewHostsGenerator() HostsGenerator {
	return hostsGenerator{}
}

func (hg hostsGenerator) BuildHostsFromProfileName(name string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", errors.Wrap(err, "fatal error on obtain home dir")
	}

	homeDirConfig := fmt.Sprintf("%s/%s", usr.HomeDir, model.ProfilesPath)
	bs, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", homeDirConfig, name))

	if err != nil {
		return "", errors.Wrap(err, "failed to get profile file")
	}

	var profile model.Profile
	err = yaml.Unmarshal(bs, &profile)

	if err != nil {
		return "", errors.Wrap(err, "failed to unmarshall profile")
	}

	output := fmt.Sprintf("## %s\n", name)
	for ip, hosts := range profile.HostMap {
		output += fmt.Sprintf("%s %s\n", ip, strings.Join(hosts, " "))
	}

	return output, nil
}
