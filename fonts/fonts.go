package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	fonts, err := returnAllTypes("list.txt")
	if err != nil {
		fmt.Printf("problem listing types: %v", err)
		return
	}
	for _, font := range fonts {
		fmt.Println(font)
		figure.NewFigure(font, font, true).Print()
		time.Sleep(1 * time.Second)
		fmt.Println(strings.Repeat("-", 80))
	}

}

func returnAllTypes(file string) ([]string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read file '%v': %w", file, err)
	}
	split := strings.Split(string(bytes), "\n")
	return split, nil
}
