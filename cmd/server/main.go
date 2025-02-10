package main

import "github.com/jpoz/esbuild-example/pkg/server"

func main() {
	server := server.New(":8080")
	server.Listen()
}
