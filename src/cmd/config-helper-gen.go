package cmd

import (
	"fmt"
	"github.com/go-for-go/yak/src/services/Generator"
	"github.com/spf13/cobra"
)

var (
	InputFile  string
	OutputFile string
)

func init() {
	configHelperGen.PersistentFlags().StringVarP(&InputFile, "input", "i", "", "set input file")
	configHelperGen.MarkPersistentFlagRequired("input")
	configHelperGen.PersistentFlags().StringVarP(&OutputFile, "output", "o", "", "set output file")
	configHelperGen.MarkPersistentFlagRequired("output")

	rootCmd.AddCommand(configHelperGen)
}

var configHelperGen = &cobra.Command{
	Use:   "config-helper-gen",
	Short: `generate object orient config helper `,
	Run: func(cmd *cobra.Command, args []string) {
		input := cmd.Flag("input").Value.String()
		output := cmd.Flag("output").Value.String()

		g := Generator.Gen{}
		err := g.Run(input, output)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Println("ok")
	},
}
