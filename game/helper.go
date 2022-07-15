package game

import (
	"fmt"
)

//全局捕获
func catch() {
	if err := recover(); err != nil {

		//打印异常
		fmt.Println(fmt.Sprintf("Catch faild:%s", err))

		//获取堆栈信息(暂时先不获取吧.)
		/*buf := new(bytes.Buffer)
		fmt.Fprintf(buf, "%v\n", err)
		for i := 1; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		}
		fmt.Println(buf.String())*/
	}
}

//执行方法
func run(f func()) {
	defer catch()
	f()
}
