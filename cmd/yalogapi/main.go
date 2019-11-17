package main

import (
	"fmt"
	"log"
	"os"

	"github.com/d2one/yalogapi/internal/yalogapi"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func main() {
	var config string
	app := &cli.App{
		Name:  "yalogapi",
		Usage: "fight the loneliness!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Required:    true,
				Aliases:     []string{"c"},
				Name:        "config",
				Usage:       "config",
				Destination: &config,
			},
		},
		Action: func(c *cli.Context) error {
			var err error
			viper.SetConfigType("yaml")
			viper.SetConfigFile(config)
			fmt.Println(config)
			// read in config file and check your errors
			if err := viper.ReadInConfig(); err != nil {
				fmt.Println(err)
			}
			conf := &yalogapi.Config{}
			err = viper.Unmarshal(conf)
			if err != nil {
				fmt.Printf("unable to decode into config struct, %v", err)
			}
			err = conf.Validate()
			if err != nil {
				fmt.Println(err)
			}
			yalogapi := yalogapi.NewYaLogApi(conf)
			yalogapi.Run()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
