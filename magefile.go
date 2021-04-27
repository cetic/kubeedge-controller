//+build mage

package main

import (
	"github.com/magefile/mage/sh"

	// mage:import
	_ "go-dev.ec.homeshore.be/csc/internal/mage"
	// mage:import
	_ "go-dev.ec.homeshore.be/csc/plugins/mage"
)

// Run GolangCI-Lint
func Lint() error {
	return sh.RunV("golangci-lint", "run")
}

// Run the tests
func Test() error {
	return sh.RunV("go", "test", "-v", "./internal/...")
}
