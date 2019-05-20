package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
			fmt.Println("Error : please set file name")
			os.Exit(1)
		}
		if c.Int("n") != 10 {
			N = c.Int("n")
		}

		// 実行
		tail(N, filename)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func tail(n int, filename string) {
	// []byteとしてファイルを読み込む
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(2)
	}

	// 配列に変換
	r := regexp.MustCompile("(\r\n|\n\r|\n|\r)")
	s := r.Split(string(content), -1)

	// 配列を逆順に並び替える
	var ss []string
	for i := range s {
		ss = append(ss, s[len(s)-1-i])
	}
	// 書き出す
	count := n
	if n > len(ss) {
		count = len(ss)
	}

	var i int
	for i = count - 1; i >= 0; i-- {
		fmt.Println(ss[i])
	}
}
