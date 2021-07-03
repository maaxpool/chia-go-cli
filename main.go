package main

import "github.com/LuttyYang/chia-go-cli/cmd"

func main() {
	err := cmd.RootCmd.Execute()

	if err != nil {
		panic(err)
	}
}
