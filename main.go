package main

import middleware "go_limit_demo/limit"

func main() {

	// 指定速率，也就是每秒允许多少个访问进来
	// 指定是否阻塞，一般使用非阻塞，阻塞会一直等待桶令牌
	limiter := middleware.NewTokenLimiter(10, false)

	// 1.尝试获取令牌 根据设置参数自动判断是阻塞还是非阻塞
	limiter.Allow()

	// 2.尝试非阻塞获取令牌 性能高于 Allow
	limiter.TryAcquire()

	// 3.尝试阻塞获取令牌 会阻塞到令牌桶中存在令牌为止
	limiter.Acquire()

	// 性能方面由快到慢排序排序  2 1 3
}
