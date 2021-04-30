package mage

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"io/ioutil"
	"os"
	"os/exec"
)

type Tools mg.Namespace

// Install all tools
func (Tools) All() {
	mg.SerialDeps(
		Tools.Mage,
	)
}

// Install Mage
func (Tools) Mage() error {
	if _, err := exec.LookPath("mage"); err == nil {
		fmt.Println("mage already installed")
		return nil
	}
	fmt.Println("Installing mage...")
	mageDir, err := ioutil.TempDir("", "mage")
	if err != nil {
		return err
	}
	defer os.RemoveAll(mageDir)
	err = sh.Run("git", "clone", "https://github.com/magefile/mage", mageDir)
	if err != nil {
		return err
	}
	cmd := exec.Command("go", "run", "bootstrap.go")
	cmd.Dir = mageDir
	return cmd.Run()
}
