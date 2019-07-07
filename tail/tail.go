package tail

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func Tail(n int, filename string) {


	// []byteとしてファイルを読み込む
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Oops! Error was occurred, Please ensure filename was input")
		fmt.Println("Input file name is: ", filename)
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


	/*
	////
	// Stream で作成
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%v", line)
		if err == io.EOF {
			break
		}
	}
	*/
}
