package main

import (
	"context"
	"pixel-haven/internal/config"
	"pixel-haven/server"
)

func main() {

	// 创建示例配置
	conf := &config.Config{}
	// 创建上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动服务器
	server.Start(ctx, conf)
}
