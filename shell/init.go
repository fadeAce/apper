package shell

import (
	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
	"strings"
)

func InitShell() (shell *ishell.Shell) {
	shell = ishell.NewWithConfig(&readline.Config{Prompt: "apper > "})

	// display welcome info.
	shell.Println(
		`
------------------  Hally cloud limited  -----------------------

****************************************************************
**  apper is a high performance scrapper.                     **
**                                                            **
**  apper interactive mode turns on !                         **
**                                                            **	
**    - you can interactive with apper daemon with commands,  **
**      for more information, try help.                       **
****************************************************************

----------------------------------------------------------------
			`,
	)

	// register a function for "terminate" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "terminate",
		Help: "terminate the daemon server, then exit apper shell",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})
	// register a function for "terminate" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "stats",
		Help: "statistics for scrapper data",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})
	// register a function for "terminate" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "state",
		Help: "state for apper in time",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})
	// register a function for "terminate" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "stop",
		Help: "stop a work from listed work id",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})
	// register a function for "terminate" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "start",
		Help: "start a work from a new config file",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})
	// register a function for "terminate" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "terminate",
		Help: "terminate the daemon server, then exit apper shell",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})
	// register a function for "terminate" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "ls",
		Help: "list all works at queue",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})
	return
}
