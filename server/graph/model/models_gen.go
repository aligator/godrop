// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewNode struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	IsFolder    bool    `json:"isFolder"`
	MimeType    *string `json:"mimeType"`
	File        *string `json:"file"`
}

type Node struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	IsFolder    bool    `json:"isFolder"`
	MimeType    *string `json:"mimeType"`
	File        *string `json:"file"`
	Children    []Node  `json:"children"`
}
