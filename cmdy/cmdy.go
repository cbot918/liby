package cmdy

import (
	"os/exec"

	u "github.com/cbot918/liby/util"
)

type Cmdy struct {
	// Command []string
	// ExecString string
	// Cmd *exec.Cmd
}

func New() *Cmdy{
	c := new(Cmdy)
	return c
}

func (c *Cmdy) Run(input []string){

	var execStr string

	for _,val := range input {
		execStr += val+";"
	}

	cmd :=exec.Command("/bin/sh", "-c", execStr)
	err := cmd.Run()
	u.Checke(err, "run error")
}

func (c *Cmdy) RunAndGet(input []string) string{
		var execStr string

	for _,val := range input {
		execStr += val+";"
	}

	out, err:=exec.Command("/bin/sh", "-c", execStr).Output()
	// err := cmd.Run()
	u.Checke(err, "run error")
	return string(out)
}