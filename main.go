package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

var (
	port = flag.Int("port", 8888, "server port")
)

func main() {
	flag.Parse()

	router := rest.MustNewServer(rest.RestConf{
		Port: *port,
	})

	router.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: homeHandler,
	})

	router.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/register",
		Handler: registerHandler,
	})

	router.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/login",
		Handler: loginHandler,
	})

	defer router.Stop()

	fmt.Printf("Server is running on port %d...\n", *port)
	router.Start()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	httpx.OkJson(w, "Welcome to the home page!")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// 实现用户注册逻辑
	// 这里只是一个示例，实际应用中需要根据具体需求编写注册逻辑

	// 假设注册成功后返回用户信息
	user := map[string]interface{}{
		"username": "john_doe",
		"email":    "john@example.com",
	}

	httpx.OkJson(w, user)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 实现用户登录逻辑
	// 这里只是一个示例，实际应用中需要根据具体需求编写登录逻辑

	// 假设登录成功后返回用户信息
	user := map[string]interface{}{
		"username": "john_doe",
		"email":    "john@example.com",
	}

	httpx.OkJson(w, user)
}
