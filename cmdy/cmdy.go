package cmdy

import (
	"os/exec"

	u "github.com/cbot918/liby/util"
)

type Cmdy struct {
	Command []string
	ExecString string
	// Cmd *exec.Cmd
}

func New() *Cmdy{
	c := new(Cmdy)

	c.Command = []string{"touch 1.test","touch 2.test","touch 3.test"}

	for _,val := range c.Command {
		c.ExecString += val+";"
	}

	return c
}


func (c *Cmdy) Run(){
	cmd :=exec.Command("/bin/sh", "-c", c.ExecString)
	err := cmd.Run()
	u.Checke(err, "run error")
}