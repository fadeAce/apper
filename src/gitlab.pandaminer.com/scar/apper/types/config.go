package types

type Conf struct {
	Sites map[string]Site `yaml:"sites"`
}

type ApperConf struct {
	Database string `yaml:"database"`
}

type Single struct {
	Type string
	Rule string
	Key  string
}

type Site struct {
	Single []Single `yaml:"singles"`
}
