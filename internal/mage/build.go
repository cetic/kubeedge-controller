package mage

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/magefile/mage/sh"
)

func buildDir(bin, name, dir string) error {
	output := fmt.Sprintf("bin/%s", bin)
	fmt.Printf("Building %s (%s)\n", name, output)
	if err := sh.Rm(output); err != nil {
		return err
	}
	ldflags := "-X main.Version=${VERSION} -X main.Revision=${REVISION} -X main.BuildTime=${BUILD_TIME} -s -w"
	args := []string{
		"build",
		"-ldflags",
		ldflags,
		"-a",
		"-installsuffix",
		"cgo",
		"-o",
		output,
	}
	args = append(args, goFiles(fmt.Sprintf("%s/%s", dir, bin))...)
	env, err := flagEnv()
	if err != nil {
		return err
	}
	if err := sh.RunWith(env, "go", args...); err != nil {
		return err
	}
	return nil
}

func build(bin, name string) error {
	return buildDir(bin, name, "cmd")
}

func flagEnv() (map[string]string, error) {
	hash, err := sh.Output("git", "rev-parse", "--short=10", "HEAD")
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"CGO_ENABLED": "0",
		"VERSION":     "0.1.0-git",
		"REVISION":    hash,
		"BUILD_TIME":  time.Now().Format(time.RFC3339),
	}, nil
}

func goFiles(dir string) []string {
	res := make([]string, 0)
	for _, file := range []string{"main.go", "plugin.go"} {
		p := fmt.Sprintf("%s/%s", dir, file)
		if _, err := os.Stat(p); err == nil {
			res = append(res, p)
		}
	}
	return res
}

func copyFile(src, dest string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dest, input, 0755)
	if err != nil {
		return err
	}
	return nil
}
