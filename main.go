package main

import "github.com/dmnyu/ftk-tools/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
