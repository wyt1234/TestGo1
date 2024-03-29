package main

import (
	"context"
	"fmt"
	"time"
)

// https://blog.csdn.net/qq_41078889/article/details/123908782
func monitor(ctx context.Context, number int) {
	for {
		select {
		// 其实可以写成 case <- ctx.Done()
		// 这里仅是为了让你看到 Done 返回的内容
		case v := <-ctx.Done():
			fmt.Printf("监控器%v，接收到通道值为：%v，监控结束。\n", number, v)
			return
		default:
			fmt.Printf("监控器%v，正在监控中...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 5; i++ {
		go monitor(ctx, i)
	}
	time.Sleep(1 * time.Second)
	// 关闭所有 goroutine 取消 context
	cancel()
	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)
	fmt.Println("主程序退出！！")
}
