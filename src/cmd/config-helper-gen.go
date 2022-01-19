package cmd

import (
	"fmt"
	"github.com/go-for-go/yak/src/services/Generator"
	"github.com/spf13/cobra"
)

var (
	InputFile  string
	OutputFile string
	PgGo10     bool
)

func init() {
	configHelperGen.PersistentFlags().StringVarP(&InputFile, "input", "i", "", "set input file")
	configHelperGen.MarkPersistentFlagRequired("input")
	configHelperGen.PersistentFlags().StringVarP(&OutputFile, "output", "o", "", "set output file")
	configHelperGen.MarkPersistentFlagRequired("output")
	configHelperGen.PersistentFlags().BoolVar(&PgGo10, "pg10", false, "use pg-go 10 version")

	rootCmd.AddCommand(configHelperGen)
}

var configHelperGen = &cobra.Command{
	Use:   "config-helper-gen",
	Short: `generate object orient config helper `,
	Run: func(cmd *cobra.Command, args []string) {
		input := cmd.Flag("input").Value.String()
		output := cmd.Flag("output").Value.String()
		pg10, err := cmd.Flags().GetBool("pg10")
		if err != nil {
			fmt.Printf("error when get flag pg10, set value as false")
			pg10 = false
		}

		g := Generator.Gen{Pg10: pg10}
		err = g.Run(input, output)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Println("ok")
	},
}
