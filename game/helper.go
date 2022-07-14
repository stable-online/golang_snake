package game

import "fmt"

//全局捕获
func catch() {
	if err := recover(); err != nil {
		//记录日志等操作
		fmt.Println(fmt.Sprintf("record error info:%s", err))
	}
}

//执行方法
func run(f func()) {
	defer catch()
	f()
}
