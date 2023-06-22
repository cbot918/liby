package loady

import (
	"fmt"
)

type Loady struct{}

func New() *Loady {
	return &Loady{}
}

func (l *Loady) Run(message string) {
	fmt.Println(message)
}

func (l *Loady) Concurrent(fn func(), times int) {
	for i := 0; i < times; i++ {
		fmt.Println(i)
		go fn()
	}
}

func (l *Loady) Test(fn func()) {
	fn()
}
