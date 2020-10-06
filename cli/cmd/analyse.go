package cmd

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"

	"github.com/lolPants/markov/cli/pkg/markov"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	analyseCmd = &cobra.Command{
		Use:   "analyse",
		Short: "Read lines from stdin and output a model file to stdout",
		Run: func(cmd *cobra.Command, args []string) {
			if checkIsPipe(os.Stdin) == false {
				fmt.Fprintln(os.Stderr, "lines must be input via shell redirection or piping")

				os.Exit(1)
				return
			}

			if viper.GetBool("gzip-analyse") && checkIsPipe(os.Stdout) == false {
				fmt.Fprintln(os.Stderr, "output must be piped or redirected when `--gzip` is set")

				os.Exit(1)
				return
			}

			reader := bufio.NewReader(os.Stdin)
			writer := bufio.NewWriter(os.Stdout)
			defer writer.Flush()

			model := markov.NewModel()
			err := model.AnalyseReader(reader)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}

			if viper.GetBool("gzip-analyse") {
				gz := gzip.NewWriter(writer)
				err = model.ExportWriter(gz)

				gz.Close()
			} else {
				err = model.ExportWriter(writer)
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(analyseCmd)

	analyseCmd.Flags().BoolP("gzip", "Z", false, "enable gzip compression")
	viper.BindPFlag("gzip-analyse", analyseCmd.Flags().Lookup("gzip"))

	analyseCmd.Flags().BoolP("pretty", "P", false, "pretty print model JSON")
	viper.BindPFlag("pretty", analyseCmd.Flags().Lookup("pretty"))
}
