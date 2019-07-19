package main

import (
	"sync"
)

func main() {
	s := new(StakingPlugin)

}

type StakingPlugin struct {
	Name 	 string
	Once     *sync.Once
}

var sk *StakingPlugin

func (s *StakingPlugin) New () {
	if nil == sk {

	}
}