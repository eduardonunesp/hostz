package parser

import (
	"os"
	"strings"

	"github.com/eduardonunesp/hostz/internals/model"
	"github.com/pkg/errors"
)

type HostsParser interface {
	ReadHostsFile(hostsPath string) ([]byte, error)
	ParseProfile(hostsFileContent []byte) string
	ParseHosts(hostsFileContent []byte) model.HostList
}

type hostsParser struct{}

func NewHostsParser() HostsParser {
	return hostsParser{}
}

func (hs hostsParser) ReadHostsFile(hostsPath string) ([]byte, error) {
	bs, err := os.ReadFile(hostsPath)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read hosts file")
	}

	return bs, nil
}

func (hs hostsParser) ParseProfile(hostsFileContent []byte) string {
	lines := strings.Split(strings.Trim(string(hostsFileContent), " \t\r\n"), "\n")

	if len(lines) == 0 {
		return "No profile specification found"
	}

	return lines[0]
}

func (hs hostsParser) ParseHosts(hostsFileContent []byte) model.HostList {
	hostList := model.HostList{}
	for _, line := range strings.Split(strings.Trim(string(hostsFileContent), " \t\r\n"), "\n") {
		line = strings.Replace(strings.Trim(line, " \t"), "\t", " ", -1)
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}

		pieces := strings.SplitN(line, " ", 2)

		if len(pieces) > 1 && len(pieces[0]) > 0 {
			var alias string

			if len(pieces) > 2 {
				alias = pieces[2]
			}

			hostList = append(hostList, model.Host{
				IP:    pieces[0],
				Name:  pieces[1],
				Alias: alias,
			})
		}
	}

	return hostList
}
