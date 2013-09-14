package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bobappleyard/readline"
	"github.com/codegangsta/cli"
	"github.com/wsxiaoys/terminal/color"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
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
							os.Exit(1)
						}

						printResponse(res)

						buf := new(bytes.Buffer)
						body, _ := ioutil.ReadAll(res.Body)
						json.Indent(buf, body, "", "    ")
						s := buf.String()
						s = strings.Replace(s, "@", "@@", -1)
						reg, _ := regexp.Compile("(\".*?:)(.*).?([,\r,\n])")
						s = reg.ReplaceAllString(s, "@b$1@c$2@|$3")
						color.Println(s)
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

func printResponse(res *http.Response) {
	color.Println(res.Proto, statusColor(res.StatusCode)+res.Status)

	for k := range res.Header {
		color.Println(k+":", "@{!c}"+res.Header.Get(k))
	}

	fmt.Println("")
}

func statusColor(code int) string {
	if code >= 400 {
		return "@{r!}"
	} else {
		return "@{g!}"
	}
}
