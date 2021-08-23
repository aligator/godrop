package dto

import "github.com/aligator/godrop/server/repository/model"

func NodeFromModel(node model.Node) Node {
	result := Node{
		ID:          node.ID,
		Name:        node.Name,
		Description: node.Description,
		IsFolder:    node.IsFolder,
		MimeType:    &node.MimeType,
	}

	if node.MimeType == "" {
		result.MimeType = nil
	}

	if node.Children != nil {
		for _, child := range node.Children {
			result.Children = append(result.Children, NodeFromModel(child))
		}
	}

	return result
}

func CreateNodeFromDTO(node CreateNode) model.CreateNode {
	result := model.CreateNode{
		Name:        node.Name,
		Description: node.Description,
		IsFolder:    node.IsFolder,
	}

	if node.MimeType == nil {
		result.MimeType = ""
	} else {
		result.MimeType = *node.MimeType
	}

	if node.File == nil {
		result.File = []byte{}
	} else {
		result.File = []byte(*node.File)
	}

	return result
}
