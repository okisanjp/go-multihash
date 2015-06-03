package main

import (
	"crypto/md5"
    "crypto/sha1"
    "flag" // コマンドラインのフラグ解析をするパッケージ
    "fmt"
)

func main() {
    flag.Parse() // 引数をパース
    h := []byte(flag.Arg(0)) // 1個目の引数のみ使う
    hash(h)
}

func hash(h []byte){
    fmt.Printf("RAW  : %s\n",h) // 文字列として表示
    fmt.Printf("MD5  : %x\n",md5.Sum(h)) // 
    fmt.Printf("SHA1 : %x\n", sha1.Sum(h))
}
