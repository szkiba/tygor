package cmd

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//nolint:lll
func parseCmd() *cobra.Command {
	flags := new(parseFlags)

	cmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "parse file",
		Short: "Convert k6 extension's API definition to JSON data model.",
		Long: `From the TypeScript declaration file,  tygor parse subcommand generates the API model in JSON format. The API model can be processed by external programs without the complexity of TypeScript parsing.

The only mandatory argument of the tygor parse subcommand is the name of the declaration file (which file name must end with a .d.ts suffix).		
`,
		Example: "$ " + appname + " parse hitchhiker.d.ts | jq",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return parseRun(checkFileSuffix(args[0]), flags)
		},
		DisableAutoGenTag: true,
	}

	cmd.Flags().
		StringVarP(&flags.output, "output", "o", "", "output file (default: standard output)")

	return cmd
}

type parseFlags struct {
	output string
}

//nolint:forbidigo
func parseRun(src string, flags *parseFlags) error {
	filename := filepath.Clean(src)

	mod, err := readAndParse(filename)
	if err != nil {
		return err
	}

	var file *os.File
	var writer io.Writer

	if len(flags.output) != 0 {
		file, err = os.Create(flags.output)
		if err != nil {
			return err
		}

		writer = file
	} else {
		writer = os.Stdout
	}

	encoder := json.NewEncoder(writer)

	err = encoder.Encode(mod)
	if err != nil {
		return err
	}

	if file != nil {
		return file.Close()
	}

	return nil
}
