package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有取消函数的context
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			// 监听Context是否被取消
			case <-ctx.Done():
				fmt.Println("Operation cancelled")
				return
			default:
				fmt.Println("Running task...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	// 让上面的goroutine先运行几次
	time.Sleep(5 * time.Second)

	// 调用取消函数，触发上面的goroutine退出
	cancel()

	// 为了能看到"Operation cancelled"，这里再等待一下
	time.Sleep(1 * time.Second)
}
