// Package cmd contains tygor CLI interface.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var (
	version = "dev"
	appname = "tygor"
)

//nolint:lll
func rootCmd() *cobra.Command {
	root := &cobra.Command{ //nolint:exhaustruct
		Use:   appname + " file",
		Short: "CLI tool that enables the development of k6 extensions with an API-First approach.",
		Long: `The functionality of k6 can be extended using JavaScript Extensions, which can be created in the go programming language. Tygor allows you to develop these extensions using an API-First approach. A TypeScript declaration file can be used as IDL to define the JavaScript API of the extension.

From the TypeScript declaration file, tygor generates the go interfaces needed to implement the API, as well as the binding code between the go implementation and the JavaScript runtime. In addition, tygor is also able to generate a skeleton implementation to help create a go implementation.

The skeleton file can be used as a sample for the implementation. Since it contains a special go build tag (//go:build skeleton), its presence will not interfere with the real implementation. To start the implementation, simply copy the skeleton file under a different name (or rename it) and delete the comments at the beginning of the file. If the declaration file changes, the bindings and skeleton can be regenerated at any time, and the skeleton can be used to help implement the changes.

The only mandatory argument is the name of the declaration file (which file name must end with a .d.ts suffix). In addition, different flags can be used to modify the generation output. 

The tygor command generates go source code by default, but it can also generate other outputs. Other outputs can be generated using subcommands. Using it without the subcommand is equivalent to using the gen subcommand.

Use the -h flag to get detailed help on subcommands and flags.
`,
		Example:           "$ " + appname + " --skeleton hitchhiker.d.ts",
		Version:           version,
		Args:              cobra.ExactArgs(1),
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	root.SetVersionTemplate(
		`{{with .Name}}{{printf "%s" .}}{{end}}{{printf " version %s\n" .Version}}`,
	)

	root.RunE = genCmdRun(root)

	root.AddCommand(genCmd())
	root.AddCommand(parseCmd())
	root.AddCommand(docCmd())

	return root
}

// Execute is the main CLI entry point.
func Execute() {
	root := rootCmd()

	err := root.Execute()
	if err != nil {
		os.Exit(1) //nolint:forbidigo
	}
}
