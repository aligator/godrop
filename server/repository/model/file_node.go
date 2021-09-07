package model

import (
	"github.com/aligator/godrop/server/graph/dto"
)

type NodeState = dto.NodeState

const (
	NodeStateUpload = dto.NodeStateUpload
	NodeStateReady  = dto.NodeStateReady
)

type FileNode struct {
	ID          string
	Name        string
	Description string
	IsFolder    bool
	State       NodeState
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
}
