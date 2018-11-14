package types

type ConfJ struct {
	Sites map[string]SiteJ `json:"sites"`
}

type ApperConfJ struct {
	Database string `json:"database"`
}

type SingleJ struct {
	Type string
	Rule string
	Key  string
}

type SiteJ struct {
	Single []SingleJ `json:"singles"`
}

type Command struct {
	Configs ConfJ  `json:"config"`
	Cmd     string `json:"cmd"`
}
