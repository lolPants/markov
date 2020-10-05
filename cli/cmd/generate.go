package cmd

import (
	"bufio"
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
			reader := bufio.NewReader(os.Stdin)
			writer := bufio.NewWriter(os.Stdout)
			defer writer.Flush()

			model := markov.NewModel()
			err := model.ImportReader(reader)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
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
}
