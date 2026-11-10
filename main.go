package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nots - \033[94mNO T\033[0mrailing \u001B[94mS\u001B[0mpaces\nUsage:\n" +
			"\t\u001B[93mnots\u001B[0m <filename(s)...>")
		os.Exit(1)
	}
	changed := false
	for _, filename := range os.Args[1:] {
		bs, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file %q: %v\n", filename, err)
			continue
		}
		lineSeq := bytes.Lines(bs)
		lines := slices.Collect(lineSeq)
		if lines == nil {
			return
		}
		lineEnding := []byte("\n")
		winLineEnd := []byte("\r\n")
		if bytes.HasSuffix(lines[0], winLineEnd) {
			lineEnding = winLineEnd
		}
		outlines := make([][]byte, 0, 1000)
		for _, line := range lines {
			after := append(bytes.TrimRight(line, " \t\r\n"), lineEnding...)
			changed = changed || bytes.Compare(after, line) != 0
			outlines = append(outlines, after)
		}
		if changed {
			err = os.WriteFile(filename, slices.Concat(outlines...), 0644)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error writing file %q: %v\n", filename, err)
				continue
			}
		}
	}
}
