//go:build mage

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/princjef/mageutil/bintool"
)

// download required build tools
func Tools() error {
	return tools()
}

// run the golangci-lint linter
func Lint() error {
	mg.Deps(tools)

	return sh.Run("golangci-lint", "run")
}

// Build build the binary
func Build() error {
	return build()
}

// run tests
func Test() error {
	return test("-short")
}

// run integration tests
func It() error {
	return test()
}

// show HTML coverage report
func Cover() error {
	return cover()
}

// remove temporary build files
func Clean() error {
	return clean()
}

// lint, test, build
func All() error {
	if err := Lint(); err != nil {
		return err
	}

	if err := It(); err != nil {
		return err
	}

	return Build()
}

var Default = All

// ---------------------------------------

var (
	toolsdir string
	workdir  string
)

func init() {
	cwd, err := os.Getwd()
	must(err)

	workdir = filepath.Join(cwd, "build")
	toolsdir = filepath.Join(workdir, ".tools")

	os.MkdirAll(toolsdir, 0o755)

	path := fmt.Sprintf("%s%c%s", toolsdir, os.PathListSeparator, os.Getenv("PATH"))
	os.Setenv("PATH", path)
}

func must(err error) {
	if err != nil {
		mg.Fatal(1, err)
	}
}

func exists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func goinstall(target string) error {
	return sh.RunWith(map[string]string{"GOBIN": toolsdir}, "go", "install", target)
}

// tools downloads k6 golangci-lint configuration, golangci-lint, goreleaser and gotestsum.
func tools() error {
	resp, err := http.Get("https://raw.githubusercontent.com/grafana/k6/master/.golangci.yml")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to download linter configuration")
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(".golangci.yml", content, 0o644)
	if err != nil {
		return err
	}

	if !bytes.HasPrefix(content, []byte("# v")) {
		return errors.New("missing version comment")
	}

	linterVersion := strings.TrimSpace(string(content[3:bytes.IndexRune(content, '\n')]))

	tester, err := bintool.New(
		"gotestsum{{.BinExt}}",
		"1.11.0",
		"https://github.com/gotestyourself/gotestsum/releases/download/v{{.Version}}/gotestsum_{{.Version}}_{{.GOOS}}_{{.GOARCH}}.tar.gz",
		bintool.WithFolder(toolsdir),
	)
	if err != nil {
		return err
	}

	if err := tester.Ensure(); err != nil {
		return err
	}

	os := cases.Title(language.English, cases.Compact).String(runtime.GOOS)
	arch := strings.ReplaceAll(runtime.GOARCH, "amd64", "x86_64")

	releaser, err := bintool.New(
		"goreleaser{{.BinExt}}",
		"1.22.0",
		"https://github.com/goreleaser/goreleaser/releases/download/v{{.Version}}/goreleaser_"+os+"_"+arch+"{{.ArchiveExt}}",
		bintool.WithFolder(toolsdir),
	)
	if err != nil {
		return err
	}

	if err := releaser.Ensure(); err != nil {
		return err
	}

	linter, err := bintool.New(
		"golangci-lint{{.BinExt}}",
		linterVersion,
		"https://github.com/golangci/golangci-lint/releases/download/v{{.Version}}/golangci-lint-{{.Version}}-{{.GOOS}}-{{.GOARCH}}{{.ArchiveExt}}",
		bintool.WithFolder(toolsdir),
	)
	if err != nil {
		return err
	}

	return linter.Ensure()
}

func lint() error {
	mg.Deps(tools)

	return sh.Run("golangci-lint", "run")
}

func coverprofile() string {
	return filepath.Join(workdir, "coverage.txt")
}

func test(args ...string) error {
	mg.Deps(tools)

	maxproc := "4"

	if runtime.GOOS == "windows" {
		maxproc = "1"
	}

	env := map[string]string{
		"GOMAXPROCS": maxproc,
	}

	testargs := []string{"--format"}
	if mg.Verbose() {
		testargs = append(testargs, "standard-verbose")
	} else {
		testargs = append(testargs, "testdox")
	}

	testargs = append(testargs, "--")
	testargs = append(testargs, args...)

	testargs = append(testargs,
		"-count", "1", "-p", maxproc, "-race", "-coverprofile="+coverprofile(), "./...")

	_, err := sh.Exec(env, os.Stdout, os.Stderr, "gotestsum", testargs...)

	return err
}

func build() error {
	mg.Deps(tools)

	var ext string

	if runtime.GOOS == "windows" {
		ext = ".exe"
	}

	_, err := sh.Exec(nil, os.Stdout, os.Stderr, "goreleaser",
		"build",
		"-o",
		filepath.Join(workdir, "tygor"+ext),
		"--snapshot",
		"--clean",
		"--single-target",
	)

	return err

}

func cover() error {
	mg.Deps(test)
	_, err := sh.Exec(nil, os.Stdout, os.Stderr, "go", "tool", "cover", "-html="+coverprofile())
	return err
}

func clean() error {
	return sh.Rm("build")
}
