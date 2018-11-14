package client

import (
	"gitlab.pandaminer.com/scar/apper/types"
	"github.com/nats-io/go-nats"
	"encoding/json"
	"gitlab.pandaminer.com/scar/apper/const"
	"gitlab.pandaminer.com/scar/apper/core"
)

func Distribute(m *nats.Msg) {
	cmd := types.Command{}
	json.Unmarshal(m.Data, &cmd)
	switch cmd.Cmd {
	case _const.CMD_START:
		// add task to cushion
		sitemap := cmd.Configs
		// todo: sitemap - test
		task := core.Generate(sitemap)
		// todo: Pending
		core.Panel.Pending(task)
	case _const.CMD_STOP:

	case _const.CMD_LS:

	}
}