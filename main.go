package main

import (
	"bufio"
	"educational-lsp/rpc"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {}
