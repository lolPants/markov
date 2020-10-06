package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lolPants/markov/cli/pkg/markov"
	"github.com/spf13/cobra"
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

			reader := bufio.NewReader(os.Stdin)
			writer := bufio.NewWriter(os.Stdout)
			defer writer.Flush()

			model := markov.NewModel()
			err := model.AnalyseReader(reader)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}

			err = model.ExportWriter(writer)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(analyseCmd)
}
