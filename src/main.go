package main

import "github.com/katsuokaisao/github-actions-go/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
