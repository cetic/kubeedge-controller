package mage

import (
	"github.com/magefile/mage/mg"
)

type Bin mg.Namespace

// Compile all binaries
func (Bin) All() {
	mg.SerialDeps(
		Bin.Controller,
	)
}

func (Bin) Controller() error {
	return build("controller-srv", "Controller Service")
}
