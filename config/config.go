package config

type Conf struct {
	Sites    map[string]Site `yaml:"sites"`
	Database string          `yaml:"database"`
}

type Single struct {
	Type string
	Rule string
	Key  string
}

//type Site struct {
//	Single []Single `yaml:"singles"`
//}
type Site struct {
	Single []Single `yaml:"singles"`
}
