package main

import "fmt"

//接口
type Map interface {
	put(key, value string)
	get(key string) string
}

type HashMap struct {

}

var pool = make(map[string]string)

// 实现接口
func (self *HashMap) put(key, value string)  {
	pool[key] = value
}

// 实现接口
func (self *HashMap) get(key string) (value string) {
	value = pool[key]
	return
}

func main() {
	hm := HashMap{}
	hm.put("name", "美女")

	fmt.Println("name == ", hm.get("name"))
}
