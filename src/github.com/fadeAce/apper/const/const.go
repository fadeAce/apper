package _const

const (
	DEFAULT_SUM_PIPE  = 200
	DEFAULT_TIMEOUT  = 15
	DEFAULT_SUM_VALUE = iota
	CACHING_STATE_NORMAL
	CACHING_STATE_ERROR
	PIP_IDLE
	PIP_WORKING
	PIP_DONE
)

const (
	CMD_START = "start"
	CMD_STOP  = "stop"
	CMD_LS    = "ls"
)

const (
	TASK_TXN_PREFFIX = "TXN_"
	KEY_CONFIG       = "config"
	PIP_PREFFIX      = "pip_"
)

const (
	TYPE_JSON    = "json"
	TYPE_HTML    = "html"
	TYPE_INVALID = "else"
)
