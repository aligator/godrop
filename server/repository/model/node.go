package model

type Node struct {
	ID          string
	Name        string
	Description string
	IsFolder    bool
	MimeType    string
	FileURL     string
	Children    []Node
}

type CreateNode struct {
	Name        string
	Description string
	IsFolder    bool
	MimeType    string
	File        []byte
}
