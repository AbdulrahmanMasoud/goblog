package cmd

import (
	"fmt"
	"github.com/AbdulrahmanMasoud/goblog/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"

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

/*
- First I init the methode setConfig to get the config values from the yml file and unmarshal it using vipre
- Used the returned configs in main function
*/
func serve() {
	configs := setConfig()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app_name": viper.Get("App.Name"), // We can get the keys through config struct
			"email":    configs.Auth.Email,
			"password": configs.Auth.Password,
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}

func setConfig() config.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error Reading the config")
	}

	var configs config.Config
	err := viper.Unmarshal(&configs)

	if err != nil {
		fmt.Println("Unable to decode into struct, %v", err)
	}

	return configs
}
