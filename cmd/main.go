package main

import (
	"log"
	"mitm-proxy/pkg/auth"
	"mitm-proxy/pkg/config"
	"net/http"
)

func handleRequest(response http.ResponseWriter, request *http.Request) {
	// ... 代理的核心代码
}

func main() {
	cfg, err := config.LoadConfig("path/to/config.json")
	if err != nil {
		log.Fatal(err)
	}

	// 使用cfg.CACert和cfg.CAKey加载CA证书和私钥
	// ...

	http.HandleFunc("/", auth.BasicAuth(handleRequest, "yourUsername", "yourPassword", "Please enter your username and password"))
	http.ListenAndServe(":8080", nil)
}
