package types

type CmdJSON struct {
	Cmd   string            `json:"cmd"`
	Param map[string]string `json:"param"`
}

type Apperserver struct {
	Cfg *ApperConf
}
