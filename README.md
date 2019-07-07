mytail 
====
## Name
mytail - cli tool which resemble `tail` command.  
This package use [urfave/cli](https://github.com/urfave/cli)
for creating CLT tool.  
If you are not familiar with this, please refer to the repository.  
mytail - これは `tail` コマンドに似た CLIツールです。  
CLIツール化のために [urfave/cli](https://github.com/urfave/cli) を使用しています。  
もしこちらに詳しくないようでしたら、恐縮ですがリポジトリをご参照ください。  

## Overview
i made this cli tool for tha exam of `Gopher道場 - 課題2`.
i write requirements of this exam to Description-area.  
この CLIツールは、`Gopher道場 - 課題2` として作成しました。  
課題の詳細については Description に記します。  

## Usage
1.Please clone this repository:  
このリポジトリをクローンしてください。:  
(課題内容の漏洩防止のため、private リポジトリで作成しています。)  
```
git clone https://github.com/benibana2001/mytail.git
```
2.Please install this package to type `go install` :  
`go install` を実行して、このパッケージをインストールしてください。:  
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
-n オプションを指定することで、N行だけ表示することができます。  
もしオプションを設定しない場合、デフォルトとして10行を表示します。
```
mytail -n 3 sample.txt
 or
mytail -n=3 sample.txt
```

##### -o:
You can save output as a file.  
結果をファイルとして保存します。
```
mytail -o new.txt sample.txt
 or
mytail -o=new.txt sample.txt
```

## Description
quote requirements.  
課題内容を引用します。  

>課題2
次の仕様を満たすtailコマンドに似たコマンドをGoで作成してください。
● 引数で渡された1つのファイルの末尾の最大N行をそのまま出力
● Nはデフォルトで10
● オプション-nでNを指定できる
実行例
$ mytail -n=3 hoge.txt
X
Y
Z
追加機能
必須ではありませんが、余力があれば上記の機能に加えて、追加機能を実装してください。
追加機能の仕様は特に制限はなく、本家のtailコマンドにないものでも構いません。
追加機能とは、例えば次のように複数のファイルを取れるようにするというようなもので
す。
$ mytail -n=3 hoge.txt fuga.txt
==> hoge.txt <==
X
Y
Z
==> fuga.txt <==
x
y
z

#### メモ  
##### Gopher道場に応募した理由とGopher道場で学びたいことについて　　
Goのスキルアップ、およびGoを通して低レイヤーの知識、マイクロサービスの知見を高めたいと思っています。
志望した理由ですが、以下の理由から個人での学習に不安を感じているためです。
・これまでスクリプト言語しかまともに触ったことがない(JS, PHP)
・CSの基礎知識があまり豊富でない(OSの仕組みについて基礎から学習中です)
・周囲にGoの経験者がいない
Gopher道場にて現場のノウハウを吸収できたらいいなと思っています。

##### 課題2の解く上で工夫した点や難しいと感じた点について
どの程度のレベルが求められているのかわからず不安でしたが、
機能追加やテストが容易にできるように意識しました。
具体的には、メインとなる構造体を容易し、そこにレシーバーで関数を付与していく形をとりました。
（当然のやり方とは思いますが...）

難しいと感じた点は、ストリームの部分です。
はじめはメモリの節約のために `bufio.NewReader` を使用して改行区切りで１行づつ読み込むのがよいかなと考えました。  
しかし上手く形にできず...。
結局 `ioutil.ReadFile` で全て読み込んだ後に、全ての文字列を改行で区切ってスライスに入れる形としました。
どうにかしたいと思っています。
