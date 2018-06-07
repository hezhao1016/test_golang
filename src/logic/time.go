package main

import (
	"time"
	"fmt"
)

// 日期

func main() {
	// 当前时间
	now := time.Now()
	fmt.Println(now)

	// 睡眠, 单位为:1ns (纳秒) ps: 1毫秒 = 1000,000 纳秒
	time.Sleep(1000000000)  // 1秒
	fmt.Println("恢复运行...")

	// 获取两个时间的间隔
	timeout := time.Since(now)
	fmt.Println(timeout)
}
