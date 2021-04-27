//+build mage

package main

import (
	"github.com/magefile/mage/sh"

	// mage:import
	_ "github.com/cetic/kubeedge-controller/internal/mage"
)


// Run the tests
func Test() error {
	return sh.RunV("go", "test", "-v", "./internal/...")
}
