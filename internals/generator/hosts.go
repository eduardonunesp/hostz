package generator

type HostsGenerator interface{}

type hostsGenerator struct{}

func NewHostsGenerator() HostsGenerator {
	return hostsGenerator{}
}

func (hg hostsGenerator) BuildHostsFile() string {
	return ""
}
