package ctree

import (
	"strings"
)

type PathTree struct {
	Root *PathFile
}

func NewTree() *PathTree {
	tree := PathTree{
		Root: NewFile(""),
	}
	tree.Root.IsRoot = true
	return &tree
}

func (tree *PathTree) AddPath(path string) {
	path_parts := strings.Split(path, "/")
	parent := tree.Root

	for _, part := range path_parts {
		if part == "" {
			continue
		}

		part_name := strings.Trim(part, " ")
		child_match := parent.FindChild(part_name)

		if child_match == nil {
			path_file := NewFile(part_name)
			parent.AddChild(path_file)
			path_file.Parent = parent
			parent = path_file
		} else {
			parent = child_match
		}
	}
}

func (tree *PathTree) Walk(path_f func(current *PathFile, depth int)) {
	tree.walk(tree.Root, 0, path_f)
}

func (tree *PathTree) walk(current *PathFile, depth int, path_f func(current *PathFile, depth int)) {
	if current == nil {
		current = tree.Root
	}

	if path_f != nil {
		path_f(current, depth)
	}

	for _, child := range current.Children {
		tree.walk(child, depth+1, path_f)
	}
}
