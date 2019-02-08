package utils

import (
	"emailservice/constants"
	"github.com/urfave/cli"
	"log"
	"os"
)

func InitFlags() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "userName",
			Usage:       "User Name for SMTP server",
			Destination: &constants.Username,
			EnvVar:      "USER_NAME",
		},
		cli.StringFlag{
			Name:        "password",
			Usage:       "Password For SMTP server",
			Destination: &constants.Password,
			EnvVar:      "PASSWORD",
		},
		cli.StringFlag{
			Name:        "fromEmail",
			Usage:       "From email address",
			Destination: &constants.FromAddress,
			EnvVar:      "F_ADDR",
		},
	}
	app.Action = func(c *cli.Context) error {
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}
