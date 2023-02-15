package ready

import (
	"log"
	"os"
)

type Ready struct{}

func New() *Ready{
	r := new(Ready)
	return r
}

func (r *Ready) Get(path string) string{
	data, err :=os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}