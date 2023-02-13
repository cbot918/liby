package configy

import (
	"log"

	"github.com/spf13/viper"
)

type Configy struct {
	V viper.Viper
	C Config
}

type Config struct{
	Vagrant string
}


func New() *Configy{
	c := new(Configy)
	c.V = *viper.New()

	return c
}

func (c *Configy) Set(path string, fname string, fsubname string) *Configy {
	c.V.AddConfigPath(path)
	c.V.SetConfigFile(fsubname)
	c.V.SetConfigName(fname)
	
	if err := c.V.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return c
}

func (c *Configy) Get (input interface{}){
	c.V.Unmarshal(&input)
}