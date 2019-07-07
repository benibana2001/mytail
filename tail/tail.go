package tail

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)
type Mytail struct {
	Data []string
	FileName string
}

func Create(n int, filename string) *Mytail{
	m := Mytail{}
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

	if s[len(s)-1] == "" {
		s = s[:len(s)-1]
	}

	// 最終行が空行の場合は削除

	if n > len(s) {
		n = len(s)
	}
	var ss []string
	var i int
	for i = len(s)-n; i<len(s); i++ {
		ss = append(ss, s[i])
	}
	m.Data = ss
	return &m

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

func (m Mytail) SaveFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Oops! failed to create new file")
	}
	file.Write([]byte(ss2bs(m.Data)))
}

func (m Mytail) Print() {
	for _, v := range m.Data{
		fmt.Println(v)
	}
}

func ss2bs(ss []string) []byte{
	var s string
	for _, v := range ss {
		s = s + v + "\n"
	}
	return []byte(s)
}
