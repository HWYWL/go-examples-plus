package main

import "fmt"

func main() {
	MyFunc1("校花")
	MyFunc2("Hi", "，美女")
}

func MyFunc1(name string)  {
	fmt.Println("Hello ", name)
}

func MyFunc2(say string, name string)  {
	fmt.Printf("%v %v\n", say, name)
}
