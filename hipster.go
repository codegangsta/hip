package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
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

			fmt.Print("> ")
			console.Run(strings.Fields("cmd " + readLine()))
		}
	}
	app.Run(os.Args)
}

func readLine() string {
	buf := bufio.NewReader(os.Stdin)
	line, err := buf.ReadString('\n')

	if err != nil {
		panic(err)

	}

	line = strings.TrimRight(line, "\n")

	return line

}
