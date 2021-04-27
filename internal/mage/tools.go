package mage

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/hashicorp/go-getter"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Tools mg.Namespace

// Install all tools
func (Tools) All() {
	mg.SerialDeps(
		Tools.Mage,
		Tools.GolangciLint,
		Tools.Migrate,
		Tools.Nats,
		Tools.Micro,
		Tools.Admin,
		Tools.Jsonvalidator,
		Tools.Token,
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

// Install GolangCI-Lint
func (Tools) GolangciLint() error {
	if _, err := exec.LookPath("golangci-lint"); err == nil {
		fmt.Println("golangci-lint already installed")
		return nil
	}
	fmt.Println("Installing golangci-lint...")
	lintDir, err := ioutil.TempDir("", "golangci-lint")
	if err != nil {
		return err
	}
	defer os.RemoveAll(lintDir)
	dir := fmt.Sprintf("golangci-lint-1.25.0-%s-amd64", runtime.GOOS)
	ext := ".tar.gz"
	if runtime.GOOS == "windows" {
		ext = ".zip"
	}
	zip := fmt.Sprintf("https://github.com/golangci/golangci-lint/releases/download/v1.25.0/%s%s", dir, ext)
	fmt.Printf("Downloading and extracting %s...\n", zip)
	err = getter.Get(lintDir, zip)
	if err != nil {
		return err
	}
	bin := "golangci-lint"
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	src := filepath.Join(lintDir, dir, bin)
	dest := filepath.Join(os.Getenv("GOPATH"), "bin", bin)
	fmt.Printf("Copying %s to %s\n", src, dest)
	return copyFile(src, dest)
}

func (Tools) Migrate() error {
	if _, err := exec.LookPath("migrate"); err == nil {
		fmt.Println("migrate already installed")
		return nil
	}
	fmt.Println("Installing migrate...")
	migrateDir, err := ioutil.TempDir("", "migrate")
	if err != nil {
		return err
	}
	defer os.RemoveAll(migrateDir)
	dir := fmt.Sprintf("migrate.%s-amd64", runtime.GOOS)
	if runtime.GOOS == "windows" {
		dir += ".exe"
	}
	zip := fmt.Sprintf("https://github.com/golang-migrate/migrate/releases/download/v4.11.0/%s.tar.gz", dir)
	fmt.Printf("Downloading and extracting %s...\n", zip)
	err = getter.Get(migrateDir, zip)
	if err != nil {
		return err
	}
	src := filepath.Join(migrateDir, dir)
	bin := filepath.Join(os.Getenv("GOPATH"), "bin", "migrate")
	fmt.Printf("Copying %s to %s\n", src, bin)
	return copyFile(src, bin)
}

// Install Nats server
func (Tools) Nats() error {
	bin := filepath.Join("bin", "nats-server")
	if _, err := os.Stat(bin); err == nil {
		fmt.Println("nats-server already installed")
		return nil
	}
	natsDir, err := ioutil.TempDir("", "nats")
	if err != nil {
		return err
	}
	defer os.RemoveAll(natsDir)
	dir := fmt.Sprintf("nats-server-v2.1.4-%s-amd64", runtime.GOOS)
	zip := fmt.Sprintf("https://github.com/nats-io/nats-server/releases/download/v2.1.4/%s.zip", dir)
	fmt.Printf("Downloading and extracting %s...\n", zip)
	err = getter.Get(natsDir, zip)
	if err != nil {
		return err
	}
	binName := "nats-server"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	src := filepath.Join(natsDir, dir, binName)
	fmt.Printf("Copying %s to %s\n", src, bin)
	return copyFile(src, bin)
}

// Install Micro
func (Tools) Micro() error {
	bin := filepath.Join("bin", "micro")
	if _, err := os.Stat(bin); err == nil {
		fmt.Println("micro already installed")
		return nil
	}
	microDir, err := ioutil.TempDir("", "micro")
	if err != nil {
		return err
	}
	defer os.RemoveAll(microDir)
	dir := fmt.Sprintf("micro-v1.18.0-%s-amd64", runtime.GOOS)
	ext := ".tar.gz"
	if runtime.GOOS == "windows" {
		ext = ".zip"
	}
	zip := fmt.Sprintf("https://github.com/micro/micro/releases/download/v1.18.0/%s%s", dir, ext)
	fmt.Printf("Downloading and extracting %s...\n", zip)
	err = getter.Get(microDir, zip)
	if err != nil {
		return err
	}
	binName := "micro"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	src := filepath.Join(microDir, binName)
	fmt.Printf("Copying %s to %s\n", src, bin)
	return copyFile(src, bin)
}

func (Tools) Admin() error {
	return buildDir("admin", "Admin Tools", "tools")
}

func (Tools) Jsonvalidator() error {
	return buildDir("jsonvalidator", "JSON Validator", "tools")
}

func (Tools) Token() error {
	return buildDir("token", "Token", "tools")
}
