package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/99designs/gqlgen/internal/gopath"
	"github.com/spf13/cobra"
)

var configFilename string
var verbose bool

var output string
var models string
var schemaFilename string
var packageName string
var modelPackageName string
var serverFilename string

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFilename, "config", "c", "", "the file to configuration to")
	rootCmd.PersistentFlags().StringVarP(&serverFilename, "server", "s", "", "the file to write server to")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "show logs")

	rootCmd.PersistentFlags().StringVar(&output, "out", "", "the file to write to")
	rootCmd.PersistentFlags().StringVar(&models, "models", "", "the file to write the models to")
	rootCmd.PersistentFlags().StringVar(&schemaFilename, "schema", "", "the graphql schema to generate types from")
	rootCmd.PersistentFlags().StringVar(&packageName, "package", "", "the package name")
	rootCmd.PersistentFlags().StringVar(&modelPackageName, "modelpackage", "", "the package name to use for models")
}

var rootCmd = &cobra.Command{
	Use:   "gqlgen",
	Short: "go generate based graphql server library",
	Long: `This is a library for quickly creating strictly typed graphql servers in golang.
			See https://gqlgen.com/ for a getting started guide.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to determine current workding dir: %s\n", err.Error())
			os.Exit(1)
		}

		if !gopath.Contains(pwd) {
			fmt.Fprintf(os.Stderr, "gqlgen must be run from inside your $GOPATH\n")
			os.Exit(1)
		}
		if verbose {
			log.SetFlags(0)
		} else {
			log.SetOutput(ioutil.Discard)
		}
	},
	Run: genCmd.Run, // delegate to gen subcommand
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
