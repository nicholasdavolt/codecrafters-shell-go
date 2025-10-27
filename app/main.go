package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// TODO: Uncomment the code below to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Fprint(os.Stdout, input+": command not found")

}
