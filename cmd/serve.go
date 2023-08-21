package cmd

import (
	"github.com/AbdulrahmanMasoud/goblog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "serve",
	Short: "Will serve app in dev env",
	Long:  `The app will be served on ports that's add in config.yml file`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	bootstrap.Serve()
}
