package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gernest/qlql/components/api"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Name = "qlql"
	app.Usage = "Awesome GUI for ql database"
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Starts a server that exposes ql API",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port",
					Usage: "ports to bind the server",
					Value: 8090,
				},
			},
			Action: Server,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Server(ctx *cli.Context) error {
	port := ctx.Int("port")
	s := api.NewServer()
	log.Println("starting ql server at http://localhost:", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s)
}
