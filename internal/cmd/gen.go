package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/szkiba/tygor/internal/gen"
	"github.com/szkiba/tygor/internal/idl"
)

var errInvalidFilename = errors.New("invalid filename")

const dtsSuffix = ".d.ts"

func checkFileSuffix(filename string) string {
	if !strings.HasSuffix(filename, dtsSuffix) {
		cobra.CheckErr(
			fmt.Errorf("%w: %s should ends with %s", errInvalidFilename, filename, dtsSuffix),
		)
	}

	return filename
}

//nolint:lll
func genCmd() *cobra.Command {
	cmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "gen file",
		Short: "Generate golang source code from k6 extension's API definition.",
		Long: `From the TypeScript declaration file, tygor gen subcommand generates the go interfaces needed to implement the API, as well as the binding code between the go implementation and the JavaScript runtime. In addition, tygor gen subcommand is also able to generate a skeleton implementation to help create a go implementation.

The skeleton file can be used as a sample for the implementation. Since it contains a special go build tag (//go:build skeleton), its presence will not interfere with the real implementation. To start the implementation, simply copy the skeleton file under a different name (or rename it) and delete the comments at the beginning of the file. If the declaration file changes, the bindings and skeleton can be regenerated at any time, and the skeleton can be used to help implement the changes.

The only mandatory argument is the name of the declaration file (which file name must end with a .d.ts suffix).
`,
		Example: "$ " + appname + " gen --skeleton hitchhiker.d.ts",
		Args:    cobra.ExactArgs(1),

		SilenceUsage:      true,
		DisableAutoGenTag: true,
	}

	cmd.RunE = genCmdRun(cmd)

	return cmd
}

func genCmdRun(cmd *cobra.Command) func(cmd *cobra.Command, args []string) error {
	flags := new(genFlags)

	cmd.Flags().
		StringVarP(&flags.output, "output", "o", "", "output directory (default: same as input)")
	cmd.Flags().
		BoolVarP(&flags.skeleton, "skeleton", "s", false, "enable skeleton generation (default: disabled)")
	cmd.Flags().
		StringVarP(&flags.pkg, "package", "p", "", "go package name (default: module name)")

	return func(cmd *cobra.Command, args []string) error {
		return genRun(checkFileSuffix(args[0]), flags)
	}
}

type genFlags struct {
	output   string
	skeleton bool
	pkg      string
}

//nolint:forbidigo
func readAndParse(src string) (*idl.Module, error) {
	filename := filepath.Clean(src)

	source, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return idl.Parse(filename, string(source))
}

func genRun(src string, flags *genFlags) error {
	filename := filepath.Clean(src)

	mod, err := readAndParse(filename)
	if err != nil {
		return err
	}

	output := flags.output
	if len(output) == 0 {
		output = filepath.Dir(filename)
	}

	pkg := flags.pkg
	if len(pkg) == 0 {
		pkg = mod.Name
	}

	if len(pkg) == 0 {
		pkg = strings.TrimSuffix(filepath.Base(filename), dtsSuffix)
	}

	opts := []gen.Option{gen.WithGenerator(appname), gen.WithPackage(pkg)}

	err = gen.Code(mod, opts...).Save(filepath.Join(output, pkg+CodeSuffix+".go"))
	if err != nil {
		return err
	}

	if !flags.skeleton {
		return nil
	}

	opts = append(opts, gen.WithTag(SketchTag))

	return gen.Sketch(mod, opts...).Save(filepath.Join(output, pkg+SketchSuffix+".go"))
}

const (
	// CodeSuffix used as filename suffix for generated go bindings code.
	CodeSuffix = "_bindings"

	// SketchSuffix used as filename suffix for generated go empty implementation.
	SketchSuffix = "_skeleton"

	// SketchTag used as go build tag for generated go empty implementation.
	SketchTag = "skeleton"
)
