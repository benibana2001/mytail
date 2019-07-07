package main

import (
	"fmt"
	"github.com/benibana2001/mytail/tail"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "n",
			Value: "10",
			Usage: "length of tail command",
		},
	}

	app.Action = func(c *cli.Context) error {
		// 初期値の設定
		N := 10
		var filename string
		if c.NArg() > 0 {
			filename = c.Args().Get(0)
		}else{
			fmt.Println("Oops! please set file name")
			os.Exit(1)
		}
		if c.Int("n") != 10 {
			N = c.Int("n")
		}

		// 実行
		tail.Tail(N, filename)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


