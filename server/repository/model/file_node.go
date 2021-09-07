package model

type NodeState = string

const (
	NodeStateUpload NodeState = "UPLOAD"
	NodeStateReady  NodeState = "READY"
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
