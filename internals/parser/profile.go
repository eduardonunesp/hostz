package parser

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/eduardonunesp/hostz/internals/model"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ProfileParser interface {
	GetProfileNames() ([]string, error)
}

type profileParser struct{}

func NewProfileParser() ProfileParser {
	return profileParser{}
}

func (pp profileParser) GetProfileNames() ([]string, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "fatal error on obtain home dir")
	}

	homeDirConfig := fmt.Sprintf("%s/%s", usr.HomeDir, model.ProfilesPath)
	files, err := ioutil.ReadDir(homeDirConfig)

	if err != nil {
		return nil, errors.Wrap(err, "cannot read dir")
	}

	var profileNames []string

	for _, file := range files {
		bs, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", homeDirConfig, file.Name()))

		if err != nil {
			return nil, errors.Wrap(err, "failed to get profile file")
		}

		var profile model.Profile
		err = yaml.Unmarshal(bs, &profile)

		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshall profile")
		}

		profileNames = append(profileNames, profile.Name)
	}

	return profileNames, nil
}
