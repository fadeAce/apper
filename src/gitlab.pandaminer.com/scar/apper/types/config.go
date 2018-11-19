package types

type Conf struct {
	Sites map[string]Site `yaml:"sites"`
}

type ApperConf struct {
	Database       string `yaml:"database"`
	Nats           string `yaml:"nats"`
	CushionSize    int    `yaml:"task_pool"`
	ThreadPoolSize int    `yaml:"th_pool"`
	Timeout        int    `yaml:"time_out"`
}

type Single struct {
	Type string
	Rule string
	Key  string
}

type Site struct {
	Single []Single `yaml:"singles"`
}
