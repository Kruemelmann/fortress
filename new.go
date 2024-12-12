package fortress

import (
	"log"

	"github.com/kruemelmann/fortress/src/basic"
	"github.com/kruemelmann/fortress/src/types"
)

type FortressMode int

const (
	BasicAuth FortressMode = iota
)

func New(mode FortressMode) types.Fortress {
	switch mode {
	case BasicAuth:
		return &basic.BasicAuth{}
	default:
		log.Fatal("Mode is unknown to fortress\n")
	}
	return nil
}
