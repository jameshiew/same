package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jameshiew/same/internal/branch"
)

const (
	name      = "same"
	homepage  = "https://github.com/jameshiew/same"
	docstring = "same lists Git branches which have no diff with master"
)

const versionTemplate = `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "%s" .Version}}
`

var rootCmd = &cobra.Command{
	Use:   name,
	Short: docstring,
	Long:  docstring + " - homepage is " + homepage,
	Run: func(cmd *cobra.Command, args []string) {
		same, err := branch.GetSame(cmd.Context(), "master")
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for _, br := range same {
			fmt.Println(br)
		}
	},
}

func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func init() {
	rootCmd.SetVersionTemplate(versionTemplate)
	log.SetLevel(log.WarnLevel) // TODO make configurable via a flag
}
