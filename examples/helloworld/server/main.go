package main

import (
	"log"
	"net/http"

	"github.com/guihouchang/swagger-api/openapiv2"
)

func main() {
	// 创建默认的handler（使用默认路径前缀 /q/swagger-ui）
	defaultHandler := openapiv2.NewHandler()

	// 创建自定义路径前缀的handler（用于测试Nginx rewrite场景）
	customHandler := openapiv2.NewHandler(
		openapiv2.WithPathPrefix("/user/q/swagger-ui"),
	)

	// 启动默认服务器
	go func() {
		log.Println("Default Swagger UI server starting on :8000")
		log.Println("Visit: http://localhost:8000/q/swagger-ui/")
		if err := http.ListenAndServe(":8000", defaultHandler); err != nil {
			log.Fatal("Default server failed:", err)
		}
	}()

	// 启动自定义路径前缀服务器
	log.Println("Custom path prefix Swagger UI server starting on :8001")
	log.Println("Visit: http://localhost:8001/user/q/swagger-ui/")
	if err := http.ListenAndServe(":8001", customHandler); err != nil {
		log.Fatal("Custom server failed:", err)
	}
}
