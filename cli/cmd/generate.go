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
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Read model from stdin and output generated lines to stdout",
		Run: func(cmd *cobra.Command, args []string) {
			if checkIsPipe(os.Stdin) == false {
				fmt.Fprintln(os.Stderr, "model must be input via shell redirection or piping")

				os.Exit(1)
				return
			}

			reader := bufio.NewReader(os.Stdin)
			writer := bufio.NewWriter(os.Stdout)
			defer writer.Flush()

			model := markov.NewModel()

			if viper.GetBool("gzip-generate") {
				gz, err := gzip.NewReader(reader)
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
				}

				err = model.ImportReader(gz)
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
				}
			} else {
				err := model.ImportReader(reader)
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
				}
			}

			lines := viper.GetInt("lines")
			for i := 0; i < lines; i++ {
				writer.WriteString(model.Generate())
				writer.WriteRune('\n')
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().UintP("lines", "L", 1, "number of lines to generate")
	viper.BindPFlag("lines", generateCmd.Flags().Lookup("lines"))

	generateCmd.Flags().BoolP("gzip", "Z", false, "enable gzip compression")
	viper.BindPFlag("gzip-generate", generateCmd.Flags().Lookup("gzip"))
}
