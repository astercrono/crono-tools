package ctree

import "fmt"

type PathFile struct {
	Name     string
	Children []*PathFile
	Parent   *PathFile
	IsRoot   bool
}

func NewFile(name string) *PathFile {
	return &PathFile{
		Name:     name,
		Children: []*PathFile{},
		Parent:   nil,
		IsRoot:   false,
	}
}

func (file *PathFile) PrintFile() {
	fmt.Printf("File Name: %s\n", file.Name)
}

func (parent *PathFile) AddChild(child *PathFile) {
	parent.Children = append(parent.Children, child)
}

func (file *PathFile) FindChild(Name string) *PathFile {
	if file.Children == nil {
		return nil
	}

	for _, file := range file.Children {
		if file.Name == Name {
			return file
		}
	}

	return nil
}
