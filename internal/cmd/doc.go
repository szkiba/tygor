package cmd

import (
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/szkiba/tygor/internal/doc"
)

//nolint:lll
func docCmd() *cobra.Command {
	flags := new(docFlags)

	cmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "doc file",
		Short: "Generate documentation from k6 extension's API definition.",
		Long: `From the TypeScript declaration file, tygor doc subcommand generates API documentation.

API documentation is generated to standard output in Markdown format by default. If the --html flag is used, the output format will be HTML.

The output can also be saved to a file using the --output flag. In this case, the default format is determined from the file extension: in the case of .htm and .html extensions, it will be in HTML format, otherwise it will be in Markdown format. Using the --html flag, the HTML format can also be forced for other file extensions.

API documentation can also be inserted (and updated) into an existing Markdown or HTML document using the --inject flag. The insertion takes place in the place marked by so-called marker comments:

    <!-- begin:api -->
    generated API documentation goes here
    <!-- end:api -->

The generated API documentation starts at heading level 1 by default. The starting heading level can be specified by using the --heading flag, which can be useful, for example, when inserting into an outer document.

The documentation may include the usual extension documentation sections, such as build instructions, download instructions, a link to the examples folder, etc. The required GitHub repository can be specified using the --github-repo flag. Otherwise, the tygor doc subcommand tries to guess the GitHub repository from the git configuration (if it exists). This automation can be disabled with the --no-auto flag.
By default, GitHub repository and generateable boilerplate sections are automatically detected. This is done by examining the git configuration, the GitHub workflows configuration, and the examples directory.

The only mandatory argument to the doc subcommand is the name of the declaration file (which file name must end with a .d.ts suffix).
`,
		Example: "$ " + appname + " doc -o README.md hitchhiker.d.ts",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return docRun(checkFileSuffix(args[0]), flags)
		},

		DisableAutoGenTag: true,
	}

	cmd.Flags().
		StringVarP(&flags.output, "output", "o", "", "output file (default: standard output)")
	cmd.Flags().
		StringVarP(&flags.outer, "inject", "i", "", "inject into outer file")
	cmd.Flags().
		StringVarP(&flags.template, "template", "t", "", "go template file for markdown generation")
	cmd.Flags().
		BoolVar(&flags.html, "html", false, "enable HTML output (default: based on file ext)")
	cmd.Flags().
		UintVar(&flags.heading, "heading", 1, "initial heading level")
	cmd.Flags().
		StringVar(&flags.githubRepo, "github-repo", "", "GitHub repository (owner/name)")
	cmd.Flags().
		BoolVar(&flags.linkReleases, "link-releases", false, "enable GitHub releases link")
	cmd.Flags().
		BoolVar(&flags.linkPackages, "link-packages", false, "enable GitHub container packages link")
	cmd.Flags().
		BoolVar(&flags.linkExamples, "link-examples", false, "enable examples folder link")
	cmd.Flags().
		BoolVar(&flags.noAuto, "no-auto", false, "disable automatic GitHub repo and link flags detection")

	cmd.MarkFlagsMutuallyExclusive("output", "inject")

	return cmd
}

type docFlags struct {
	output   string
	outer    string
	html     bool
	heading  uint
	template string

	githubRepo   string
	linkReleases bool
	linkPackages bool
	linkExamples bool
	noAuto       bool
}

func (flags *docFlags) toOptions() []doc.Option {
	opts := []doc.Option{
		doc.WithGitHubRepo(flags.githubRepo),
		doc.WithLinkReleases(flags.linkReleases),
		doc.WithLinkPackages(flags.linkPackages),
		doc.WithLinkExamples(flags.linkExamples),
	}

	if flags.html || isHTML(flags.output) || isHTML(flags.outer) {
		opts = append(opts, doc.WithFormat(doc.FormatHTML))
	}

	if flags.heading != 0 {
		heading := flags.heading
		if heading > 5 {
			heading = 5
		}
		opts = append(opts, doc.WithHeading(heading))
	}

	return opts
}

func exists(filename string) bool {
	_, err := os.Stat(filename) //nolint:forbidigo

	return err == nil
}

//nolint:forbidigo
func (flags *docFlags) autoDetect(from string) {
	dir, err := filepath.Abs(from)
	if err != nil {
		return
	}

	var filename string

	for {
		parent := filepath.Dir(dir)
		if parent == dir {
			return
		}

		filename = filepath.Join(dir, ".git", "config")
		if exists(filename) {
			break
		}

		dir = parent
	}

	if len(flags.githubRepo) == 0 {
		content, err := os.ReadFile(filepath.Clean(filename))
		if err != nil {
			return
		}

		re := regexp.MustCompile(
			`(?s)\[remote "origin"\][^[]*url\s*=.*github.com[:]?[/]?([^/]*/[^/.]*)`,
		)

		all := re.FindSubmatch(content)
		if all == nil || len(all) <= 1 {
			return
		}

		flags.githubRepo = string(all[1])
	}

	if exists(filepath.Join(dir, "examples")) {
		flags.linkExamples = true
	}

	filename = filepath.Join(dir, ".github", "workflows", "release.yml")
	if exists(filename) {
		flags.linkReleases = true

		content, err := os.ReadFile(filename)
		if err == nil && strings.Contains(string(content), "ghcr.io") {
			flags.linkPackages = true
		}
	}
}

//nolint:forbidigo
func docRun(src string, flags *docFlags) error {
	mod, err := readAndParse(src)
	if err != nil {
		return err
	}

	var file *os.File
	var writer io.Writer

	if !flags.noAuto {
		flags.autoDetect(filepath.Dir(src))
	}

	opts := flags.toOptions()

	if len(flags.output) != 0 {
		file, err = os.Create(flags.output)
		if err != nil {
			return err
		}

		writer = file
	}

	if len(flags.outer) != 0 {
		outer, e := os.ReadFile(flags.outer)
		if e != nil {
			return e
		}

		opts = append(opts, doc.WithOuter(outer))

		file, err = os.OpenFile(flags.outer, os.O_RDWR, 0o600)
		if err != nil {
			return err
		}

		writer = file
	}

	if writer == nil {
		writer = os.Stdout
	}

	if len(flags.template) != 0 {
		template, e := os.ReadFile(flags.template)
		if e != nil {
			return e
		}

		opts = append(
			opts,
			doc.WithTemplate(string(template)),
			doc.WithTemplateName(filepath.Base(flags.template)),
		)
	}

	out, err := doc.Generate(mod, opts...)
	if err != nil {
		return err
	}

	_, err = writer.Write(out)
	if err != nil {
		return err
	}

	if file != nil {
		return file.Close()
	}

	return nil
}

func isHTML(filename string) bool {
	return strings.HasSuffix(filename, ".html") || strings.HasSuffix(filename, "htm")
}
