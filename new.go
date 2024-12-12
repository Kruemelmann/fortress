package fortress

import (
	"log"

	"github.com/kruemelmann/fortress/src/basic"
)

type FortressMode int

const (
	BasicAuth FortressMode = iota
)

func New(mode FortressMode) *basic.BasicAuth {
	switch mode {
	case BasicAuth:
		return &basic.BasicAuth{}
	default:
		log.Fatal("Mode is unknown to fortress\n")
	}
	return nil
}
