package internals

import (
	"fmt"
	"os"
	"os/user"

	"github.com/eduardonunesp/hostz/internals/model"
	"github.com/pkg/errors"
)

func DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}

	return false
}

func FileExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}

	return false
}

func GetHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", errors.Wrap(err, "fatal error on obtain home dir")
	}

	homeDirConfig := fmt.Sprintf("%s/%s", usr.HomeDir, model.ProfilesPath)

	return homeDirConfig, nil
}

func GetDefaultProfile() model.Profile {
	return model.Profile{
		Name: "default",
		HostList: model.HostList{
			model.Host{
				IP:    "127.0.0.1",
				Name:  "localhost",
				Alias: "loopback",
			},
			model.Host{
				IP:   "255.255.255.255",
				Name: "broadcasthost",
			},
			model.Host{
				IP:   "127.0.0.1",
				Name: "localhost.localdomain",
			},
			model.Host{
				IP:   "127.0.0.1",
				Name: "local",
			},
			model.Host{
				IP:   "::1",
				Name: "localhost",
			},
		},
	}
}
