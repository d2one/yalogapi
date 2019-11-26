package cmd

import (
	"fmt"
	"os"

	"github.com/d2one/yalogapi/internal/yalogapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runCommand,
}

var conf *yalogapi.Config

func runCommand(cmd *cobra.Command, args []string) {
	yalogapi := yalogapi.NewYaLogApi(conf)
	yalogapi.Run()
}

func init() {
	cobra.OnInitialize(initConfig)

	runCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "AWS region (required)")
	runCmd.MarkFlagRequired("config")

	runCmd.PersistentFlags().StringP("mode", "m", "", "available values [history, regular, regular_early]")
	viper.BindPFlag("mode", runCmd.PersistentFlags().Lookup("mode"))

	runCmd.PersistentFlags().String("startDate", "", "startDate")
	viper.BindPFlag("startDate", runCmd.PersistentFlags().Lookup("startDate"))

	runCmd.PersistentFlags().String("endDate", "", "endDate")
	viper.BindPFlag("endDate", runCmd.PersistentFlags().Lookup("endDate"))

	runCmd.PersistentFlags().StringP("source", "s", "", "available values [hits, visits]")
	viper.BindPFlag("source", runCmd.PersistentFlags().Lookup("source"))

	rootCmd.AddCommand(runCmd)

}

func initConfig() {
	if cfgFile == "" {
		return
	}

	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
		os.Exit(1)
	}
	chConf := viper.New()
	chConf.SetConfigFile("configs/types.yaml")
	if err := chConf.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Using config for types file:", chConf.ConfigFileUsed())
	confTypes := &yalogapi.Types{}
	if err := chConf.Unmarshal(confTypes); err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
		os.Exit(1)
	}
	conf.Types = confTypes
	if err := conf.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
