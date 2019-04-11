package cmd

import (
	"rilutham/stovia/config"
	"rilutham/stovia/lib/log"
	"strconv"

	"rilutham/stovia/lib/pq/migration"

	"github.com/spf13/cobra"
)

// MigrationCmd :nodoc:
var MigrationCmd = &cobra.Command{
	Use:   "migrate {up|down|drop}",
	Short: "Migrate database",
}

func migrateRun(cmd *cobra.Command, args []string) (err error) {
	if len(args) < 1 {
		log.For("CLI", "migrate").Error("migrate {up|down|drop}")
		return cmd.Usage()
	}

	directionMap := map[string]migration.Direction{
		"up":   migration.Up,
		"down": migration.Down,
		"drop": migration.Drop,
	}

	steps, err := cmd.PersistentFlags().GetString("steps")
	if err != nil {
		log.For("CLI", "migrate").Error(err)
		return err
	}

	s, err := strconv.Atoi(steps)
	if err != nil {
		log.For("CLI", "migrate").Error(err)
		return err
	}

	return migration.Run(config.DatabaseDSN(), directionMap[args[0]], s)
}

func migrate(cmd *cobra.Command, args []string) (err error) {
	return cmd.Usage()
}

func init() {
	MigrationCmd.RunE = migrateRun
	MigrationCmd.PersistentFlags().StringP("steps", "s", "0", "migration steps")
}
