mytail 
====
## Name
mytail - cli tool which resemble `tail` command.  
This package use [urfave/cli](https://github.com/urfave/cli)
for creating CLT tool.  

## Overview
i made this cli tool for my practice.

## Usage
1.Please clone this repository:  
```
git clone https://github.com/benibana2001/mytail.git
```
2.Please install this package to type `go install` :    
```
cd $GOPATH/src/github.com/benibana2001/mytail
go install
```
### run
example:  
```
mytail sample.txt
```

### add flag (option)
##### -n: 
You can display only N lines.
If you don't set, display default lines (N = 10).  
```
mytail -n 3 sample.txt
 or
mytail -n=3 sample.txt
```

##### -o:
You can save output as a file.  
```
mytail -o new.txt sample.txt
 or
mytail -o=new.txt sample.txt
```

#### 改善点  
メインとなる構造体を容易し、そこにレシーバーで関数を付与していく形をとっている。
当初はメモリの節約のためにストリームのままio.Writerに持って行こうと考えたが、実力不足により実装できなかった。
妥協して、 `ioutil.ReadFile` で全て読み込みstringに変換した後に、改行で区切ってスライスに入れる形としている。
