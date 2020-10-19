package parser

import (
	"io/ioutil"
	"strings"

	"github.com/eduardonunesp/hostz/internals/model"
	"github.com/pkg/errors"
)

type HostsParser interface {
	ReadHostsFile(hostsPath string) ([]byte, error)
	ParseHosts(hostsFileContent []byte) model.HostMap
}

type hostsParser struct{}

func NewHostsParser() HostsParser {
	return hostsParser{}
}

func (hs hostsParser) ReadHostsFile(hostsPath string) ([]byte, error) {
	bs, err := ioutil.ReadFile(hostsPath)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read hosts file")
	}

	return bs, nil
}

func (hs hostsParser) ParseHosts(hostsFileContent []byte) model.HostMap {
	hostsMap := model.HostMap{}
	for _, line := range strings.Split(strings.Trim(string(hostsFileContent), " \t\r\n"), "\n") {
		line = strings.Replace(strings.Trim(line, " \t"), "\t", " ", -1)
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}

		pieces := strings.SplitN(line, " ", 2)
		if len(pieces) > 1 && len(pieces[0]) > 0 {
			if names := strings.Fields(pieces[1]); len(names) > 0 {
				if _, ok := hostsMap[pieces[0]]; ok {
					hostsMap[pieces[0]] = append(hostsMap[pieces[0]], names...)
				} else {
					hostsMap[pieces[0]] = names
				}
			}
		}
	}

	return hostsMap
}
