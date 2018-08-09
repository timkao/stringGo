package main

import "fmt"

func concat() {
	str := "Beginning of the string " + "Ending of the string"
	fmt.Println(str)
}

/*
	在循环中使用加号 + 拼接字符串并不是最高效的做法，更好的办法是使用函数 strings.Join()（第 4.7.10 节），
	有没有更好地办法了？有！使用字节缓冲（bytes.Buffer）拼接更加给力（第 7.2.6 节）！
*/
