package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\x00')

	fmt.Println(str)
}
