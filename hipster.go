package main

import (
	"bytes"
	"fmt"
	"github.com/bobappleyard/readline"
	"github.com/codegangsta/cli"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Action = func(ctx *cli.Context) {
		if len(ctx.Args()) == 0 {
			cli.ShowAppHelp(ctx)
			os.Exit(1)
		}

		address := ctx.Args()[0]

		for {
			console := cli.NewApp()
			console.Action = func(c *cli.Context) {
				fmt.Println("Command not found. Type 'help' for a list of commands.")
			}

			console.Commands = []cli.Command{
				{
					Name:      "get",
					ShortName: "g",
					Usage:     "Make a get request",
					Action: func(c *cli.Context) {
						url := address + c.Args()[0]
						res, err := http.Get(url)
						if err != nil {
							fmt.Println(err)
						}
						buf := new(bytes.Buffer)
						buf.ReadFrom(res.Body)
						s := buf.String()
						fmt.Println(s)
					},
				},
			}

			line, err := readline.String("> ")
			if err == io.EOF {
				break

			}
			if err != nil {
				fmt.Println("error: ", err)
				break

			}
			readline.AddHistory(line)
			console.Run(strings.Fields("cmd " + line))
		}
	}
	app.Run(os.Args)
}
