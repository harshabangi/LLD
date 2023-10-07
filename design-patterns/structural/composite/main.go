package main

import "fmt"

/*
Composite is a structural design pattern that allows composing objects into
a tree-like structure and work with it as if it was a singular object
*/

type Component interface {
	search(keyword string)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword '%s' in file '%s'\n", keyword, f.name)
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching for keyword '%s' in folder '%s'\n", keyword, f.name)
	for _, c := range f.components {
		c.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1"}

	folder1.add(file1)

	folder2 := &Folder{name: "Folder2"}
	folder2.add(folder1)
	folder2.add(file2)
	folder2.add(file3)

	folder2.search("rose")
}
