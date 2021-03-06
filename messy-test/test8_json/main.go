package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"test9_time"`
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	fmt.Println(string(b)) //{"name":"Alice","body":"Hello","test9_time":1294706395881547000}
	var msg1 **Message
	var msg2 = &Message{}
	var msg3 = Message{}
	var msg Message
	//很神奇，传msg就不行，传&msg就可以
	//大致看了下源码，首先如果传msg1其实就相当于传了一个nil，在反射阶段就直接报错了
	//而传入指针的地址会通过反射创建一个值 见：decode.go/indirect 488行
	//json.Unmarshal(b, msg)
	json.Unmarshal(b, &msg)
	fmt.Println(msg) //{Alice Hello 1294706395881547000}
	json.Unmarshal(b, &msg1)
	fmt.Println(*msg1) //&{Alice Hello 1294706395881547000}
	json.Unmarshal(b, &msg2)
	fmt.Println(msg2) // &{Alice Hello 1294706395881547000}
	json.Unmarshal(b, &msg3)
	fmt.Println(msg3) //{Alice Hello 1294706395881547000}
}
