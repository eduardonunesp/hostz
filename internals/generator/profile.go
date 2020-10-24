package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/eduardonunesp/hostz/internals"
	"github.com/eduardonunesp/hostz/internals/model"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type ProfileGenerator interface {
	CreateProfileFileFromName(name string) error
	CreateProfileFromHostList(name string, hostMap model.HostList) error
}

type profileGenerator struct{}

func NewProfileGenerator() ProfileGenerator {
	return profileGenerator{}
}

func (pg profileGenerator) buildFromProfile(profile model.Profile) error {
	bs, err := yaml.Marshal(profile)
	if err != nil {
		return errors.Wrap(err, "failed to create profile")
	}

	usr, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "fatal error on obtain home dir")
	}

	homeDirConfig := fmt.Sprintf("%s/%s", usr.HomeDir, model.ProfilesPath)
	homeConfigFile := fmt.Sprintf("%s/%s", homeDirConfig, profile.Name)

	if isDir := internals.DirExists(homeDirConfig); !isDir {
		if err := os.Mkdir(homeDirConfig, 0755); err != nil {
			panic(fmt.Errorf("fatal error on create config file: %s", err))
		}
	}

	if configExists := internals.FileExists(homeConfigFile); !configExists {
		if err := ioutil.WriteFile(homeConfigFile, []byte{}, 0644); err != nil {
			panic(fmt.Errorf("unable to write config file: %s", err))
		}
	}

	if err := ioutil.WriteFile(homeConfigFile, bs, 0644); err != nil {
		return errors.Wrap(err, "fatal error write configuration")
	}

	return nil
}

func (pg profileGenerator) CreateProfileFileFromName(name string) error {
	profile := model.Profile{
		Name: name,
	}

	return pg.buildFromProfile(profile)
}

func (pg profileGenerator) CreateProfileFromHostList(name string, hostList model.HostList) error {
	profile := model.Profile{
		Name:     name,
		HostList: hostList,
	}

	return pg.buildFromProfile(profile)
}
