package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nots - \033[94mNO T\033[0mrailing \u001B[94mS\u001B[0mpaces\nUsage:\n\t\u001B[93mnots\u001B[0m <filename(s)...>")
		os.Exit(1)
	}
	for _, filename := range os.Args[1:] {
		bs, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file %q: %v\n", filename, err)
			continue
		}
		lines := bytes.Lines(bs)
		outlines := make([][]byte, 0, 1000)
		for line := range lines {
			outlines = append(outlines, bytes.TrimRight(line, " \t"))
		}
		err = os.WriteFile(filename, bytes.Join(outlines, []byte("")), 0644)
		if err != nil {
			fmt.Printf("Error writing file %q: %v\n", filename, err)
			continue
		}
	}
}
