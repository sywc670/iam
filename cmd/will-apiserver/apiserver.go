package main

import apiserver "github.com/marmotedu/iam/internal/will-apiserver"

func main() {
	apiserver.NewApp("apiserver").Run()
}
