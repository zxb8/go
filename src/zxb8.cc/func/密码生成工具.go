package main

import "flag"

var(
	length int
	charset string
)
const (
	NumStr = "0123456789"
	CharStr = "abcdefghijklmnopqrstuvwxyz"

)

func parseArgs(){
	flag.IntVar(&length,"l",16,"-l 生成密码的长度")
	flag.StringVar(&charset,"t","num",`-t 指定生成密码的字符集,
		num:只是用数字[0-9],
		char:只使用英文字母[a-zA-z],
		mix:使用数字和字母,
		advance:使用数字、字母及特殊字符
	`)
	flag.Parse()
}


func main(){
}
