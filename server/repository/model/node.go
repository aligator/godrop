package model

type FileNode struct {
	ID          string
	Name        string
	Description string
	IsFolder    bool
	MimeType    string
	Children    []FileNode
}

type CreateFileNode struct {
	Name        string
	Description string
	IsFolder    bool
	MimeType    string
	File        []byte
}
