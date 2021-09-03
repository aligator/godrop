package model

import "io"

type FileNode struct {
	ID          string
	Name        string
	Description string
	IsFolder    bool
	MimeType    string
	Children    []FileNode
	Size        int64
}

type CreateFileNode struct {
	Path        string
	Name        string
	Description string
	IsFolder    bool
	MimeType    string
	File        io.Reader
}
