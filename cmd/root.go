package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jameshiew/same/internal/git"
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
		same, err := git.GetSameBranchNames(context.Background(), "master")
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for _, br := range same {
			fmt.Println(br)
		}
	},
}

func Execute(version string, logLevel string) {
	lvl, err := log.ParseLevel(logLevel)
	if lvl == 0 {
		lvl = log.WarnLevel
	}
	log.SetLevel(lvl)
	if err != nil {
		log.Warnf("Error parsing log level: %v", err)
	}
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error running: %v", err)
	}
}

func init() {
	rootCmd.SetVersionTemplate(versionTemplate)
}
