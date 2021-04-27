package mage

import (
	"fmt"
	"os/exec"

	"github.com/magefile/mage/mg"
)

type Proto mg.Namespace

// Generate all proto
func (Proto) All() {
	mg.SerialDeps(
		Proto.Configuration,
		Proto.Invoicing,
		Proto.Identityaccess,
		Proto.Metering,
	)
}

func (Proto) Configuration() error {
	return proto("configuration")
}

func (Proto) Invoicing() error {
	return proto("invoicing")
}

func (Proto) Identityaccess() error {
	return proto("identityaccess")
}

func (Proto) Metering() error {
	return proto("metering")
}

func proto(service string) error {
	fmt.Printf("Generating proto for %s\n", service)
	args := []string{
		"--proto_path", ".",
		"--micro_out", ".",
		"--go_out", ".",
		fmt.Sprintf("proto/%s/%s.proto", service, service),
	}
	cmd := exec.Command("protoc", args...)
	cmd.Dir = fmt.Sprintf("internal/micro/%s", service)
	return cmd.Run()
}
