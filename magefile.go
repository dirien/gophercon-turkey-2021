//go:build mage
// +build mage

package main

import (
	"bytes"
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"os/exec"
)

var Default = Build

var (
	platforms = []string{"darwin", "linux", "windows"}
	archs     = []string{"amd64"}
)

const (
	binary  = "gophercon-turkey-2021"
	ldflags = "-X main.version=%s -X main.build=%s"
	version = "0.1.0"
	dist    = "dist"
)

func getBuild() string {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		return "unknown"
	}
	return string(cmdOutput.Bytes())
}

func Test() error {
	return sh.RunV("go", "test", "-v", "./...")
}

func Lint() error {
	fmt.Println("linting...")
	return sh.RunV("golangci-lint", "run", "--deadline", "10m", "./...")
}

func Docker() error {
	fmt.Println("building docker image...")
	mg.Deps(BuildAll)
	return sh.RunV("docker", "build", "-t", fmt.Sprintf("%s:%s", binary, version), ".")
}

func Build() error {
	mg.Deps(Clean, InstallDeps)
	fmt.Println("Building...")
	return sh.Run("go", "build", "-o", fmt.Sprintf("%s/%s", dist, binary), "-ldflags", fmt.Sprintf(ldflags, version, getBuild()), ".")
}

func BuildAll() error {
	mg.Deps(Clean)
	fmt.Println("Building  all...")
	for _, platform := range platforms {
		for _, arch := range archs {
			sh.Run("go", "build", "-o", fmt.Sprintf("%s/%s-%s-%s", dist, binary, platform, arch), "-ldflags", fmt.Sprintf(ldflags, version, getBuild()), ".")
		}
	}
	return nil
}

func InstallDeps() error {
	fmt.Println("Installing Deps...")
	return sh.Run("go", "mod", "tidy")
}

func Clean() {
	fmt.Println("cleaning...")
	os.RemoveAll(dist)
}
