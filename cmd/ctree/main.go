package main

import (
	"bufio"
	"cronolabs/tools/internal/ctree"
	"fmt"
	"os"
	"strings"
)

func main() {
	tree := ctree.NewTree()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tree.AddPath(scanner.Text())
	}

	tree.Walk(func(file *ctree.PathFile, depth int) {
		indent := tabout(depth)
		end_mark := mark_end(file)
		fmt.Printf("%s%s%s\n", indent, file.Name, end_mark)
	})
}

func tabout(depth int) string {
	if depth > 0 {
		return strings.Repeat("    ", depth)
	}
	return ""
}

func mark_end(file *ctree.PathFile) string {
	if len(file.Children) > 0 || file.IsRoot {
		return "/"
	}
	return ""
}
