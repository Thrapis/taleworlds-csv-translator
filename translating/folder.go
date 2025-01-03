package translating

import (
	"fmt"
	"path/filepath"
)

type Folder struct {
	Name    string
	Path    string
	Files   []File
	Folders []Folder
}

func (f Folder) FullPath() string {
	return filepath.Join(f.Path, f.Name)
}

func (f Folder) String() string {
	return fmt.Sprintf("Folder: %s", f.Name)
}

func (f *Folder) PrintDeep() {
	fmt.Println(f)

	for _, file := range f.Files {
		fmt.Println(file)
	}

	for _, folder := range f.Folders {
		folder.PrintDeep()
	}
}

type File struct {
	FullName string
	Path     string
}

func (f File) FullPath() string {
	return filepath.Join(f.Path, f.FullName)
}

func (f File) String() string {
	return fmt.Sprintf("File: %s", f.FullName)
}
