package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// Create a new CLI application with the urfave/cli package
	app := &cli.App{
		Name:  "HealthChecker", // Set the application name
		Usage: "A tiny tool that checks whether a website is running or down",
		// Flags define the command-line options for the application
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Doamin number to check.",
				Required: true,
			},
			// Another StringFlag for the port number, which is not required (defaults to 80 if not provided)
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
