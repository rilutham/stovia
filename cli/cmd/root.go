package cmd

import (
	"fmt"

	"rilutham/stovia/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version bool
)

// Root :nodoc:
var Root = &cobra.Command{
	Use: "stovia",
}

func root(cmd *cobra.Command, args []string) (err error) {
	if version {
		fmt.Println(config.Version())
		return nil
	}

	return cmd.Usage()
}

func init() {
	Root.Flags().BoolVarP(&version, "version", "v", false, "show application version")
	Root.PersistentFlags().StringP("config", "c", "", "specify config path")
	viper.BindPFlag("config_path", Root.PersistentFlags().Lookup("config"))

	Root.RunE = root
	Root.AddCommand(
		MigrationCmd,
	)
}
