/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import "go-stuff/msg-util/cmd"

func main() {
	cmd.Execute()
}

type messageHandler struct{
	message string
	topic string
	sender string
	receiver string
}