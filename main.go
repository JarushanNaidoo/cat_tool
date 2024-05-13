package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

struct commands {

}

func main() {
	// read input after cat until new line
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading instruction")
		return
	}
	fmt.Print(input)
}
