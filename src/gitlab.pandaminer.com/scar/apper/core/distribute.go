package core

import (
	"gitlab.pandaminer.com/scar/apper/types"
	"github.com/nats-io/go-nats"
	"encoding/json"
	"gitlab.pandaminer.com/scar/apper/const"
)

func Distribute(m *nats.Msg) {
	cmd := types.Command{}
	json.Unmarshal(m.Data, &cmd)
	switch cmd.Cmd {
	case _const.CMD_START:
		// add task to cushion
		sitemap := cmd.Configs
		task := generate(sitemap)
		task.matchPIP()
		task.run()
	case _const.CMD_STOP:

	case _const.CMD_LS:

	}
}

func readConfig() {

}
