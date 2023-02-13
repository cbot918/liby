package main

import (
	"fmt"

	u "github.com/cbot918/liby/util"

	"github.com/cbot918/liby/configy"
)

type Config struct {
	Conf struct {
		Vagrant string
		Foo string
	}
}

func main(){
	cfg := &Config{}
	configy.New().Set("./","app","yaml").Get(cfg)


	fmt.Println(cfg)
	u.Loggj(cfg)
	
}
