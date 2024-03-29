package main

import (
	"fmt"
	"github.com/260by/tools/gssh"
)

func main() {
	ssh := &gssh.Server{
		Options: gssh.ServerOptions{
			Addr: "192.168.1.173",
			Port: "22",
			User: "root",
			KeyFile: "/home/keith/public_key/local",
		},
	}

	// 上传文件
	err := ssh.Put("tmp/test-logs/request*", "/tmp/ttt")
	if err != nil {
		fmt.Println(err)
	}

	// 下载文件
	// err := ssh.Get("/data/test-logs/", "tmp")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
