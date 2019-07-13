package util

import (
	"bufio"
	"fmt"
	"os"
)

func Ask(msg string) string {
	fmt.Printf("%s\n> ", msg)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	return Trim(value)
}
